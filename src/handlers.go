package main

import (
	//"bytes"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
)

func setSessionHandlers(s *discordgo.Session) {
	s.AddHandler(textCommandHandler)
}

func setVoiceHandlers(vc *discordgo.VoiceConnection) {
	vc.AddHandler(speechHandler)
}

func speechHandler(vc *discordgo.VoiceConnection, vs *discordgo.VoiceSpeakingUpdate) {
	if vs.Speaking {
		if vs.UserID == master {
			listening = true
		}
	} else {
		listening = false
	}
}

func textCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Message.Content[0], "$") {
		separatedContent := strings.SplitN(m.Message.Content, " ", 3)

		switch command := separatedContent[0]; command {
		case "$ping":
			go sendTargetedMessage(s, m.Message.ChannelID, m.Message.Author.ID, "pong")
		case "$update":
		case "$uptime":
			go sendTargetedMessage(s, m.Message.ChannelID, m.Message.Author.ID, time.Since(startTime).String())
		case "$locale":
			if len(separatedContent) > 1 {
				locale = separatedContent[1]
			}
		case "$train":
			if separatedContent[1] == "voice" {
				listening = true
			} else if separatedContent[1] == "text" {

			}
		case "$translate":
			if len(separatedContent) == 3 {
				go translate(s, m.Message.ChannelID, m.Message.Author.ID, separatedContent[1], separatedContent[2])
			} else if len(separatedContent) == 2 {
				go translate(s, m.Message.ChannelID, m.Message.Author.ID, locale, separatedContent[1])
			} else {
				msgs, err := s.ChannelMessages(m.Message.ChannelID, 1, m.Message.ID, "", "")
				if err != nil {
					log.Panicln(err)
				}
				go translate(s, m.Message.ChannelID, m.Message.Author.ID, locale, msgs[0].Content)
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
			go sendTargetedMessage(s, m.Message.ChannelID, m.Message.Author.ID, gitUri)
		default:
		}
	}
}
