package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/zeze322/wt-guided-weaponry/internal/api"
	"github.com/zeze322/wt-guided-weaponry/internal/db/postgresdb"
	"github.com/zeze322/wt-guided-weaponry/internal/logger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load env file")
	}

	var (
		port        = os.Getenv("PORT")
		postgresURL = os.Getenv("POSTGRES_URL")
	)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	postgres, err := postgresdb.New(ctx, postgresURL)
	if err != nil {
		log.Fatal(err)
	}

	defer postgres.Close(ctx)

	logger := logger.SetupLogger()

	server := api.NewServer(logger, port, postgres)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
