package sqlx

import (
	"database/sql/driver"
	"encoding/base64"
	"fmt"
)

// Base64String is a base64 encoded string to persist
// and retrieve base64 encoded strings.
type Base64String string

// Value encodes the string to base64.
func (s Base64String) Value() (driver.Value, error) {
	dst := []byte(s)
	return base64.StdEncoding.EncodeToString(dst), nil
}

// Scan decodes the base64 encoded string.
func (s *Base64String) Scan(src any) error {
	ss, ok := src.(string)
	if !ok {
		return fmt.Errorf("%q is %T, not string", ss, src)
	}
	dst, err := base64.StdEncoding.DecodeString(ss)
	if err != nil {
		return fmt.Errorf("decoding %q: %w", ss, err)
	}
	*s = Base64String(dst)
	return nil
}

// String returns the base64 encoded value as a string.
func (s Base64String) String() string {
	return string(s)
}
