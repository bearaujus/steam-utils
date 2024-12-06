package library

import (
	"context"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/spf13/cobra"
)

func NewCommands(ctx context.Context, cfg *config.Config) []*cobra.Command {
	return []*cobra.Command{
		newSetAutoUpdateCmd(ctx, cfg),
	}
}
