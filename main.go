package main

import (
	"os"

	_ "github.com/fernand-o/go-bot-custom-responses-plugin"
	"github.com/go-chat-bot/bot/slack"
)

func main() {
	slack.Run(os.Getenv("SLACK_TOKEN"))
}
