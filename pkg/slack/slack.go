package slack

import (
	"fmt"

	"github.com/slack-go/slack"
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
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
