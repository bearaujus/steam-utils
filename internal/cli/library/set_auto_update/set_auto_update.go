package set_auto_update

import (
	"context"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/bearaujus/steam-utils/internal/usecase/library/set_auto_update"
	"github.com/spf13/cobra"
)

func New(ctx context.Context, cfg *config.Config) []*cobra.Command {
	return []*cobra.Command{
		new0Cmd(ctx, cfg),
		new1Cmd(ctx, cfg),
	}
}

func new0Cmd(ctx context.Context, cfg *config.Config) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "0",
		Short: "Always keep all games updated",
		Args:  cobra.NoArgs,
		RunE:  set_auto_update.NewCmdRunner(ctx, cfg),
	}
	return cmd
}

func new1Cmd(ctx context.Context, cfg *config.Config) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "1",
		Short: "Only update a game when you launch it",
		Args:  cobra.NoArgs,
		RunE:  set_auto_update.NewCmdRunner(ctx, cfg),
	}
	return cmd
}
