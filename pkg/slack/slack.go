package slack

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

var (
	log = logrus.New()
)

func Send(slackToken string, channelID string, message string) {
	api := slack.New(slackToken)

	channelID, timestamp, err := api.PostMessage(
		channelID,
		slack.MsgOptionText(message, false),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	log.Info("Notification sent on Slack channel " + channelID + " at " + timestamp)
}
