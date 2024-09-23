package config

import "fmt"

var (
	FirecrackerVersion string
	ContainerdVersion  string
)

type EnvVars []string

// ExpectedPrerequisiteVersionsEnvVars get the expected prerequisite versions in the environment variables.
func ExpectedPrerequisiteVersionsEnvVars() EnvVars {
	return []string{
		fmt.Sprintf("MICROSTACKD_VERSION=%s", TagVersion),
		fmt.Sprintf("FIRECRACKER_VERSION=%s", FirecrackerVersion),
		fmt.Sprintf("CONTAINERD_VERSION=%s", ContainerdVersion),
	}
}
