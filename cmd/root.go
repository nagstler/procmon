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
	cfgFile string
	log     = logrus.New()
)

var rootCmd = &cobra.Command{
	Use:   "procmon [Process Name]",
	Short: "ProcMon is a process monitoring tool",
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is provided (the Process Name)
	Run: func(cmd *cobra.Command, args []string) {
		procName := args[0] // Get the Process Name from the arguments

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}

		slackToken := viper.GetString("slack.token")
		if slackToken == "" {
			log.Fatal("Slack token must be set")
		}

		channelID := viper.GetString("slack.channel")
		if channelID == "" {
			log.Fatal("Slack channel ID must be set")
		}

		monitor.Start(procName, slackToken, channelID)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.procmon.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
