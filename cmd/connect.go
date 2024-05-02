package cmd

import (
	"fmt"
	"strings"

	"github.com/dnitros/family-tree/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect two persons into a relationship",
	Long: `Connect two persons into a relationship
Mandatory usage of in below format:
	family tree connect <person1> as <relationship> of <person2>
For example:
	family tree connect John Doe as son of Mary Ann.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			if err := cmd.Help(); err != nil {
				cobra.CheckErr(err)
			}
			fmt.Println("\nNo arguments passed")
			fmt.Println("family-tree connect <person1> as <relationship> of <person2>")
			return
		}
		connectionString := strings.Join(args, " ")
		asIndex := strings.Index(connectionString, " as ")
		ofIndex := strings.Index(connectionString, " of ")
		if asIndex == -1 || ofIndex == -1 || asIndex > ofIndex {
			if err := cmd.Help(); err != nil {
				cobra.CheckErr(err)
			}
			fmt.Println("Invalid args for connect passed")
			fmt.Println("Mandatory usage:")
			fmt.Println("family-tree connect <person1> as <relationship> of <person2>")
			return
		}
		leftName, relationship, rightName := extractConnectArgs(connectionString)
		utils.ValueExists("tree.person", leftName, rightName)
		utils.ValueExists("tree.relationship", relationship)
		addConnection(leftName, relationship, rightName)
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}

func extractConnectArgs(connectionString string) (leftName, relationship, rightName string) {
	splits := strings.Split(connectionString, " as ")
	leftName = splits[0]
	splits = strings.Split(splits[1], " of ")
	relationship = splits[0]
	rightName = splits[1]
	return
}

func addConnection(leftName, relationship, rightName string) {
	key := fmt.Sprintf("tree.connect.%s.%s", relationship, strings.ToLower(rightName))
	connects := viper.GetStringSlice(key)
	for _, connect := range connects {
		if leftName == connect {
			return
		}
	}
	viper.Set(key, append(connects, leftName))
	err := viper.WriteConfig()
	cobra.CheckErr(err)
}
