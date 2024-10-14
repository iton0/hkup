package cmd

import (
	"github.com/spf13/cobra"
)

// https://github.com/spf13/cobra/blob/main/site/content/user_guide.md
var (
	rootCmd = &cobra.Command{
		Use:     "hkup",
		Short:   "hkup CLI",
		Long:    `hkup is a management tool for git hooks`,
		Args:    cobra.MinimumNArgs(1),
		Version: "1.0.0",
	}
)

func init() {}

func Execute() {
	rootCmd.Execute()
}
