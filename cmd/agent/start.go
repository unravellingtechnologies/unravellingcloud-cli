/*
Copyright © 2023 Unravelling Technologies GmbH <unravelling@unravelling.io>
*/
package agent

import (
	"github.com/spf13/cobra"
	"github.com/unravellingtechnologies/unravelling-cloud/cli/lib/server"
	"github.com/unravellingtechnologies/unravelling-cloud/cli/lib/terraform"
)

// StartCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(startCmd *cobra.Command, args []string) {
		server.StartServer(&server.ServerOpts{Port: port, Host: host})
	},
}

func init() {
	server.Mount(server.Router, "/api/terraform", terraform.Router)
}
