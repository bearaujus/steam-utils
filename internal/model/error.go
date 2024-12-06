package model

import (
	"github.com/bearaujus/berror"
)

var (
	ErrSteamPathIsNotSet = berror.NewErrDefinition("%v. please include flag '--%v'", berror.OptionErrDefinitionWithDisabledStackTrace())
	ErrInvalidSteamPath  = berror.NewErrDefinition("%v. please update argument on flag '--%v'", berror.OptionErrDefinitionWithDisabledStackTrace())
)
