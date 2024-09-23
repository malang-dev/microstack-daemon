package main

import (
	"fmt"
	"microstack/cmd/microstackd/cmd"
	"microstack/pkg/logs"
	"os"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
)

func NewMicrostakdCommand() *cobra.Command {
	conf := &logs.Config{Verbosity: 0, Format: "text", Output: "stdout"}
	root := &cobra.Command{
		Use:           "microstackd",
		Short:         "A self-sufficient runtime for microVMs",
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Configure logging
			if err := logs.Configure(conf); err != nil {
				return fmt.Errorf("configuring logging: %w", err)
			}

			return nil
		},
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

	root.AddCommand(cmd.NewInstallCommand())
	root.AddCommand(cmd.NewUninstallCommand())
	root.AddCommand(cmd.NewVersionCommand())

	return root
}

func main() {
	if err := Run(); err != nil {
		os.Exit(1)
	}
}

// Run runs the main cobra command of this application
func Run() error {
	c := NewMicrostakdCommand()
	return c.Execute()
}
