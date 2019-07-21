package main

import (
	"github.com/nlopes/slack"
	"os"
)

func newSlackClient() *slack.Client {
	token := os.Getenv("SLACK_API_TOKEN")
	return slack.New(token)
}

func Send() {
	client := newSlackClient()
	test_message := slack.MsgOptionText("test", true)
	client.PostMessage("#github", test_message)
}
