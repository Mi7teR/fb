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
		fmt.Printf("Branch %s not found in %s\n", branchName, repoPath)
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

	err := filepath.Walk(reposDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error proccessing path %s: %w", path, err)
		}
		if info.IsDir() && isGitRepo(path) {
			checkBranchInRepo(path, branchName)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
