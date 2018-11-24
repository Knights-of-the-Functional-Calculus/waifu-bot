package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/jinzhongmin/gtra"
)

func errorLog(err error) string {
	log.Panicln(err)
	return "Failed"
}

func sendTargetedMessage(s *discordgo.Session, user *discordgo.User, strParams ...string) {
	content := fmt.Sprintf("%s %s", user.Mention(), strParams[1])
	_, err := s.ChannelMessageSend(strParams[0], content)
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("User %s requested info: %s\n", user.ID, strParams[1])
}

func translate(s *discordgo.Session, user *discordgo.User, strParams ...string) {
	t := gtra.New()
	res := t.To(strParams[1], errorLog)
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
		log.Printf("%+v\n", vs)
		if vs.UserID == userID {
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
		log.Println("Speech API Token: ", witAITokenMap[locale])
		prepRawAudioForWitAPI(dgv)
	}
}
