package cmd

import (
	"fmt"
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

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
	rootCmd.Flags().IntVarP(&pid, "pid", "p", 0, "process id to monitor")
}

func initConfig() {
	viper.SetConfigName("config") // look for config.yaml
	viper.AddConfigPath(".")      // in the current directory

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
