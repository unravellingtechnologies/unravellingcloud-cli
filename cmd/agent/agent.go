/*
Copyright © 2022 Unravelling Technologies GmbH <unravelling@unravelling.io>
*/
package agent

import (
	"github.com/spf13/cobra"
)

// Variables used for flags.
var terraformEnabled bool
var workingDir string
var host string
var port int

// AgentCmd represents the agent command
var AgentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Use the CLI to start an agent",
	Long:  `The agent command allows to start an agent locally. This agent can then be accessed via GraphQL.`,
}

func init() {
	AgentCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agentCmd.PersistentFlags().BoolP("terraformEnabled", "t", true, "Whether this agent should support running Terraform actions locally")
	AgentCmd.PersistentFlags().StringVarP(&host, "host", "H", "localhost", "The host to listen on")
	AgentCmd.PersistentFlags().IntVarP(&port, "port", "P", 8080, "The port to listen on")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	AgentCmd.Flags().BoolVarP(&terraformEnabled, "terraform", "t", true, "Whether this agent should support running Terraform actions locally")
	AgentCmd.Flags().StringVarP(&workingDir, "workingDir", "w", "", "Working directory for the Agent")
}
