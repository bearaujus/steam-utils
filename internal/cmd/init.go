package cmd

import (
	"context"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/spf13/cobra"
)

func New(ctx context.Context, config *config.Config) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "steam-utils",
		Short: "A brief description of your application",
		Args:  cobra.NoArgs,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   true,
			DisableNoDescFlag:   true,
			DisableDescriptions: true,
			HiddenDefaultCmd:    true,
		},
	}
	rootCmd.SetContext(ctx)
	rootCmd.AddCommand(newLibraryCmd(ctx, config))
	return rootCmd
}
