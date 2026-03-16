package cmd

import (
	"github.com/spf13/cobra"
)

var labCmd = &cobra.Command{
	Use:   "lab",
	Short: "Start Floyd as an MCP server",
	Run: func(cmd *cobra.Command, args []string) {
		// This leverages the existing 'ai' command logic to trigger MCP
		aiCmd.Run(cmd, []string{"mcp"})
	},
}

func init() {
	rootCmd.AddCommand(labCmd)
}
