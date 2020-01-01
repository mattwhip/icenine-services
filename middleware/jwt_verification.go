package middleware

import (
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

const (
	// CtxKeyUserID is the context key to retrieve the user ID set by JWTVerification middleware
	CtxKeyUserID = "user_id"
)

// JWTVerification parses a JSON Web Token from the request headers, verifies the signature,
// and adds the subject from the token claims (user ID) to the context for handler retrieval.
var JWTVerification = func(signingSecret string) buffalo.MiddlewareFunc {
	return func(h buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {

			// Retrieve JWT authorization header
			authHeaders := c.Request().Header["Authorization"]
			jwtString := ""
			for _, header := range authHeaders {
				if strings.HasPrefix(header, "JWT") {
					splits := strings.Split(header, " ")
					if len(splits) != 2 {
						return errors.New("Invalid JWT Authorization header format")
					}
					jwtString = splits[1]
					break
				}
			}
			if jwtString == "" {
				return errors.New("JWT Authorization header not found")
			}

			// Parse takes the token string and a function for looking up the key. The latter is especially
			// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
			// head of the token to identify which key to use, but the parsed token (head and claims) is provided
			// to the callback, providing flexibility.
			token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return []byte(signingSecret), nil
			})
			if err != nil {
				return err
			}

			// Parse the user ID from the claims
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if val, ok := claims["sub"]; ok {
					userID := val.(string)
					c.Set(CtxKeyUserID, userID)
				}
			} else {
				return errors.New("Invalid JWT")
			}

			// Execute the next middleware/handler in the chain
			if err := h(c); err != nil {
				return err
			}

			return nil
		}
	}
}
