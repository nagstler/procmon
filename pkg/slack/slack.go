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
	fmt.Println("Message successfully sent on Slack:")
	fmt.Println("\tChannel:", channelID)
	fmt.Println("\tTime:", timestamp)
}
