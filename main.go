package main

import (
	"context"
	"fmt"
	"github.com/bearaujus/steam-utils/internal/cli"
	"github.com/bearaujus/steam-utils/internal/config"
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
	fmt.Println(name, version, arch, goos, file)
	cfg := &config.Config{}
	var rootCLI = cli.NewRoot(context.TODO(), cfg)
	err := config.LoadConfig(rootCLI, cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = rootCLI.Execute()
}
