package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	discordToken  = os.Getenv("DISCORD_TOKEN")
	witAITokenMap = map[string]string{
		"en": os.Getenv("WIT_AI_TOKEN_EN"),
		"ja": os.Getenv("WIT_AI_TOKEN_JA"),
	}
	witTextURI   = os.Getenv("WIT_AI_TEXT_URI")
	witSpeechURI = os.Getenv("WIT_AI_SPEECH_URI")
	locale       = "en"
	debug, _     = strconv.ParseBool(os.Getenv("DEBUG"))
	master       = os.Getenv("DISCORD_MASTER_ID")
	gitURI       = os.Getenv("GIT_REPO_URI")
	startTime    = time.Now()
)

func main() {

	log.Println("Connecting to Discord with Token: ", discordToken)
	// Connect to Discord

	discord, err := discordgo.New(fmt.Sprintf("Bot %s", discordToken))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Opening Socket...")
	// Open Websocket
	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer discord.Close()

	setSessionHandlers(discord)

	<-make(chan struct{})
}
