package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4" 
	"google.golang.org/api/option"
)

func Init_firebase() *firebase.App {
	ctx := context.Background()

	config := &firebase.Config{
		DatabaseURL: "https://portfolio-b755c-default-rtdb.asia-southeast1.firebasedatabase.app/",
	}

	// This is the secure 2026 method we discussed!
	opt := option.WithAuthCredentialsFile(option.ServiceAccount, "firebase-key.json")

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("Error initializing firebase app: %v", err)
	}

	return app
}