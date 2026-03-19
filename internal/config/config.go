package config

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func Init_firebase() *firebase.App {
	ctx := context.Background()

	config := &firebase.Config{
		DatabaseURL: os.Getenv("FIREBASE_DB_URL"),
	}

	jsonData := os.Getenv("FIREBASE_SERVICE_ACCOUNT")

	
	opt := option.WithCredentialsJSON([]byte(jsonData))

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("Error initializing firebase app: %v", err)
	}

	return app
}