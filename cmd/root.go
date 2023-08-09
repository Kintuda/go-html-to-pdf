package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	ctx := context.Background()

	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "Notification cli",
	}

	rootCmd.AddCommand(NewServerCmd(ctx))

	return rootCmd
}
