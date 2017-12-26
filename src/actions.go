package main

import (
	"bytes"
	"github.com/bwmarrin/discordgo"
	"github.com/jinzhongmin/gtra"
	"log"
)

func mentionTarget(targetUserID string) *bytes.Buffer {
	var targetBuffer bytes.Buffer
	targetBuffer.WriteString("<@")
	targetBuffer.WriteString(targetUserID)
	targetBuffer.WriteString("> ")
	return &targetBuffer
}

func sendTargetedMessage(s *discordgo.Session, strParams ...string) {
	content := mentionTarget(strParams[1])
	content.WriteString(strParams[2])
	_, err := s.ChannelMessageSend(strParams[0], content.String())
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("User %s requested info: %s\n", strParams[1], strParams[2])
}

func translate(s *discordgo.Session, strParams ...string) {
	t := gtra.New(strParams[3])
	res, err := t.To(strParams[2])
	if err != nil {
		log.Panicln(err)
	}
	sendTargetedMessage(s, strParams[0], strParams[1], res)
}
