package tools

import (
    "os"
    "os/exec"
    "time"
    "fmt"
    "strings"
)

func SyncProject(project *Project) {
    if err := os.Chdir(project.Directory); err != nil {
        fmt.Println("Error changing directory:", err)
        return
    }

    cmdGetURL := exec.Command("git", "remote", "get-url", "origin")
    currentURL, err := cmdGetURL.CombinedOutput()
    if err != nil {
        fmt.Printf("Error getting current remote URL: %s\nOutput: %s\n", err, string(currentURL))
        return
    }

    cmdSetURL := exec.Command("git", "remote", "set-url", "origin", project.GitURL)
    if output, err := cmdSetURL.CombinedOutput(); err != nil {
        fmt.Printf("Error setting remote URL: %s\nOutput: %s\n", err, string(output))
        return
    }

    defer func() {
        // Restore the original remote URL
        cmdRestoreURL := exec.Command("git", "remote", "set-url", "origin", strings.TrimSpace(string(currentURL)))
        if output, err := cmdRestoreURL.CombinedOutput(); err != nil {
            fmt.Printf("Error restoring original remote URL: %s\nOutput: %s\n", err, string(output))
        }
    }()


    cmdStatus := exec.Command("git", "status", "--porcelain")
    statusOutput, err := cmdStatus.CombinedOutput()
    if err != nil {
        fmt.Printf("Error checking git status: %s\nOutput: %s\n", err, string(statusOutput))
        return
    }

    if len(statusOutput) > 0 {
        cmdAdd := exec.Command("git", "add", ".")
        if output, err := cmdAdd.CombinedOutput(); err != nil {
            fmt.Printf("Error staging changes: %s\nOutput: %s\n", err, string(output))
            return
        }

        cmdDiffCached := exec.Command("git", "diff", "--cached", "--quiet")
        if err := cmdDiffCached.Run(); err == nil {
            fmt.Println("No changes to commit.")
        } else {
            commitMessage := time.Now().Format("2006-01-02 15:04:05")
            cmdCommit := exec.Command("git", "commit", "-m", commitMessage)
            if output, err := cmdCommit.CombinedOutput(); err != nil {
                fmt.Printf("Error committing changes: %s\nOutput: %s\n", err, string(output))
                return
            }
        }
    }

    cmdLog := exec.Command("git", "log", "origin/"+project.BranchName+"..HEAD")
    logOutput, err := cmdLog.CombinedOutput()
    if err != nil {
        fmt.Printf("Error checking for commits to push: %s\nOutput: %s\n", err, string(logOutput))
        return
    }
    if len(logOutput) == 0 {
        fmt.Println("No changes to sync.")
        return
    }

    cmdPull := exec.Command("git", "pull", "--rebase", "origin", project.BranchName)
    output, err := cmdPull.CombinedOutput()
    if err != nil {
        if strings.Contains(string(output), "CONFLICT") {
            fmt.Println("Merge conflicts detected. Please resolve them manually.")
        } else {
            fmt.Printf("Error pulling changes from upstream: %s\nOutput: %s\n", err, string(output))
        }
        return
    }

    cmdPush := exec.Command("git", "push", "origin", project.BranchName)
    if output, err := cmdPush.CombinedOutput(); err != nil {
        fmt.Printf("Error pushing changes: %s\nOutput: %s\n", err, string(output))
        return
    }

    fmt.Printf("Successfully synced the local directory with '%s' branch of '%s'\n", project.BranchName, project.GitURL)
}