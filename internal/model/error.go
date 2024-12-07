package model

import (
	"github.com/bearaujus/berror"
)

var (
	ErrSteamPathIsNotSet           = berror.NewErrDefinition("default steam installation path is not detected. please include flag '--%v'", berror.OptionErrDefinitionWithDisabledStackTrace())
	ErrSteamPathIsInvalid          = berror.NewErrDefinition("%v. please update argument on flag '--%v'", berror.OptionErrDefinitionWithDisabledStackTrace())
	ErrReadDirectory               = berror.NewErrDefinition("fail to read directory: %v", berror.OptionErrDefinitionWithDisabledStackTrace())
	ErrReadFile                    = berror.NewErrDefinition("fail to read file: %v", berror.OptionErrDefinitionWithDisabledStackTrace())
	ErrWriteFile                   = berror.NewErrDefinition("fail to write file: %v", berror.OptionErrDefinitionWithDisabledStackTrace())
	ErrParseSteamACFFile           = berror.NewErrDefinition("fail to parse steam acf: %v", berror.OptionErrDefinitionWithDisabledStackTrace())
	ErrGetValueFromSteamACFFile    = berror.NewErrDefinition("fail to get value from steam acf: %v", berror.OptionErrDefinitionWithDisabledStackTrace())
	ErrUpdateValueFromSteamACFFile = berror.NewErrDefinition("fail to update value to steam acf: %v", berror.OptionErrDefinitionWithDisabledStackTrace())
)
