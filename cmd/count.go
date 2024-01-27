package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count connected relationships of a person",
	Long: `Count connected relationships of a person
For example:
	family-tree count sons of Mary Ann
	family-tree count daughters of John Doe
	family-tree count wives of John Doe`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 || args[1] != "of" {
			if err := cmd.Help(); err != nil {
				cobra.CheckErr(err)
			}
			fmt.Println("Invalid arguments for count passed")
			fmt.Println("Mandatory usage:")
			fmt.Println("family-tree count <relationship in plural> of <person>")
			return
		}

		relationship, person := extractCountArgs(args)
		relationship = singularRelationship(relationship)
		key := "tree.connect." + relationship + "." + strings.ToLower(person)
		fmt.Println(len(viper.GetStringSlice(key)))
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
}

func extractCountArgs(input []string) (relationship, person string) {
	connectionString := strings.Join(input, " ")
	splits := strings.Split(connectionString, " of ")
	relationship = splits[0]
	person = splits[1]
	return
}

func singularRelationship(input string) string {
	switch {
	case strings.HasSuffix(input, "ies"):
		return strings.TrimSuffix(input, "ies") + "y"
	case strings.HasSuffix(input, "ren"):
		return strings.TrimSuffix(input, "ren")
	case strings.HasSuffix(input, "wives"):
		return strings.TrimSuffix(input, "ves") + "fe"
	default:
		return strings.TrimSuffix(input, "s")
	}
}
