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
	pid     int
	log     = logrus.New()
)

var rootCmd = &cobra.Command{
	Use:   "procmon",
	Short: "ProcMon is a process monitoring tool",
	Run: func(cmd *cobra.Command, args []string) {
		if pid == 0 {
			log.Fatal("Process ID must be provided")
		}

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

		monitor.Start(int32(pid), slackToken, channelID)
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
	rootCmd.Flags().IntVarP(&pid, "pid", "p", 0, "process id to monitor")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".procmon")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
