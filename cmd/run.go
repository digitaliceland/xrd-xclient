package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/brynjarh/xclient/pkg/web"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Launch the X road client and listen on a port",
	Long:  `Launch the X road client and listen on a port`,
	Run: func(cmd *cobra.Command, args []string) {
		addr := "127.0.0.1:5000"
		fmt.Println("Starting a web server on http://", addr)
		web.StartWWW(addr)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
