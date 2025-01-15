package main

import (
	"fmt"
	"log"
	"strings"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Println("git-tools version 1.0.0")
		return
	}

	currentBranch := executeCommand("git", "rev-parse", "--abbrev-ref", "HEAD")
	branchParts := strings.Split(currentBranch, "/")

	if len(branchParts) < 3 {
		log.Fatalf("Invalid branch format: %s", currentBranch)
	}

	branchType := branchParts[0]
	version := branchParts[1]
	jiraID := branchParts[2]
	commitMessage := executeCommand("git", "log", "-1", "--pretty=%B")
	title := commitMessage
	targetBranch := fmt.Sprintf("dev/%s", version)

	if !strings.Contains(commitMessage, jiraID) {
		title = fmt.Sprintf("[%s]%s", jiraID, commitMessage)
	}

	// Check for valid branch types
	validBranchTypes := map[string]bool{
		"feature": true, "feat": true, "bugfix": true,
		"hotfix": true, "task": true, "subtask": true, "subTask": true,
	}

	if !validBranchTypes[branchType] {
		result := executeCommand("git", "push", "--set-upstream", "origin", currentBranch)
		fmt.Println(result)
		return
	}

	// Handle task or subtask branches
	if branchType == "task" || branchType == "subtask" || branchType == "subTask" {
		targetBranch = findParentBranch()
	}
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Jira ID: %s\n", jiraID)
	fmt.Printf("Commit message: %s\n", commitMessage)
	fmt.Printf("Title: %s\n", title)
	fmt.Printf("\033[34mTarget branch: %s\033[0m\n", targetBranch)
	// if !installGlabIfNeeded() {
	// 	log.Fatal("Cannot continue without glab.")
	// }

	// 通过glab 命令检查是否有mr存在，如果存在，则不再创建新的mr
	// mrExists := executeCommand("glab", "mr", "list", "--source-branch=" + currentBranch, "--target-branch=" + targetBranch)
	// result := ""
// 	if (mrExists!= "") {
// 		result = executeCommand("git", "push", "--set-upstream", "origin", currentBranch)
// 		fmt.Println(result)
// 		return
// 	} else {
// 		// Push and create MR
// 		result = executeCommand("git", "push", "--set-upstream", "origin", currentBranch,
// 			"-o", "merge_request.create",
// 			"-o", fmt.Sprintf("merge_request.title=%s", title),
// 			"-o", fmt.Sprintf("merge_request.target=%s", targetBranch),
// )
// 	}

	result := executeCommand("git", "push", "--set-upstream", "origin", currentBranch,
		"-o", "merge_request.create",
		"-o", fmt.Sprintf("merge_request.title=%s", title),
		"-o", fmt.Sprintf("merge_request.target=%s", targetBranch),
	)
	fmt.Println(result)
}
