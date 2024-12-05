package collections

import (
	"context"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/spf13/cobra"
)

func NewCollectionsCommands(ctx context.Context, config *config.Config) []*cobra.Command {
	return []*cobra.Command{
		newDisableAutoUpdateCmd(ctx, config),
	}
}
