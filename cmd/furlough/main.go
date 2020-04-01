package main

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	token := os.Args[len(os.Args)-1]

	if token == "" {
		log.Fatal("token must be set")
	}

	rtm := slack.New(token).NewRTM()

	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		e, ok := msg.Data.(*slack.UserChangeEvent)
		if !ok {
			continue
		}

		fmt.Printf("user %s changed\n", e.User.Name)
	}
}
