package library

import (
	"context"
	"github.com/bearaujus/steam-utils/internal/cli/library/set_auto_update"
	"github.com/bearaujus/steam-utils/internal/cli/library/set_background_downloads"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/spf13/cobra"
)

func New(ctx context.Context, cfg *config.Config) []*cobra.Command {
	return []*cobra.Command{
		newSetAutoUpdateCmd(ctx, cfg),
		newSetBackgroundDownloadsCmd(ctx, cfg),
	}
}

func newSetAutoUpdateCmd(ctx context.Context, cfg *config.Config) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "set-auto-update",
		Short: "Set auto update behavior on all collections on your Steam library",
		Args:  cobra.NoArgs,
	}
	for _, childCmd := range set_auto_update.New(ctx, cfg) {
		cmd.AddCommand(childCmd)
	}
	return cmd
}

func newSetBackgroundDownloadsCmd(ctx context.Context, cfg *config.Config) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "set-background-downloads",
		Short: "Set background downloads behavior on all collections on your Steam library",
		Args:  cobra.NoArgs,
	}
	for _, childCmd := range set_background_downloads.New(ctx, cfg) {
		cmd.AddCommand(childCmd)
	}
	return cmd
}
