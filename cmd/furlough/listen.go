package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/slack-go/slack"
)

func listen(ctx context.Context, token string) error {
	rtm := slack.New(token).NewRTM()

	go rtm.ManageConnection()

	e := json.NewEncoder(os.Stdout)

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			if err := e.Encode(map[string]interface{}{
				"@timestamp": time.Now().Format(time.RFC3339Nano),
				"type":       msg.Type,
				"data":       msg.Data,
			}); err != nil {
				return fmt.Errorf("json encode: %w", err)
			}

			switch e := msg.Data.(type) {
			case *slack.InvalidAuthEvent:
				return errors.New("invalid auth token")
			case *slack.ConnectionErrorEvent:
				return fmt.Errorf("connection error: %v", e)
			}
		case <-ctx.Done():
			return fmt.Errorf("context done: %w", ctx.Err())
		}
	}
}
