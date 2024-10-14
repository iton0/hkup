# HkUp
Your CLI tool with benefits built by [iton0](https://github.com/iton0) in [Go](https://go.dev/)!

## Introduction
Git hooks automate and implement processes in your workflow, increasing code quality and consistency.

However, many developers avoid git hooks due to a lack of awareness and the perceived complexity of setup, discouraging them from using this feature.

**HkUp** simplifies the management of git hooks, allowing you to focus on the logic and usage of your hooks instead.

## Installation
Install HkUp from either of the options below:

#### 1. Prebuilt Binary
External Dependencies:
- `git`
- `curl`
- `grep`

Run the script below (supports Linux, macOS, and Windows):

```bash
curl -sSL https://raw.githubusercontent.com/iton0/hkup/main/scripts/install | sh
```
> [!Tip]
> To update HkUp, rerun the above script.
> It will replace the current version.

#### 2. Build from source
Prerequisite: Go 1.23.2 or later

```bash
go install github.com/iton0/hkup@latest
```

## Usage Quickstart
This section provides basic information about core usage. For detailed options run `hkup --help`.

#### Initializing hkup
Run the following command in your git repository to initialize HkUp:
```bash
hkup init
```

This command creates a **.hkup** folder and sets the local **core.hooksPath** variable. If the folder already exists, it will simply update the path variable. The path is relative, ensuring that moving your repository won’t affect hook sourcing.

#### Adding & Removing hooks
Add or remove hooks easily with:
```bash
hkup add <hook-name>
hkup remove <hook-name>
```

#### Info & Docs
There are two commands that will help you with both HkUp and git hooks:

**`hkup list {hook|lang}`**
Outputs list of either available hooks or supported languages.

**`hkup doc <hook-name>`**
Opens your browser with Git documentation for the specified git hook, helping you understand its usage.

## Future TODOs
- [ ] add either flags or subcommand for init to specify dir and worktree; also if you want the hkup folder to be hidden or not
- [ ] functionality to save custom setups (ie gitdir and workdir are not in same location)
- [ ] make an update subcommand
- [ ] store custom git hooks as templates for future use (via add template subcmd)
    - Allow users to create, store, and share templates for common hooks. Users can fetch these templates over the network.
- [ ] branch-specific hooks
- [ ] logo maybe?

## Contributing
Contributing will become available once HkUp moves out of its alpha testing phase.
