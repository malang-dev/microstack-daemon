package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	root := &cobra.Command{
		Use:           "version",
		Aliases:       []string{"v"},
		Short:         "Show current installed version",
		SilenceErrors: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s\nCommit Hash: %s\n", "v0.0.1", "8f8dcd1a1346e596c9ba889bfc002556d1a82ac0")
		},
	}

	return root
}
