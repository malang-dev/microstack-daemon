package cmd

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
)

func NewUninstallCommand() *cobra.Command {
	root := &cobra.Command{
		Use:           "uninstall",
		Short:         "Uninstall pre-requisites",
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println()
			pterm.DefaultBigText.WithLetters(
				putils.LettersFromStringWithStyle("Micro", pterm.NewStyle(pterm.FgLightGreen)),
				putils.LettersFromStringWithStyle("stack", pterm.NewStyle(pterm.FgLightBlue))).
				Render()
			fmt.Println()
			return cmd.Help()
		},
	}

	return root
}
