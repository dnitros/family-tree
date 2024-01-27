package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"slices"
)

func ValueExists(key string, names ...string) {
	values := viper.GetStringSlice(key)
	for _, value := range names {
		if !slices.Contains(values, value) {
			fmt.Println(value + " not found in Database")
		}
	}
}
