/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "ttalk",
		Short: "A brief description",
		Long:  `A longer description`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) {
			return initializeConfig(cmd)
		},
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ttalk.yaml)")
}

func initializeConfig(cmd *cobra.Command) error {
	viper.SetEnvPrefix("TTALK")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "*", "-", "*"))
	viper.AutomaticEnv()
	if cfgFile != nil {
	}
	return nil
}
