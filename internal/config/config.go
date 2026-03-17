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
		DatabaseURL: "https://portfolio-b755c-default-rtdb.asia-southeast1.firebasedatabase.app/",
	}

	jsonData := os.Getenv("FIREBASE_SERVICE_ACCOUNT")

	
	opt := option.WithCredentialsJSON([]byte(jsonData))

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("Error initializing firebase app: %v", err)
	}

	return app
}