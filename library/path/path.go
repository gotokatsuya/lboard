package path

import (
	"os"
	"path"
)

// GetAppPath return the source root for app.
func GetAppPath() string {
	return path.Join(os.Getenv("LBOARD_ROOT"))
}
