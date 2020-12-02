package auth

import (
	"context"
	"errors"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

var app *firebase.App

// Init initializes the Firebase authentication service
func Init() {
	newApp, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalln("Error initializing Firebase app.")
	}
	app = newApp
}

// VerifyToken verifies the string token with the Firebase authenticaion client.
func VerifyToken(token string) (*auth.Token, error) {
	ctx := context.Background()
	client, err := app.Auth(ctx)
	if err != nil {
		return &auth.Token{}, err
	}

	authToken, err := client.VerifyIDToken(ctx, token)
	if err != nil {
		return &auth.Token{}, errors.New("ID token verification failed")
	}
	return authToken, nil
}
