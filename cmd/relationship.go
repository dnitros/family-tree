package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"

	"github.com/spf13/cobra"
)

var relationshipCmd = &cobra.Command{
	Use:   "relationship",
	Short: "Add a relationship",
	Long: `Add a relationship into the family tree
Currently only supports adding a single relationship at a time.
For example:
	family-tree add relationship father
	family-tree add relationship son`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			if err := cmd.Help(); err != nil {
				cobra.CheckErr(err)
			}
			fmt.Println("\nNo arguments passed")
			fmt.Println("family-tree add relationship <name>")
			return
		}
		relationshipName := strings.ToLower(strings.Join(args, " "))
		addRelationship(relationshipName)
	},
}

func init() {
	addCmd.AddCommand(relationshipCmd)
}

func addRelationship(relationshipName string) {
	relationships := viper.GetStringSlice("tree.relationship")
	for _, relationship := range relationships {
		if relationshipName == relationship {
			return
		}
	}
	viper.Set("tree.relationship", append(relationships, relationshipName))
	err := viper.WriteConfig()
	cobra.CheckErr(err)
}
