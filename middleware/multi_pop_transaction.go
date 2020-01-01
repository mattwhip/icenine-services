package middleware

import (
	"time"

	"github.com/pkg/errors"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// MultiPopTransaction is based on the PopTransaction Buffalo middlewar
// that wraps each request in a transaction. This middleware takes an
// additional suffix that will be added to the key in the context:
//    "tx_<suffix>" instead of "tx"
// This is intended for use where multiple databases must be connected.
// Otherwise the standard PopTransaction middleware should be used for
// the default single database case.
var MultiPopTransaction = func(db *pop.Connection, suffix string) buffalo.MiddlewareFunc {
	return func(h buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			// wrap all requests in a transaction and set the length
			// of time doing things in the db to the log.
			err := db.Transaction(func(tx *pop.Connection) error {
				start := tx.Elapsed
				defer func() {
					finished := tx.Elapsed
					elapsed := time.Duration(finished - start)
					c.LogField("db_"+suffix, elapsed)
				}()
				c.Set("tx_"+suffix, tx)
				if err := h(c); err != nil {
					return err
				}
				if res, ok := c.Response().(*buffalo.Response); ok {
					if res.Status < 200 || res.Status >= 400 {
						return errNonSuccess
					}
				}
				return nil
			})
			if err != nil && errors.Cause(err) != errNonSuccess {
				return err
			}
			return nil
		}
	}
}

var errNonSuccess = errors.New("non success status code")
