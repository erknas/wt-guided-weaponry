package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/zeze322/wt-guided-weaponry/internal/api"
	"github.com/zeze322/wt-guided-weaponry/internal/db/mongodb"
	"github.com/zeze322/wt-guided-weaponry/internal/db/postgresdb"
	"github.com/zeze322/wt-guided-weaponry/internal/logger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load env file")
	}

	var (
		port            = os.Getenv("PORT")
		postgresURL     = os.Getenv("POSTGRES_URL")
		mongoURL        = os.Getenv("MONGO_URL")
		mongoDatabase   = os.Getenv("MONGODB_DATABASE")
		mongoCollection = os.Getenv("MONGODB_COLLECTION")
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	postgresConn, err := postgresdb.New(ctx, postgresURL)
	if err != nil {
		log.Fatal(err)
	}

	defer postgresConn.Close(ctx)

	mongoClient, err := mongodb.New(ctx, mongoURL, mongoDatabase, mongoCollection)
	if err != nil {
		log.Fatal(err)
	}

	defer mongoClient.Close(ctx)

	logger := logger.SetupLogger()

	server := api.NewServer(logger, port, postgresConn, mongoClient)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
