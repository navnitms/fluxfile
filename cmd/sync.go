package cmd

import (
	"fluxfile/tools"
	"fmt"

	"github.com/spf13/cobra"
)

var all bool

var syncCmd = &cobra.Command{
	Use:   "sync [identifier]",
	Short: "Sync the local directory with the specified GitHub repository",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config, err := tools.LoadConfig()
		if err != nil {
			fmt.Println("Error loading configuration:", err)
			return
		}

		if all {
			// Sync all projects
			for _, project := range config.Projects {
				fmt.Println("Fluxing project:", project.ProjectName)
				tools.SyncProject(&project)
			}
		} else {
			if len(args) != 1 {
				fmt.Println("Error: Identifier is required unless --all is specified.")
				return
			}
			identifier := args[0]

			project, index := tools.FindProjectByIdentifier(config, identifier)
			if index == -1 {
				fmt.Println("Error: Project not found.")
				return
			}

			tools.SyncProject(project)
		}
	},
}

func init() {
	syncCmd.Flags().BoolVar(&all, "all", false, "Sync all projects")
	rootCmd.AddCommand(syncCmd)
}
