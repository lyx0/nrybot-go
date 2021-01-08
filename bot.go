package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gempir/go-twitch-irc"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	botUser := os.Getenv("TWITCH_USER")
	botPass := os.Getenv("TWITCH_PASSWORD")
	// or client := twitch.NewAnonymousClient() for an anonymous user (no write capabilities)
	client := twitch.NewClient(botUser, botPass)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
		if message.Message == "FeelsDankMan TeaTime" {
			client.Say("nouryxd", "pajaDink")
		}
	})

	client.Say("nouryxd", "pajaDink")

	client.Join("nouryxd")

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
