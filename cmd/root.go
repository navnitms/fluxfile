package cmd

import (
    "github.com/spf13/cobra"
    "fmt"
)

var rootCmd = &cobra.Command{
    Use:   "fluxfile",
    Short: "A tool to sync files with Git and other sources",
    Long:  "fluxfile is a command-line application to sync files of local-first applications with Git and other locations",
}

func Execute() {
    rootCmd.CompletionOptions.DisableDefaultCmd = true
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        return
    }
}
