package cmd

import (
	"github.com/iton0/hkup/internal/git"
	"github.com/iton0/hkup/internal/logic"
	"github.com/iton0/hkup/internal/util"
	"github.com/spf13/cobra"
)

var (
	docCmd = &cobra.Command{
		Use:       "doc <hook-name>",
		Aliases:   []string{"docs"},
		Short:     "Documentation for git hook",
		ValidArgs: util.ConvertMapKeysToSlice(git.Hooks()),
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE:      logic.Doc,
	}
)

func init() {
	rootCmd.AddCommand(docCmd)
}