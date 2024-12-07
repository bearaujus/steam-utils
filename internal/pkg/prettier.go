package pkg

import (
	"fmt"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/fatih/color"
	"strings"
)

const (
	maxWidth = 100
)

func PrintTitle(cfg *config.Config) {
	PrintSep()
	title := color.New(color.FgHiYellow, color.Bold).Sprint(strings.ToUpper(cfg.LdFlags.Name))
	version := color.New(color.FgHiBlue).Sprint(cfg.LdFlags.Version)
	fmt.Printf("%v (%v) - %v\n", title, version, "Sets of utilities for managing your Steam")
	PrintSep()
}

func PrintSep() {
	fmt.Println(strings.Repeat("-", maxWidth))
}
