package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the config file content",
	Long:  `Clear the config file content`,
	Run: func(cmd *cobra.Command, args []string) {
		cfgFile := viper.ConfigFileUsed()
		if err := os.Remove(cfgFile); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _, err := os.OpenFile(cfgFile, os.O_RDONLY|os.O_CREATE, 0666); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
