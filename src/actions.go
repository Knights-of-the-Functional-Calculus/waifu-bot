package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jinzhongmin/gtra"
	"log"
)

func sendTargetedMessage(s *discordgo.Session, user *discordgo.User, strParams ...string) {
	content := fmt.Sprintf("%s %s", user.Mention(), strParams[1])
	_, err := s.ChannelMessageSend(strParams[0], content)
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("User %s requested info: %s\n", user.ID, strParams[1])
}

func translate(s *discordgo.Session, user *discordgo.User, strParams ...string) {
	t := gtra.New(strParams[2])
	res, err := t.To(strParams[1])
	if err != nil {
		log.Panicln(err)
	}
	sendTargetedMessage(s, user, strParams[0], res)
}

func grabVoiceState(s *discordgo.Session, user *discordgo.User, channelID string) *discordgo.VoiceState {
	textChannel, err := s.State.Channel(channelID)
	if err != nil {
		log.Panicln(err)
	}
	guild, err := s.State.Guild(textChannel.GuildID)
	if err != nil {
		log.Panicln(err)
	}
	userID := user.ID
	for _, vs := range guild.VoiceStates {
		if vs.UserID != userID {
			return vs
		}
	}
	return nil
}

func readAudio(s *discordgo.Session, vs *discordgo.VoiceState) {
	dgv, err := s.ChannelVoiceJoin(vs.GuildID, vs.ChannelID, false, false)
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		recordRawAudioToFile(dgv)
	} else {
		log.Println("Using Speech API Token: ", witAITokenMap[locale])
		prepRawAudioForWitAPI(dgv)
	}
}
