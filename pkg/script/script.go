package script

import (
	"context"
	"fmt"
	"io"
	"microstack/internal/config"
	"microstack/pkg/logs"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/pkg/errors"
)

type Type string

const (
	InstallPrerequisites   Type = "install-prerequisites.sh"
	UninstallPrerequisites Type = "uninstall-prerequisites.sh"
)

var (
	downloadScriptEndpointFormat = fmt.Sprintf(
		"https://raw.githubusercontent.com/%s/%s/scripts/%%s",
		config.RepositoryName,
		config.GetTagVersionForDownloadScript(config.TagVersion),
	)
)

func LocalScriptFile(version string, t Type) string {
	return path.Join(config.BinDir, version, string(t))
}

func RemoteScriptUrl(script Type) string {
	return fmt.Sprintf(downloadScriptEndpointFormat, script)
}

func Download(script Type, version string, force bool) error {
	log := logs.GetLogger().WithField("version", version).WithField("force", force)

	url := RemoteScriptUrl(script)
	destFile := LocalScriptFile(version, script)

	log.Infof("downloading %s to save %s", url, destFile)

	err := downloadScript(
		url,
		destFile,
		force,
	)
	if err != nil {
		return errors.WithMessagef(err, "failed to download script (%s)", script)
	}

	return nil
}

func downloadScript(url string, destFile string, force bool) error {
	if _, err := os.Stat(destFile); !os.IsNotExist(err) {
		if !force {
			return nil
		}

		if err := os.RemoveAll(destFile); err != nil {
			return errors.WithStack(err)
		}
	}

	if err := os.MkdirAll(filepath.Dir(destFile), 0755); err != nil && err != os.ErrExist {
		return errors.WithStack(err)
	}

	resp, err := http.Get(url)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()

	out, err := os.Create(destFile)
	if err != nil {
		return errors.WithStack(err)
	}
	defer out.Close()

	if err := out.Chmod(0700); err != nil {
		return errors.WithStack(err)
	}

	if _, err := io.Copy(out, resp.Body); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func runScript(script string, beforeCallback func(cmd *exec.Cmd) error) error {
	if _, err := os.Stat(script); os.IsNotExist(err) {
		return errors.WithStack(err)
	}

	cmd := exec.CommandContext(context.Background(), "sudo", "-E", script)

	if beforeCallback != nil {
		if err := beforeCallback(cmd); err != nil {
			return errors.WithStack(err)
		}
	}

	if err := cmd.Run(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func Run(script Type, version string, beforeCallback func(cmd *exec.Cmd) error) error {
	log := logs.GetLogger().WithField("version", version)
	log.Infof("running script (%s)", script)

	f := LocalScriptFile(version, script)

	log.Infof("running %s", f)
	err := runScript(f, beforeCallback)

	if err != nil {
		return errors.WithMessagef(err, "failed to run script (%s)", script)
	}

	return nil
}
