package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.SetOut(os.Stdout)
}

var version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of house-facts",
	Long:  "Print the version number of house-facts",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("house-facts version: %s\n", version)
	},
}
