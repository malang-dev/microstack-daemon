package config

import (
	"os"
	"path"
)

var (
	HomeDir, _          = os.UserHomeDir()
	RootDir             = path.Join(HomeDir, ".kubefire")
	ClusterRootDir      = path.Join(RootDir, "clusters")
	BinDir              = path.Join(RootDir, "bin")
	BootstrapperRootDir = path.Join(RootDir, "bootstrappers")
)
