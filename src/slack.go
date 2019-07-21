package main

import (
	"github.com/nlopes/slack"
	"os"
	"time"
)

func newSlackClient() *slack.Client {
	token := os.Getenv("SLACK_API_TOKEN")
	return slack.New(token)
}

func postChannel() (channel string) {
	channel = os.Getenv("NOTIFY_CHANNEL")
	if channel == "" {
		channel = "#github"
	}
	return
}

func createMessagesFrom(trends []GithubTrend) (messages []slack.MsgOption) {
	attachments := []slack.Attachment{}
	for _, trend := range trends {
		attachments = append(attachments, trend.toAttachment())
	}
	for i := 0; i < len(attachments)/5; i++ {
		indexFrom := i * 5
		indexTo := (i + 1) * 5
		if indexTo > len(attachments) {
			indexTo = len(attachments) + 1
		}
		message := slack.MsgOptionAttachments(attachments[indexFrom:indexTo]...)
		messages = append(messages, message)
	}
	return
}

func threadText() string {
	now := time.Now().Format("2006/01/02 15:04:05")
	return "*【Github Trending】* at *" + now + "*"
}

func Send(trends []GithubTrend) {
	client := newSlackClient()
	channel := postChannel()
	message := slack.MsgOptionText(threadText(), true)
	_, thread, _ := client.PostMessage(channel, message)
	for _, message = range createMessagesFrom(trends) {
		client.PostMessage(channel, message, slack.MsgOptionTS(thread))
	}
}
