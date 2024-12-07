package config

import (
	"fmt"
	"github.com/bearaujus/steam-utils/pkg/steam_path"
	"github.com/spf13/cobra"
	"runtime"
)

const (
	PersistentFlagSteamPath = "steam-path"
)

type Config struct {
	SteamPath string
	LdFlags   *LdFlags
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
	return &Config{
		LdFlags: ldFlags,
	}
}

func LoadConfig(cmd *cobra.Command, config *Config) error {
	var defaultSteamPath string
	sp, err := steam_path.LoadDefaultSteamPath()
	if err == nil {
		defaultSteamPath = sp.Base()
	}
	cmd.PersistentFlags().StringVar(&config.SteamPath, PersistentFlagSteamPath, defaultSteamPath, "Path to steam installation directory")
	return nil
}
