package main

import (
	"github.com/nlopes/slack"
)

const BLUE = "#4183c4"

type GithubTrend struct {
	rank int
	repo string
	desc string
}

func (trend *GithubTrend) url() string {
	return "https://github.com" + trend.repo
}

func (trend *GithubTrend) toAttachment() (attachment slack.Attachment) {
	return slack.Attachment{
		Color:     BLUE,
		Title:     trend.repo,
		TitleLink: trend.url(),
		Footer:    "Github",
	}
}
