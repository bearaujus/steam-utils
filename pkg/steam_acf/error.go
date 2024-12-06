package steam_acf

import "errors"

var (
	ErrEmptyData  = errors.New("empty data")
	ErrParseData  = errors.New("error when parsing data")
	ErrGetData    = errors.New("error when get data")
	ErrUpdateData = errors.New("error when update data")
)
