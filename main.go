package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func checkBranchInRepo(repoPath, branchName string) {
	cmd := exec.Command("git", "show-ref", "--verify", fmt.Sprintf("refs/heads/%s", branchName))
	cmd.Dir = repoPath

	if err := cmd.Run(); err != nil {
		return
	}

	fmt.Printf("Branch %s found in %s\n", branchName, repoPath)
}

func isGitRepo(path string) bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func main() {
	dirP := flag.String("d", ".", "Path to directory with repositories")
	branchP := flag.String("b", "main", "Branch name")

	flag.Parse()

	reposDir := *dirP
	branchName := *branchP

	dirs, err := os.ReadDir(reposDir)
	if err != nil {
		fmt.Printf("Error reading directory %s: %s\n", reposDir, err)
		os.Exit(1)
	}

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		repoPath := filepath.Join(reposDir, dir.Name())

		if !isGitRepo(repoPath) {
			continue
		}

		checkBranchInRepo(repoPath, branchName)
	}
}
