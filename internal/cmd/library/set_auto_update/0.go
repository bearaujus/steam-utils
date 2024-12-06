package set_auto_update

import (
	"context"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/bearaujus/steam-utils/internal/usecase"
	"github.com/spf13/cobra"
)

func new0Cmd(ctx context.Context, config *config.Config) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "0",
		Short: "Always keep your Steam library updated",
		Args:  cobra.NoArgs,
		RunE:  usecase.NewSetAutoUpdate(ctx, config),
	}
	return cmd
}
