package opener

import (
	"errors"
	"os/exec"
	"runtime"
)

func OpenURL(url string) error {
	switch runtime.GOOS {

	case "linux":
		// Linux uses xdg-open for opening any URI or URL
		return exec.Command("xdg-open", url).Start()

	case "darwin":
		// macOS uses "open" command to open files, apps, and URLs
		return exec.Command("open", url).Start()

	case "windows":
		// Windows uses "rundll32" to open URLs with the default handler
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()

	default:
		return errors.New("unsupported operating system")
	}
}
