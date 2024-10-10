package cmd

import (
	"bufio"
	"fluxfile/tools"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a directory for syncing",
	Long:  `This command helps you initialize a local directory to sync with a Git repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		config, err := tools.LoadConfig()
		if err != nil {
			fmt.Println("Error loading configuration:", err)
			return
		}

		fmt.Print("Enter the path to the folder you want to sync: ")
		dir, _ := reader.ReadString('\n')
		dir = tools.SanitizeInput(dir)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Println("Error: Directory does not exist. Please provide a valid path.")
			return
		}

		fmt.Print("Enter the GitHub repository URL: ")
		gitURL, _ := reader.ReadString('\n')
		gitURL = tools.SanitizeInput(gitURL)
		if valid := tools.CheckForValidUrl(gitURL); !valid {
			fmt.Println("Error: You must provide a valid GitHub URL.")
			return
		}
		if gitURL == "" {
			fmt.Println("Error: You must provide a valid GitHub URL.")
			return
		}

		fmt.Print("Enter the branch name: ")
		branchName, _ := reader.ReadString('\n')
		branchName = tools.SanitizeInput(branchName)
		if branchName == "" {
			fmt.Println("Error: You must provide a valid Branch Name.")
			return
		}

		fmt.Print("Enter an identifier (name) for this project: ")
		projectName, _ := reader.ReadString('\n')
		projectName = tools.SanitizeInput(projectName)
		if projectName == "" {
			fmt.Println("Error: You must provide a project name.")
			return
		}

		if _, index := tools.FindProjectByIdentifier(config, projectName); index != -1 {
			fmt.Println("Error: Project with this identifier already exists. Please choose a different identifier.")
			return
		}

		project := tools.Project{
			Directory:   dir,
			GitURL:      gitURL,
			ProjectName: projectName,
			BranchName:  branchName,
		}

		config.Projects = append(config.Projects, project)
		err = tools.SaveConfig(config)
		if err != nil {
			fmt.Println("Error saving project configuration:", err)
			return
		}

		fmt.Printf("Successfully initialized the folder '%s' for syncing with '%s' under project name '%s'\n", dir, gitURL, projectName)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
