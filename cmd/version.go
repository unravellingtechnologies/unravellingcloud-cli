/*
Copyright © 2023 Unravelling Technologies GmbH <unravelling@unravelling.io>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "--version",
	Short: "Print the version number of unravelling cloud's CLI",
	Long:  `All software has versions. This is unravelling cloud's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Unravelling CLI: %s", "1.0.0")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
