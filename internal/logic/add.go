package logic

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iton0/hkup/internal/git"
	"github.com/iton0/hkup/internal/util"
	"github.com/spf13/cobra"
)

var (
	// flags
	Lang string
)

func Add(cmd *cobra.Command, args []string) error {
	var sheBangLine = "#!/bin/sh\n\n\n\n\n" // Defaults to bash
	hook := args[0]

	// check for lang flag and update sheBangLine accordingly
	if Lang != "" {
		if _, err := git.GetLang(Lang); err != nil {
			return err
		}
		sheBangLine = fmt.Sprintf("#!/usr/bin/env %s\n\n\n\n\n", Lang)
	}

	// check if the hkup folder exists and make the file in the .hkup directory
	if !util.DoesDirectoryExist(FullPath) {
		return fmt.Errorf("failed running \"hkup add\"\n%s does not exist", FullPath)
	}

	filePath := filepath.Join(FullPath, hook)

	// check if filePath already exists; dont want to overwrite
	if util.DoesFileExist(filePath) {
		return fmt.Errorf("%s hook already exists", hook)
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// Write the content to the file
	_, err = file.WriteString(sheBangLine)
	if err != nil {
		return fmt.Errorf("failed writing to file: %w", err)
	}

	// Set the executable permissions
	err = os.Chmod(filePath, 0755)
	if err != nil {
		return fmt.Errorf("failed changing permissions of file: %w", err)
	}

	return nil
}