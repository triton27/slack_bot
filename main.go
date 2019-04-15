package main

import (
	"os"
	"strings"
	"github.com/nlopes/slack"
)

var (
		botId string
		botName string
)

func main() {
	token := os.Getenv("SLACK_API_TOKEN")
	bot := NetBot(token)
	go bot.rtm.ManageConnection()

	for {
		select {
			case msg := <-bot.rtm.IncomingEvents:
				switch ev := msg.Data.(type) {
				case *slack.ConnectedEvent:
					botId = ev.Info.User.ID
					botName = ev.Info.User.Name
				case *slack.MessageEvent:
					user := ev.User
					text := ev.Text
					channel := ev.Channel

					if ev.Type == "message" && strings.HasPrefix(text, "<@"+botId+">") {
						bot.handleResponse(user, text, channel)
				}
			}
		}
	}
}