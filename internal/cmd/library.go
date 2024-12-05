package cmd

import (
	"context"
	"github.com/bearaujus/steam-utils/internal/cmd/library"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/spf13/cobra"
)

func newLibraryCmd(ctx context.Context, config *config.Config) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "library",
		Short: "Steam library utilities",
		Args:  cobra.NoArgs,
	}
	cmd.SetContext(ctx)
	for _, childCmd := range library.NewLibraryCommands(ctx, config) {
		cmd.AddCommand(childCmd)
	}
	return cmd
}
