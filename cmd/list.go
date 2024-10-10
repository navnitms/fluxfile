package cmd

import (
	"fluxfile/tools"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := tools.LoadConfig()
		if err != nil {
			fmt.Println("Error loading configuration:", err)
			return
		}

		fmt.Println("Projects:")
		for _, project := range config.Projects {
			fmt.Printf("- %s\n", project.ProjectName)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
