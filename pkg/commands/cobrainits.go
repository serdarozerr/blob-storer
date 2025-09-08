package commands

import cobra "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "Blob storage cli",
	Short: "A cli based blob storage application.",
	Long:  `This application is a simple blob storage service that allows users to store and retrieve binary large objects (blobs) via a command-line interface.`,
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() {

	// RootCmd.AddCommand(uploadCmd)
	RootCmd.AddCommand(InitDownload())
	// RootCmd.AddCommand(deleteCmd)
	// RootCmd.AddCommand(listCmd)
}
