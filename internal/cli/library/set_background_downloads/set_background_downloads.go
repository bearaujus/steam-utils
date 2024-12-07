package set_background_downloads

import (
	"context"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/bearaujus/steam-utils/internal/usecase/library/set_background_downloads"
	"github.com/spf13/cobra"
)

func New(ctx context.Context, cfg *config.Config) []*cobra.Command {
	return []*cobra.Command{
		new0Cmd(ctx, cfg),
		new1Cmd(ctx, cfg),
		new2Cmd(ctx, cfg),
	}
}

func new0Cmd(ctx context.Context, cfg *config.Config) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "0",
		Short: "Follow your global steam settings",
		Args:  cobra.NoArgs,
		RunE:  set_background_downloads.NewCmdRunner(ctx, cfg),
	}
	return cmd
}

func new1Cmd(ctx context.Context, cfg *config.Config) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "1",
		Short: "Always allow background downloads",
		Args:  cobra.NoArgs,
		RunE:  set_background_downloads.NewCmdRunner(ctx, cfg),
	}
	return cmd
}

func new2Cmd(ctx context.Context, cfg *config.Config) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "2",
		Short: "Never allow background downloads",
		Args:  cobra.NoArgs,
		RunE:  set_background_downloads.NewCmdRunner(ctx, cfg),
	}
	return cmd
}
