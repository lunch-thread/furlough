package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/slack-go/slack"
)

func main() {
	token := os.Getenv("TOKEN")

	if token == "" {
		log.Fatal("token must be set")
	}

	rtm := slack.New(token).NewRTM()

	go rtm.ManageConnection()

	e := json.NewEncoder(os.Stdout)

	for msg := range rtm.IncomingEvents {
		if err := e.Encode(map[string]interface{}{
			"@timestamp": time.Now().Format(time.RFC3339Nano),
			"msg":        msg.Data,
		}); err != nil {
			panic(err)
		}

		switch e := msg.Data.(type) {
		case *slack.ConnectingEvent:
		case *slack.InvalidAuthEvent:
			log.Fatalln("invalid auth token")
		case *slack.ConnectionErrorEvent:
			log.Fatalf("connection error: %v\n", e)
		}
	}
}
