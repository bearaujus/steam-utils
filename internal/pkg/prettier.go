package pkg

import (
	"fmt"
	"strings"

	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/fatih/color"
)

const maxWidth = 100

func GetTitle(cfg *config.Config) string {
	title := color.New(color.FgHiYellow, color.Bold).Sprint(strings.ToUpper(cfg.LdFlags.Name))
	version := color.New(color.FgHiBlue).Sprint(cfg.LdFlags.Version)
	return fmt.Sprintf("%v (%v) - %v", title, version, "Sets of utilities for managing your Steam")
}

func GetTitleRaw(cfg *config.Config) string {
	return fmt.Sprintf("%v (%v) - %v", cfg.LdFlags.Name, cfg.LdFlags.Version, "Sets of utilities for managing your Steam")
}

func PrintTitle(cfg *config.Config) {
	PrintSep()
	fmt.Println(GetTitle(cfg))
	PrintSep()
}

func PrintSep() {
	fmt.Println(strings.Repeat("-", maxWidth))
}
