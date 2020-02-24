package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/brynjarh/xclient/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of xclient",
	Long:  `Print the version number of xclient`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Git Commit:", version.GitCommit)
		fmt.Println("Version:", version.Version)
		fmt.Println("Go Version:", version.GoVersion)
		fmt.Println("OS / Arch:", version.OsArch)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
