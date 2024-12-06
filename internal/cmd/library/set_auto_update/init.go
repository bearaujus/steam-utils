package set_auto_update

import (
	"context"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/spf13/cobra"
)

func NewCommands(ctx context.Context, cfg *config.Config) []*cobra.Command {
	return []*cobra.Command{
		new0Cmd(ctx, cfg),
		new1Cmd(ctx, cfg),
	}
}
