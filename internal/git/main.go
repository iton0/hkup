package git

import (
	"fmt"
)

const HookDocSite = "https://git-scm.com/docs/githooks#"

var (
	// source: https://git-scm.com/docs/githooks
	// INFO: all supported hooks as of 10/24/2024
	// NOTE: update above date when changing hooks var
	//
	// key = git hook name
	// value = url section of online git hooks docs
	hooks = map[string]string{
		"applypatch-msg":        "_applypatch_msg",
		"pre-applypatch":        "_pre_applypatch",
		"post-applypatch":       "_post_applypatch",
		"pre-commit":            "_pre_commit",
		"pre-merge-commit":      "_pre_merge_commit",
		"prepare-commit-msg":    "_prepare_commit_msg",
		"commit-msg":            "_commit_msg",
		"post-commit":           "_post_commit",
		"pre-rebase":            "_pre_rebase",
		"post-checkout":         "_post_checkout",
		"post-merge":            "_post_merge",
		"pre-push":              "_pre_push",
		"pre-receive":           "pre-receive",
		"update":                "update",
		"proc-receive":          "proc-receive",
		"post-receive":          "post-receive",
		"post-update":           "post-update",
		"reference-transaction": "_reference_transaction",
		"push-to-checkout":      "_push_to_checkout",
		"pre-auto-gc":           "_pre_auto_gc",
		"post-rewrite":          "_post_rewrite",
		"sendemail-validate":    "_sendemail_validate",
		"fsmonitor-watchman":    "_fsmonitor_watchman",
		"p4-changelist":         "_p4_changelist",
		"p4-prepare-changelist": "_p4_prepare_changelist",
		"p4-post-changelist":    "_p4_post_changelist",
		"p4-pre-submit":         "_p4_pre_submit",
		"post-index-change":     "_post_index_change",
	}

	// Languages that can be used for git hooks (excluding default bash)
	// NOTE: any language can technically be used but
	// this is those that do not need to be compiled.
	supportedLangs = map[string]bool{
		"python": true,
		"ruby":   true,
		"node":   true,
		"perl":   true,
		"php":    true,
	}
)

// GetHook retrieves the value of a specified git hook by its key.
// Returns the hook value and an error if the hook is not found.
func GetHook(key string) (string, error) {
	value, exists := hooks[key]
	if !exists {
		return "", fmt.Errorf("hook not found: %s", key)
	}
	return value, nil
}

// Hooks returns the map of all defined git hooks.
func Hooks() map[string]string {
	return hooks
}

// GetLang checks if a specified language is supported for git hooks.
// Returns a boolean indicating support and an error if the language is not found.
func GetLang(key string) (bool, error) {
	value, exists := supportedLangs[key]
	if !exists {
		return false, fmt.Errorf("language not supported: %s", key)
	}
	return value, nil
}

// SupportedLangs returns the map of supported git hook languages.
func SupportedLangs() map[string]bool {
	return supportedLangs
}
