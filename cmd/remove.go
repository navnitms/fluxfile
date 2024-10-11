package cmd

import (
    "fmt"
    "fluxfile/tools"

    "github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
    Use:   "remove [identifier]",
    Short: "Remove a specific project",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        identifier := args[0]

        config, err := tools.LoadConfig()
        if err != nil {
            fmt.Println("Error loading configuration:", err)
            return
        }

        err = tools.RemoveProject(config, identifier)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        fmt.Printf("Project '%s' has been removed successfully.\n", identifier)
    },
}

func init() {
    rootCmd.AddCommand(removeCmd)
}
