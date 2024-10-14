package logic

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/iton0/hkup/internal/util"
	"github.com/spf13/cobra"
)

var (
	// local repo folder name to hold git hooks via relative path
	FullPath = filepath.Join(".", ".hkup") // NOTE: treat as constant

	// flags
)

func Init(cmd *cobra.Command, args []string) error {
	// check if cwd is git repo
	if err := exec.Command("git", "-C", ".", "rev-parse", "--is-inside-work-tree").Run(); err != nil {
		return fmt.Errorf("failed to check if cwd is git repo: %w", err)
	}

	// check if there is a hkup folder and create if necessary
	if !util.DoesDirectoryExist(FullPath) {
		if err := util.CreateFolder(FullPath); err != nil {
			return err
		}
		cmd.Printf("Initialized hkup folder at %s\n", FullPath)
	}

	if out, _ := exec.Command("git", "config", "--local", "core.hooksPath").Output(); len(out) != 0 {
		return fmt.Errorf("hooksPath already set to %s", out)
	} else {
		// update the local git core.hookPath
		if err := exec.Command("git", "config", "--local", "core.hooksPath", FullPath).Run(); err != nil {
			return fmt.Errorf("failed to set hooksPath: %w", err)
		}
		return nil
	}
}
