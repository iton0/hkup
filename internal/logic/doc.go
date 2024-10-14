package logic

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/iton0/hkup/internal/git"
	"github.com/spf13/cobra"
)

func Doc(cmd *cobra.Command, args []string) error {
	key := args[0]
	var termCmd *exec.Cmd
	var url string

	if hook, err := git.GetHook(key); err != nil {
		return err
	} else {
		url = git.HookDocSite + hook
	}

	switch runtime.GOOS {
	case "linux":
		termCmd = exec.Command("xdg-open", url)
	case "darwin":
		termCmd = exec.Command("open", url)
	case "windows":
		termCmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	termCmd.Start()

	return termCmd.Wait()
}
