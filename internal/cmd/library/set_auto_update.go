package library

import (
	"context"
	"github.com/bearaujus/steam-utils/internal/cmd/library/set_auto_update"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/spf13/cobra"
)

func newSetAutoUpdateCmd(ctx context.Context, cfg *config.Config) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "set-auto-update",
		Short: "Set auto update behavior on all collections on your Steam library",
		Args:  cobra.NoArgs,
	}
	for _, childCmd := range set_auto_update.NewCommands(ctx, cfg) {
		cmd.AddCommand(childCmd)
	}
	return cmd
}
