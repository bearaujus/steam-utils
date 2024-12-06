package usecase

import "github.com/spf13/cobra"

type CmdRunner func(cmd *cobra.Command, args []string) error
