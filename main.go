package main

import (
	"os"

	"github.com/go-chat-bot/bot/slack"
)

func main() {
	slack.Run(os.Getenv("SLACK_TOKEN"))
}
