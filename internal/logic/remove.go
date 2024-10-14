package logic

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iton0/hkup/internal/util"
	"github.com/spf13/cobra"
)

func Remove(cmd *cobra.Command, args []string) error {
	hook := args[0]

	// check if the hkup folder exists and make the file in the .hkup directory
	if !util.DoesDirectoryExist(FullPath) {
		return fmt.Errorf("failed running \"hkup remove\"\n%s folder does not exist", FullPath)
	}

	filePath := filepath.Join(FullPath, hook)

	if !util.DoesFileExist(filePath) {
		return fmt.Errorf("not supported hook: %s", hook)
	}

	// remove the file
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed deleting file: %w", err)
	}

	return nil
}
