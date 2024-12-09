package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/bearaujus/steam-utils/internal/view/cli"
	"github.com/bearaujus/steam-utils/internal/view/interactive"
	"github.com/inconshreveable/mousetrap"
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
	ctx := context.TODO()
	cfg := config.NewConfig(&config.LdFlags{
		Name:    name,
		Version: version,
		Arch:    arch,
		Goos:    goos,
		File:    file,
	})

	var err error
	if cfg.LdFlags.Goos == "windows" && mousetrap.StartedByExplorer() {
		err = interactive.New(ctx, cfg).Run(ctx)
	} else {
		err = cli.New(ctx, cfg).Run(ctx)
	}
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(time.Second * 10)
		os.Exit(1)
		return
	}

	os.Exit(0)
}
