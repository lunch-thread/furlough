package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/slack-go/slack"
)

func main() {
	token := flag.String("token", "", "slack token")
	flag.Parse()
	if *token == "" {
		log.Fatal("token must be set")
	}

	rtm := slack.New(*token).NewRTM()

	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		e, ok := msg.Data.(*slack.UserChangeEvent)
		if !ok {
			continue
		}

		fmt.Printf("user %s changed\n", e.User.Name)
	}
}
