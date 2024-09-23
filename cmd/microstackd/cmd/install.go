package cmd

import (
	"microstack/internal/config"
	"microstack/pkg/logs"
	"microstack/pkg/script"
	"os/exec"

	"github.com/spf13/cobra"
)

var forceDownload bool

func NewInstallCommand() *cobra.Command {
	root := &cobra.Command{
		Use:           "install",
		Short:         "Install all microstackd pre-requisites",
		SilenceErrors: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			logs.SetContext(cmd.Context()) // Set logs context

			if !forceDownload {
				forceDownload = !config.IsReleasedTagVersion(config.TagVersion)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Download install script
			scripts := []script.Type{
				script.InstallPrerequisites,
			}

			for _, s := range scripts {
				if err := script.Download(s, config.TagVersion, forceDownload); err != nil {
					return err
				}

				if err := script.Run(s, config.TagVersion, createSetupInstallCommandEnvsFunc()); err != nil {
					return err
				}
			}
			return nil
		},
	}

	return root
}

func createSetupInstallCommandEnvsFunc() func(cmd *exec.Cmd) error {
	return func(cmd *exec.Cmd) error {
		cmd.Env = append(
			cmd.Env,
			config.ExpectedPrerequisiteVersionsEnvVars()...,
		)

		return nil
	}
}
