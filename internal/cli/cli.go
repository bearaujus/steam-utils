package cli

import (
	"context"
	"github.com/bearaujus/steam-utils/internal/cli/library"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/spf13/cobra"
)

func NewRoot(ctx context.Context, cfg *config.Config) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   cfg.LdFlags.Name,
		Short: "Steam utilities",
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

func newLibraryCmd(ctx context.Context, cfg *config.Config) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "library",
		Short: "Steam library utilities",
		Args:  cobra.NoArgs,
	}
	for _, childCmd := range library.New(ctx, cfg) {
		cmd.AddCommand(childCmd)
	}
	return cmd
}
