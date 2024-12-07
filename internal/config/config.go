package config

import (
	"github.com/bearaujus/steam-utils/pkg/steam_path"
	"github.com/spf13/cobra"
)

const (
	PersistentFlagSteamPath = "steam-path"
)

type Config struct {
	SteamPath string
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
