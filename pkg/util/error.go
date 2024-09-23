package util

import (
	"context"
	"microstack/pkg/logs"
)

// GenericCheckErr is used by the commands to check if the action failed
// and respond with a fatal error provided by the logger (calls os.Exit)
// Ignite has its own, more detailed implementation of this in cmdutil
func GenericCheckErr(ctx context.Context, err error) {
	switch err.(type) {
	case nil:
		return // Don't fail if there's no error
	}

	logs.GetLogger().Fatal(err)
}
