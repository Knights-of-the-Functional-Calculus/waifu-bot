package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func sendAudioToWitAPI(buffer []byte, speechAPIToken string) {
	reader := bytes.NewReader(buffer)
	req, err := http.NewRequest("POST", witSpeechUri, reader)
	if err != nil {
		log.Panicln("Request could not be built", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", speechAPIToken))
	req.Header.Set("Transfer-encoding", "chunked")
	req.Header.Set("Content-Type", "audio/raw;encoding=signed-integer;bits=16;rate=16000;endian=little")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panicln("Request could not be sent", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("Response body is not readable", err)
	}
	log.Printf("%s: %v\n", resp.Status, string(body))
}

func sendTextToWitAPI(message, speechAPIToken string) {
	uri := fmt.Sprintf("%s%s", witTextUri, url.QueryEscape(message))
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", speechAPIToken))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panicln("Request could not be sent", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("Response body is not readable", err)
	}
	log.Printf("%s: %v\n", resp.Status, string(body))
}
