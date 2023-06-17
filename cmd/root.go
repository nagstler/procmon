// Importing required packages
package cmd

import (
	"fmt"
	"os"
	"procmon/pkg/monitor"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Variables for command-line flags
	cfgFile     string
	processName string

	// Initializing the logrus logger
	log = logrus.New()
)

// rootCmd defines the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "procmon",
	Short: "ProcMon is a process monitoring tool",
	Run: func(cmd *cobra.Command, args []string) {
		// Validate if process name is provided
		if processName == "" {
			log.Fatal("Process name must be provided")
		}

		// Load the configuration from file
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}

		// Read and validate Slack token from configuration
		slackToken := viper.GetString("slack.token")
		if slackToken == "" {
			log.Fatal("Slack token must be set")
		}

		// Read and validate Slack channel ID from configuration
		channelID := viper.GetString("slack.channel")
		if channelID == "" {
			log.Fatal("Slack channel ID must be set")
		}

		// Start monitoring the process
		monitor.Start(processName, slackToken, channelID)
	},
}

// Execute function adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// init function to initialize command-line flags
func init() {
	cobra.OnInitialize(initConfig)

	// Defining command-line flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.procmon.yaml)")
	rootCmd.Flags().StringVarP(&processName, "name", "n", "", "process name to monitor")
}

// initConfig function reads in config file and environment variables if set.
func initConfig() {
	// Check if config file path is provided in command-line flag
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// If not, set default path to home directory
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".procmon")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
