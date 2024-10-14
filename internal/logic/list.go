package logic

import (
	"github.com/iton0/hkup/internal/git"
	"github.com/iton0/hkup/internal/util"
	"github.com/spf13/cobra"
)

func List(cmd *cobra.Command, args []string) error {
	arg := args[0]
	var output []string

	// NOTE: default case is handled by cobra framework
	switch {
	case arg == "hook":
		output = util.ConvertMapKeysToSlice(git.Hooks())
	case arg == "lang":
		output = util.ConvertMapKeysToSlice(git.SupportedLangs())
	}

	for _, key := range output {
		cmd.Printf(" %s\n", key)
	}

	return nil
}
