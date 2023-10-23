package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// List of interested file extensions
var interestedExtensions = []string{".js", ".jsx", ".ts", ".tsx", ".vue", ".scss", ".css", ".html"}

// getChangedFiles fetches the list of files that have been changed and staged for commit.
func getChangedFiles() []string {
	cmd := exec.Command("git", "diff", "--name-only", "--cached")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error fetching changed files:", err)
		return nil
	}
	return strings.Split(strings.TrimSpace(out.String()), "\n")
}

// filterProjectFiles filters the provided list of files to return only files with interested extensions.
func filterProjectFiles(files []string) []string {
	var projectFiles []string
	for _, file := range files {
		for _, ext := range interestedExtensions {
			if strings.HasSuffix(file, ext) {
				projectFiles = append(projectFiles, file)
				break
			}
		}
	}
	return projectFiles
}

// checkForTODO_CMNT searches for the text "TODO_CMNT:" in each file and prints a message if found.
func checkForTODO_CMNT(files []string) {
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
			continue
		}
		if strings.Contains(string(content), "TODO_CMNT:") {
			fmt.Printf("The text 'TODO_CMNT:' was found in %s\n", file)
		}
	}
}

// runEslint runs the ESLint check on the provided project files.
func runEslint(files []string) {
	eslintCmd := exec.Command("yarn", append([]string{"lint", "--no-fix"}, files...)...)
	eslintCmd.Stdout = os.Stdout
	eslintCmd.Stderr = os.Stderr
	err := eslintCmd.Run()
	if err != nil {
		fmt.Println("Error running ESLint:", err)
	}
}

func main() {
	// Fetch the list of files that have been changed and staged for commit.
	changedFiles := getChangedFiles()
	if changedFiles == nil {
		return
	}

	// Filter out the list to get only files with interested extensions.
	projectFiles := filterProjectFiles(changedFiles)
	if len(projectFiles) == 0 {
		fmt.Printf("No files with extensions %s have been changed.\n", strings.Join(interestedExtensions, ", "))
		return
	}

	// Run the custom check for "TODO_CMNT:" text in the files.
	checkForTODO_CMNT(projectFiles)

	// Run ESLint check on the project files.
	runEslint(projectFiles)
}
