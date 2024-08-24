package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of busy",
	Long:  `All software has versions. This is busy's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("busy version: v0.0.1-alpha1 -- HEAD")
	},
}
