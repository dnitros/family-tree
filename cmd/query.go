package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query to find a relative",
	Long: `Query to find a relative
For example:
	family-tree query father of John Doe
	family-tree query son of Mary Ann`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			if err := cmd.Help(); err != nil {
				cobra.CheckErr(err)
			}
			fmt.Println("\nNo arguments passed")
			fmt.Println("family-tree query <relationship> of <name>")
			return
		}
		findRelative(args)
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}

func findRelative(input []string) {
	connectionString := strings.Join(input, " ")
	splits := strings.Split(connectionString, " of ")
	key := fmt.Sprintf("tree.connect.%s.%s", splits[0], strings.ToLower(splits[1]))
	sons := viper.GetStringSlice(key)
	for _, son := range sons {
		fmt.Println(son)
	}
}
