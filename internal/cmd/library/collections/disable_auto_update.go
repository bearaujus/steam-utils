package collections

import (
	"context"
	"fmt"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/spf13/cobra"
)

func newDisableAutoUpdateCmd(ctx context.Context, config *config.Config) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "disable-auto-update",
		Short: "Disable auto update on all collections on your Steam library",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(cmd.Flags().GetBool("asd"))
			return nil
		},
	}
	cmd.SetContext(ctx)
	cmd.Flags().BoolP("asd", "a", false, "show asd")
	return cmd
}
