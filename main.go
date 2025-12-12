package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN not set in environment")
	}

	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	log.Println("Openning session...")
	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer sess.Close()

	log.Println("Bot is now running. Registering commands...")
	RegisterCommands(sess)
	log.Println("Registering handlers...")
	sess.AddHandler(RegisterHandlers)
	log.Println("BOT IS ONLINE!!!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
