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
	Use:   "procmon",
	Short: "ProcMon is a process monitoring tool",
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}

		slackToken := viper.GetString("slack.token")
		if slackToken == "" {
			log.Fatal("Slack token is not set in the configuration file")
		}

		channelID := viper.GetString("slack.channel")
		if channelID == "" {
			log.Fatal("Slack channel ID is not set in the configuration file")
		}

		processNames := viper.GetStringSlice("processes")
		if len(processNames) == 0 {
			log.Fatal("No processes are specified in the configuration file")
		}

		log.Info("ProcMon started")

		monitor.Start(processNames, slackToken, channelID)

		select {}
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
		fmt.Println("PROCMON STARTED:")
		fmt.Println("\tUsing config file:", viper.ConfigFileUsed())
	}
}
