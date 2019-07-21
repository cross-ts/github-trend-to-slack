package main

type GithubTrend struct {
	rank int
	repo string
	desc string
}

func (trend *GithubTrend) url() string {
	return "https://github.com" + trend.repo
}
