package uuid

import (
	"encoding/hex"
	"fmt"

	u "github.com/gobuffalo/uuid"
	"github.com/pkg/errors"
)

const (
	// UUIDBytesSize is the number of bytes in a UUID
	UUIDBytesSize = 16
)

// New creates a new Universally Unique ID string
func New() (string, error) {
	newUUID, err := u.NewV4()
	if err != nil {
		return "", errors.Wrap(err, "failed to create new UUID")
	}
	return FromBytes(newUUID.Bytes())
}

// FromBytes converts a byte slice to a UUID string
func FromBytes(bytes []byte) (string, error) {
	if len(bytes) != UUIDBytesSize {
		return "", errors.Errorf("byte slice of length %v cannot be converted to UUID of length %v",
			len(bytes), UUIDBytesSize)
	}
	return hex.EncodeToString(bytes), nil
}

// ToBytes converts a string UUID to a byte slice
func ToBytes(uuid string) ([]byte, error) {
	if !Validate(uuid) {
		return nil, errors.New(fmt.Sprintf("string '%v' cannot be converted to a UUID byte slice",
			len(uuid)))
	}
	bytes, err := hex.DecodeString(uuid)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode UUID string as hex")
	}
	return bytes, nil
}

// Validate a uuid string to make sure it is the correct length
func Validate(uuid string) bool {
	// A UUID string is a hex representation of 16 bytes of data. In hex, each half-byte of data is
	// represented as one character (0-F). So the total length of a UUID string is twice that of the binary
	// representation.
	return len(uuid)/2 == UUIDBytesSize
}
