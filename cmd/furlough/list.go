package main

import (
	"context"
	"fmt"
	"sort"

	"github.com/slack-go/slack"
)

type userSlice []slack.User

func (p userSlice) Len() int { return len(p) }

func (p userSlice) Less(i, j int) bool { return p[i].Updated.Time().Before(p[j].Updated.Time()) }

func (p userSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func list(ctx context.Context, token string) error {
	us, err := slack.New(token).GetUsersContext(ctx)
	if err != nil {
		return fmt.Errorf("get users: %w", err)
	}

	// sort so they are in last-updated order
	sort.Sort(userSlice(us))

	// filter for only deactivated accounts and not bots
	for _, u := range us {
		if !u.IsBot && u.Deleted {
			const layout = "Jan _2 06 15:04:05"
			fmt.Printf("%s %s\n", u.Updated.Time().Format(layout), u.Name)
		}
	}

	return nil
}
