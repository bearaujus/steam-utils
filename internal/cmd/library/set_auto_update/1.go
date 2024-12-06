package set_auto_update

import (
	"context"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/bearaujus/steam-utils/internal/usecase"
	"github.com/spf13/cobra"
)

func new1Cmd(ctx context.Context, config *config.Config) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "1",
		Short: "Update your Steam library on launch",
		Args:  cobra.NoArgs,
		RunE:  usecase.NewSetAutoUpdate(ctx, config),
	}
	return cmd
}
