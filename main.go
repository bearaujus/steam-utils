package main

import (
	"context"
	"github.com/bearaujus/steam-utils/internal/cmd"
	"github.com/bearaujus/steam-utils/internal/config"
)

func main() {
	cfg := &config.Config{}
	var rootCmd = cmd.NewCommand(context.TODO(), cfg)
	err := config.LoadConfig(rootCmd, cfg)
	if err != nil {
		panic(err)
	}
	err = rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
