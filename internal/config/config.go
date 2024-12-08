package config

import (
	"fmt"
	"runtime"

	"github.com/bearaujus/steam-utils/pkg/steam_path"
)

type Config struct {
	SteamPath        steam_path.SteamPath
	DefaultSteamPath steam_path.SteamPath
	LdFlags          *LdFlags
}

type LdFlags struct {
	Name    string
	Version string
	Arch    string
	Goos    string
	File    string
}

func NewConfig(ldFlags *LdFlags) *Config {
	if ldFlags == nil {
		ldFlags = &LdFlags{}
	}
	if ldFlags.Name == "" {
		ldFlags.Name = "steam-utils"
	}
	if ldFlags.Version == "" {
		ldFlags.Version = "v0.0.0-dev"
	}
	if ldFlags.Goos == "" {
		ldFlags.Goos = runtime.GOOS
	}
	if ldFlags.Arch == "" {
		ldFlags.Arch = runtime.GOARCH
	}
	if ldFlags.File == "" {
		ldFlags.File = fmt.Sprintf("%v-%v-%v-%v", ldFlags.Name, ldFlags.Version, ldFlags.Goos, ldFlags.Arch)
	}
	cfg := &Config{LdFlags: ldFlags}
	sp, err := steam_path.LoadDefaultSteamPath()
	if err == nil {
		cfg.DefaultSteamPath = sp
	}
	return cfg
}
