package main

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	noGitErr     = errors.New("a `git` executable was not found in $PATH")
	dirExistsErr = errors.New("cannot create this directory because it already exists")
	copyEnvErr   = errors.New("unable to create a .env file automatically")
	buildErr     = errors.New("unable to build an app binary")
	unknownErr   = errors.New("something went wrong")
)

var new = &cobra.Command{
	Use:     "new",
	Short:   "Create a new Gostalt application",
	Example: "gostalt new app_name",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkGit(); err != nil {
			return err
		}

		dir := args[0]

		if err := clone(dir); err != nil {
			return err
		}

		if err := copyEnv(dir); err != nil {
			return err
		}

		if err := build(dir); err != nil {
			return err
		}

		if err := printIntro(dir); err != nil {
			return err
		}

		return nil
	},
}

// checkGit determines whether a git binary exists in the current system's $PATH.
// If it doesn't, installation can't proceed as git is needed to clone the repo.
//
// @see https://github.com/gostalt/cli/issues/1
func checkGit() error {
	_, err := exec.Command("git", "--version").Output()
	if err != nil {
		return noGitErr
	}

	return nil
}

func clone(dir string) error {
	fmt.Printf("⬇️  Cloning Gostalt to `%s` ... \n", dir)
	if _, err := exec.Command("git", "clone", "https://github.com/gostalt/gostalt.git", dir).Output(); err != nil {
		code, _ := err.(*exec.ExitError)
		switch code.ExitCode() {
		case 128:
			return dirExistsErr
		default:
			return unknownErr
		}
	}

	exec.Command("rm", "-rf", dir+"/.git").Output()
	return nil
}

func copyEnv(dir string) error {
	fmt.Printf("⚙️  Copying %s/.env.example file to %s/.env ... \n", dir, dir)
	if _, err := exec.Command("cp", dir+"/.env.example", dir+"/.env").Output(); err != nil {
		return copyEnvErr
	}

	return nil
}

func build(dir string) error {
	fmt.Println("📦  Building app binary ... ")
	cmdd := exec.Command("go", "build")
	cmdd.Dir = dir
	if _, err := cmdd.Output(); err != nil {
		return buildErr
	}

	return nil
}

func printIntro(dir string) error {
	fmt.Println("🚀 To get started:")
	fmt.Println("    cd " + dir)
	fmt.Println("    ./gostalt serve")

	return nil
}
