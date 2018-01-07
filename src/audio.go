package main

import (
	"bytes"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"layeh.com/gopus"
	"log"
	"os"
)

const (
	in_channels     int = 2
	sampleRate      int = 16000
	sampleSize      int = 320
	minimumSendSize int = 5120
)

var (
	speakers map[uint32]*gopus.Decoder
	dataBuff = &bytes.Buffer{}
)

func printVoiceConnections(session *discordgo.Session) {
	for k, v := range session.VoiceConnections {
		log.Printf("key[%s] value[%s]\n", k, v.UserID)
	}
}

// Records audio to file
func recordRawAudioToFile(v *discordgo.VoiceConnection) {
	log.Println("Start recording...")
	recv := make(chan *discordgo.Packet, 2)
	go decode(v, recv)

	f, err := os.Create(fmt.Sprintf("%s.raw", v.UserID))
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	for {
		buf, _ := getMonoFromDGPacket(recv)
		_, err := f.Write(buf)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Currently there are only plans for supporting Wit.ai
func prepRawAudioForWitAPI(v *discordgo.VoiceConnection) {
	log.Println("Listening to: ", v.UserID)
	recv := make(chan *discordgo.Packet, 2)
	go decode(v, recv)

	for {
		if !v.Ready {
			return
		}
		buf, _ := getMonoFromDGPacket(recv)
		dataBuff.Write(buf)

		if dataBuff.Len() > 0 {
			if dataBuff.Len() >= minimumSendSize {
				go sendAudioToWitAPI(dataBuff.Bytes(), witAITokenMap[locale]) //, respBody)
			}
			dataBuff.Reset()
		}
	}
}

func getMonoFromDGPacket(recv chan *discordgo.Packet) ([]byte, int) {
	p, ok := <-recv
	if !ok {
		return nil, 0
	}

	//log.Println("Buffering...")

	buf := make([]byte, sampleSize*in_channels)

	i := 0
	for ; i < len(p.PCM); i++ {
		buf[i] = uint8(p.PCM[i] >> 8)
	}

	return buf, i
}

/*
Copyright (c) 2015, Bruce
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of dgvoice nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
func decode(v *discordgo.VoiceConnection, c chan *discordgo.Packet) {
	var err error
	if c == nil {
		return
	}

	for {
		if !v.Ready || v.OpusRecv == nil {
			log.Println("Discordgo not to receive opus packets. %+v : %+v", v.Ready, v.OpusSend)
			return
		}

		p, ok := <-v.OpusRecv
		if !ok {
			return
		}

		if speakers == nil {
			speakers = make(map[uint32]*gopus.Decoder)
		}

		_, ok = speakers[p.SSRC]
		if !ok {
			speakers[p.SSRC], err = gopus.NewDecoder(sampleRate, in_channels)
			if err != nil {
				log.Panicln("error creating opus decoder", err)
				continue
			}
		}

		p.PCM, err = speakers[p.SSRC].Decode(p.Opus, sampleSize, false)
		if err != nil {
			log.Panicln("Error decoding opus data", err)
			continue
		}

		c <- p
	}
}
