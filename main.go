package main

import (
	"context"
	"fmt"
	"github.com/bearaujus/steam-utils/internal/cli"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/bearaujus/steam-utils/internal/pkg"
)

// these variable will be retrieved from -ldflags
var (
	name    string // main.name
	version string // main.version
	arch    string // main.arch
	goos    string // main.goos
	file    string // main.file
)

func main() {
	var (
		cfg = config.NewConfig(&config.LdFlags{
			Name:    name,
			Version: version,
			Arch:    arch,
			Goos:    goos,
			File:    file,
		})
		ctx     = context.TODO()
		rootCLI = cli.NewRoot(ctx, cfg)
	)
	if err := config.LoadConfig(rootCLI, cfg); err != nil {
		fmt.Println(err)
		return
	}
	pkg.PrintTitle(cfg)
	_ = rootCLI.Execute()
	fmt.Println()
}
