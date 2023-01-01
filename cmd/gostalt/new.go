package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var new = &cobra.Command{
	Use:     newCmd,
	Short:   newDescr,
	Example: newExample,
	Args:    cobra.ExactArgs(1),
	Run:     newFunc,
}

func hasGitInstalled() bool {
	// TODO: Possibly look at downloading the repo as a fallback.
	_, err := exec.Command("git", "--version").Output()
	if err != nil {
		return false
	}

	return true
}

func cloneGitRepoToDir(dir string) {
	fmt.Printf("⬇️  Cloning Gostalt to `%s` ... ", dir)
	_, err := exec.Command("git", "clone", "https://github.com/gostalt/gostalt.git", dir).Output()
	code, _ := err.(*exec.ExitError)

	if code != nil {
		fmt.Println("❌")
		switch code.ExitCode() {
		case 128:
			fmt.Println("    - " + dirExistsErr)
		}
		os.Exit(1)
	}

	fmt.Println("✅")

	exec.Command("rm", "-rf", dir+"/.git").Output()
}

func copyEnvFile(dir string) {
	fmt.Printf("⚙️  Copying %s/.env.example file to %s/.env ... ", dir, dir)
	_, err := exec.Command("cp", dir+"/.env.example", dir+"/.env").Output()
	if err != nil {
		fmt.Println("❌")
		fmt.Println("    - Gostalt was unable to create a .env file automatically")
		fmt.Printf("    - To fix this, copy the .env.example file inside the %s directory\n", dir)
		return
	}

	fmt.Println("✅")
}

func buildInitialBinary(dir string) {
	fmt.Print("📦  Building app binary ... ")
	cmdd := exec.Command("go", "build")
	cmdd.Dir = dir
	_, err := cmdd.Output()
	if err != nil {
		fmt.Println("❌")
		fmt.Println("    - Gostalt was unable to build an app binary")
		fmt.Printf("    - To fix this, run `go build` from the %s directory\n", dir)
		return
	}

	fmt.Println("✅")
}

func printGetStarted(dir string) {
	fmt.Println("🚀 To get started:")
	fmt.Println("    cd " + dir)
	fmt.Println("    ./gostalt serve")
}

func newFunc(cmd *cobra.Command, args []string) {
	dir := args[0]
	if !hasGitInstalled() {
		fmt.Println("⚡️ " + noGitErr)
		os.Exit(1)
	}

	cloneGitRepoToDir(dir)

	copyEnvFile(dir)

	buildInitialBinary(dir)

	printGetStarted(dir)
}

const (
	noGitErr     string = "A `git` executable was not found in your $PATH. Ensure git is installed on your system."
	dirExistsErr string = "Gostalt cannot create this directory because it already exists."
	unknownErr   string = "Unable to create a new application. An unknown error occurred."
)

const (
	newCmd     = "new"
	newDescr   = "Create a new Gostalt application"
	newExample = "gostalt new app_name"
)
