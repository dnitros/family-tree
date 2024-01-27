package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "family-tree",
	Short: "Create and query a family tree",
	Long: `family-tree is a cli tool to manage family-tree.
This application is a tool to add member, define relationships and make basic queries`,
}

func init() {
	cobra.OnInitialize(initConfig)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".family-tree")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, err := fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed()); err != nil {
			return
		}
	}
}
