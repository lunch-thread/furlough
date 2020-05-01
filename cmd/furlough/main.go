package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
)

func Main(ctx context.Context) error {
	token := os.Getenv("TOKEN")

	if token == "" {
		log.Fatal("token must be set")
	}

	flag.Parse()

	switch cmd := flag.Arg(0); cmd {
	case "", "listen":
		return listen(ctx, token)
	case "list":
		return list(ctx, token)
	default:
		return fmt.Errorf("unknown command: %s", cmd)
	}
}

func main() {
	if err := Main(context.Background()); err != nil {
		log.Fatal(err)
	}
}
