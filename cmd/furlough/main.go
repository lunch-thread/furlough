package main

import (
	"log"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	token := os.Getenv("TOKEN")

	if token == "" {
		log.Fatal("token must be set")
	}

	rtm := slack.New(token).NewRTM()

	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		log.Printf("got message %s\n", msg.Type)

		switch e := msg.Data.(type) {
		case *slack.ConnectingEvent:
			log.Println("connecting")
		case *slack.InvalidAuthEvent:
			log.Fatalln("invalid auth token")
		case *slack.ConnectionErrorEvent:
			log.Println("connection error:", e.Error())
		case *slack.UserChangeEvent:
			log.Printf("user %s changed, deleted: %t\n", e.User.Name, e.User.Deleted)
		}
	}
}
