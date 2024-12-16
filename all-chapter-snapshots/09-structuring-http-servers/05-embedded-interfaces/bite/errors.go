package bite

import "errors"

var (
	ErrExists         = errors.New("already exists")
	ErrNotExist       = errors.New("does not exist")
	ErrInvalidRequest = errors.New("invalid request")
	ErrInternal       = errors.New(
		"internal error: please try again later or contact support",
	)
)

// other shared types—structs, interfaces, etc., maybe in different files.
