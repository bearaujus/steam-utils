package steam_path

import "errors"

var (
	ErrEmptyPath           = errors.New("empty steam installation path")
	ErrInvalidPath         = errors.New("invalid steam installation path")
	ErrUnsupportedOS       = errors.New("unsupported operating system")
	ErrDefaultPathNotFound = errors.New("default steam installation not found")
)
