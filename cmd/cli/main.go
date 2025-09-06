package main

import cobra "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "MyCLI is a sample command-line application",
	Long:  `MyCLI is a sample command-line application built with Cobra in Go.`,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
