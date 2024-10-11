package cmd

import (
	"bufio"
	"fluxfile/tools"
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [identifier]",
	Short: "Edit details of a specific project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		identifier := args[0]
		reader := bufio.NewReader(os.Stdin)

		config, err := tools.LoadConfig()
		if err != nil {
			fmt.Println("Error loading configuration:", err)
			return
		}

		project, index := tools.FindProjectByIdentifier(config, identifier)
		if index == -1 {
			fmt.Println("Error: Project not found.")
			return
		}

		fmt.Printf("Current directory for '%s': %s\n", identifier, project.Directory)
		fmt.Print("Enter new directory (or press Enter to keep the current): ")
		newDir, _ := reader.ReadString('\n')
		newDir = tools.SanitizeInput(newDir)
		if newDir != "" {
		if _, err := os.Stat(newDir); os.IsNotExist(err) {
			fmt.Println("Error: Directory does not exist. Please provide a valid path.")
			return
		}
		if newDir != "" {
			project.Directory = newDir
		}
	}

		fmt.Printf("Current GitHub URL for '%s': %s\n", identifier, project.GitURL)
		fmt.Print("Enter new GitHub URL (or press Enter to keep the current): ")
		newGitURL, _ := reader.ReadString('\n')
		newGitURL = tools.SanitizeInput(newGitURL)
		if newGitURL != "" {
            sshEnabled, err := tools.CheckSSHEnabled(config)
            if err != nil {
                fmt.Print("Do you want to enable SSH for GitHub URLs? (yes/no, default: yes): ")
                sshAnswer, _ := reader.ReadString('\n')
                sshAnswer = tools.SanitizeInput(sshAnswer)
                if sshAnswer == "" || sshAnswer == "yes" {
                    sshEnabled = true
                } else {
                    sshEnabled = false
                }
                tools.UpdateSSHEnabled(config, sshEnabled)
            }

            if sshEnabled {
                newGitURL = tools.ToSSHFormat(newGitURL)
            }
            project.GitURL = newGitURL
        }

		fmt.Printf("Current Branch Name for '%s': %s\n", identifier, project.BranchName)
		fmt.Print("Enter new Branch Name (or press Enter to keep the current): ")
		newBranchName, _ := reader.ReadString('\n')
		newBranchName = tools.SanitizeInput(newBranchName)
		if newBranchName != "" {
			project.BranchName = newBranchName
		}

		config.Projects[index] = *project
		err = tools.SaveConfig(config)
		if err != nil {
			fmt.Println("Error saving project configuration:", err)
			return
		}

		fmt.Printf("Project '%s' has been updated successfully.\n", identifier)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
