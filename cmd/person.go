package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

var personCmd = &cobra.Command{
	Use:   "person",
	Short: "Add a person",
	Long: `Add a person into the family tree
Currently only supports adding a single person at a time.
For example:
	family-tree add person John Doe
	family-tree add person Mary`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			if err := cmd.Help(); err != nil {
				cobra.CheckErr(err)
			}
			fmt.Println("No arguments passed")
			fmt.Println("family-tree add person <name>")
			return
		}
		personName := strings.Join(args, " ")
		addPerson(personName)
	},
}

func init() {
	addCmd.AddCommand(personCmd)
}

func addPerson(personName string) {
	persons := viper.GetStringSlice("tree.person")
	for _, person := range persons {
		if personName == person {
			return
		}
	}
	viper.Set("tree.person", append(persons, personName))
	err := viper.WriteConfig()
	cobra.CheckErr(err)
}
