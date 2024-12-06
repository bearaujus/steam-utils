package cmd

import (
	"context"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/spf13/cobra"
)

func NewCommand(ctx context.Context, cfg *config.Config) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "steam-utils",
		Short: "Sets of utilities for managing your Steam",
		Args:  cobra.NoArgs,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   true,
			DisableNoDescFlag:   true,
			DisableDescriptions: true,
			HiddenDefaultCmd:    true,
		},
	}
	rootCmd.AddCommand(newLibraryCmd(ctx, cfg))
	return rootCmd
}
