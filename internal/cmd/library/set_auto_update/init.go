package set_auto_update

import (
	"context"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/spf13/cobra"
)

func NewCommands(ctx context.Context, config *config.Config) []*cobra.Command {
	return []*cobra.Command{
		new0Cmd(ctx, config),
		new1Cmd(ctx, config),
	}
}
