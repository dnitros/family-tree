package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a person/relationship",
	Long: `Add a person or a relationship which can be further used for connection.
For example:
	family-tree add person John Doe
	family-tree add relationship father`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			cobra.CheckErr(err)
		}
		fmt.Println("\nfamily-tree: error: subcommand required")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
