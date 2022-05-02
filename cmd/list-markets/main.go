package main

import (
	"context"
	"log"

	"github.com/fsctl/go-kalshi/pkg/kalshi"
	"github.com/fsctl/go-kalshi/pkg/kalshi/helpers"
)

func main() {
	ctx := context.Background()

	username, password := helpers.ReadEnvFile()

	kc, err := kalshi.NewKalshiClient(ctx, username, password)
	if err != nil {
		log.Fatalf("Error:  NewKalshiClient failed: '%v'\n", err)
		return
	}

	kc.PrintMarketsList(ctx)
}
