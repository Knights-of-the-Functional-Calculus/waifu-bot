package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
)

var (
	listening             = false
	reading               = false
	currentReadingChannel = ""
)

func setSessionHandlers(s *discordgo.Session) {
	s.AddHandler(textCommandHandler)
}

func textCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Message
	if strings.HasPrefix(msg.Content, "$") {
		separatedContent := strings.SplitN(msg.Content, " ", 3)

		switch separatedContent[0] {
		case "$ping":
			go sendTargetedMessage(s, msg.Author, msg.ChannelID, "pong")
		case "$update":
		case "$uptime":
			go sendTargetedMessage(s, msg.Author, msg.ChannelID, time.Since(startTime).String())
		case "$locale":
			if len(separatedContent) > 1 {
				locale = separatedContent[1]
			}
		case "$train":
			if len(separatedContent) > 1 {
				if separatedContent[1] == "voice" {
					vs := grabVoiceState(s, msg.Author, msg.ChannelID)
					if vs == nil {
						go sendTargetedMessage(s, msg.Author, msg.ChannelID, "You must be in a voice channel to issue this command.")
						return
					}
					if listening {
						err := s.VoiceConnections[vs.GuildID].Disconnect()
						if err != nil {
							log.Panicln(err)
						}
						go sendTargetedMessage(s, msg.Author, msg.ChannelID, "Voice training turned off.")
					} else {
						go readAudio(s, vs)
						go sendTargetedMessage(s, msg.Author, msg.ChannelID, "Voice training turned on.")
					}
					listening = !listening
				} else if separatedContent[1] == "text" {
					var responseMessage string
					if reading {
						currentReadingChannel = ""
						responseMessage = fmt.Sprintf("Text training turned off in #%s", currentReadingChannel)
					} else {
						currentReadingChannel = msg.ChannelID
						responseMessage = fmt.Sprintf("Text training turned on in #%s", currentReadingChannel)
					}
					go sendTargetedMessage(s, msg.Author, msg.ChannelID, responseMessage)
					reading = !reading
				}
			}
		case "$translate":
			if len(separatedContent) == 3 {
				go translate(s, msg.Author, msg.ChannelID, separatedContent[1], separatedContent[2])
			} else if len(separatedContent) == 2 {
				go translate(s, msg.Author, msg.ChannelID, locale, separatedContent[1])
			} else {
				msgs, err := s.ChannelMessages(msg.ChannelID, 1, msg.ID, "", "")
				if err != nil {
					log.Panicln(err)
				}
				go translate(s, msg.Author, msg.ChannelID, locale, msgs[0].Content)
			}
		case "$anime":
			if len(separatedContent) > 2 {
				switch parameter := separatedContent[1]; parameter {
				case "find":
				case "rate":
				case "recommend":
				}
			}
		case "$nsfw":
		case "$git":
			go sendTargetedMessage(s, msg.Author, msg.ChannelID, gitUri)
		default:
		}
	} else if reading && msg.ChannelID == currentReadingChannel {
		go sendTextToWitAPI(msg.Content, witAITokenMap[locale])
	}
}
