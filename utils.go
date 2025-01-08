package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"runtime"
)

func executeCommand(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing command '%s %v': %v\n%s", command, args, err, output)
	}
	return strings.TrimSpace(string(output))
}

func executeCommand2(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(output)), err
}

func findParentBranch() string {
	currentBranch := executeCommand("git", "rev-parse", "--abbrev-ref", "HEAD")
	branchParts := strings.Split(currentBranch, "/")
	if len(branchParts) < 2 {
		log.Fatalf("Invalid branch format: %s", currentBranch)
	}

	version := branchParts[1]
	commitID := executeCommand("git", "rev-parse", "HEAD")
	var parentBranch string

	for i := 0; i < 50; i++ {
		if parentBranch != "" {
			return parentBranch
		}

		// Get parent commit ID
		parentCommitID := executeCommand("git", "rev-parse", commitID+"^")

		// Find branch containing the parent commit
		branches := executeCommand("git", "branch", "--contains", parentCommitID)
		for _, branch := range strings.Split(branches, "\n") {
			branch = strings.TrimSpace(branch)
			if !strings.HasPrefix(branch, "*") && (strings.Contains(branch, fmt.Sprintf("dev/%s", version)) || strings.Contains(branch, fmt.Sprintf("feature/%s", version)) || strings.Contains(branch, fmt.Sprintf("feat/%s", version))) {
				parentBranch = branch
				break
			}
		}
		commitID = parentCommitID
	}

	parentBranch = fmt.Sprintf("dev/%s", version)
	return parentBranch
}

func checkGlabInstalled() bool {
	_, err := executeCommand2("which", "glab")
	return err == nil
}

// installGlab installs glab based on the operating system
func installGlab() error {
	fmt.Println("Installing glab...")
	platform := runtime.GOOS
	switch platform {
	case "darwin":
		_, err := executeCommand2("brew", "install", "glab")
		return err
	case "linux":
		_, err := executeCommand2("sudo", "apt", "update")
		if err != nil {
			return err
		}
		_, err = executeCommand2("sudo", "apt", "install", "-y", "glab")
		return err
	case "windows":
		return fmt.Errorf("automatic installation is not supported on Windows")
	default:
		return fmt.Errorf("unsupported platform: %s", platform)
	}
}

func installGlabIfNeeded() bool {
	if !checkGlabInstalled() {
		fmt.Println("glab is not installed on your system.")
		fmt.Print("Would you like to install it now? (y/N): ")
		var response string
		fmt.Scanln(&response)
		if strings.ToLower(response) == "y" {
			err := installGlab()
			if err != nil {
				fmt.Printf("Failed to install glab: %v\n", err)
				return false // 表示安装失败
			}
			fmt.Println("glab installed successfully!")
			return true // 表示安装成功
		} else {
			fmt.Println("glab is required to run this tool. Please install it and try again.")
			return false
		}
	}
	return true
}
