package model

import (
	"github.com/bearaujus/berror"
)

var (
	ErrFailToInitializeSteamPath   = berror.NewErrDefinition("fail to initialize steam path: %v", berror.OptionErrDefinitionWithDisabledStackTrace())
	ErrEmptyListLibraryMetadata    = berror.NewErrDefinition("empty list library metadata: no .acf files detected in %v directory. ensure that you have installed applications in your Steam library and try again", berror.OptionErrDefinitionWithDisabledStackTrace())
	ErrReadDirectory               = berror.NewErrDefinition("fail to read directory: %v", berror.OptionErrDefinitionWithDisabledStackTrace())
	ErrReadFile                    = berror.NewErrDefinition("fail to read file: %v", berror.OptionErrDefinitionWithDisabledStackTrace())
	ErrWriteFile                   = berror.NewErrDefinition("fail to write file: %v", berror.OptionErrDefinitionWithDisabledStackTrace())
	ErrParseSteamACFFile           = berror.NewErrDefinition("fail to parse steam acf: %v", berror.OptionErrDefinitionWithDisabledStackTrace())
	ErrUpdateValueFromSteamACFFile = berror.NewErrDefinition("fail to update value to steam acf: %v", berror.OptionErrDefinitionWithDisabledStackTrace())
)
