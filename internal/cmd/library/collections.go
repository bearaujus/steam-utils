package library

import (
	"context"
	"github.com/bearaujus/steam-utils/internal/cmd/library/collections"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/spf13/cobra"
)

func newCollectionsCmd(ctx context.Context, config *config.Config) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "collections",
		Short: "Steam library collections utilities",
		Args:  cobra.NoArgs,
	}
	cmd.SetContext(ctx)
	for _, childCmd := range collections.NewCollectionsCommands(ctx, config) {
		cmd.AddCommand(childCmd)
	}
	return cmd
}
