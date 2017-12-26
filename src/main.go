package main

import (
	"bytes"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	discordToken     = os.Getenv("DISCORD_TOKEN")
	discordGuildID   = os.Getenv("DISCORD_GUILDID")
	discordChannelID = os.Getenv("DISCORD_CHANNELID")
	witAITokenMap    = map[string]string{
		"en": os.Getenv("WIT_AI_TOKEN_EN"),
		"ja": os.Getenv("WIT_AI_TOKEN_JA"),
	}
	witSpeechUri = os.Getenv("WIT_AI_SPEECH_URI")
	locale       = "en"
	debug, _     = strconv.ParseBool(os.Getenv("DEBUG"))
	master       = os.Getenv("DISCORD_MASTER_ID")
	gitUri       = os.Getenv("GIT_REPO_URI")
	startTime    = time.Now()
)

func main() {

	log.Println("Connecting to Discord with Token: ", discordToken)
	// Connect to Discord

	var preppedToken bytes.Buffer
	preppedToken.WriteString("Bot ")
	preppedToken.WriteString(discordToken)
	discord, err := discordgo.New(preppedToken.String())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Opening Socket...")
	// Open Websocket
	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Joining Voice Channel at Guild: ", discordGuildID, " Channel: ", discordChannelID)
	// Connect to voice channel.
	// NOTE: Setting mute to false, deaf to true.
	dgv, err := discord.ChannelVoiceJoin(discordGuildID, discordChannelID, false, false)
	if err != nil {
		log.Fatal(err)
	}

	setSessionHandlers(discord)
	setVoiceHandlers(dgv)
	if debug {
		recordRawAudioToFile(dgv)
	} else {
		log.Println("Using Speech API Token: ", witAITokenMap[locale])
		prepRawAudioForWitAPI(dgv)
	}

	// Close connections
	dgv.Close()
	discord.Close()

	return
}
