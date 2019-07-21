# Send Github Trending to Slack

# How to use?
```
$ docker run --rm --env SLACK_API_TOKEN=<token> --env NOTIFY_CHANNEL=#github crossts/github-trending-to-slack:latest
```

or

```
$ docker run --rm --env-file .env crossts/github-trending-to-slack:latest
```
