package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/davecusatis/song-request-backend/song-request/models"
)

// SendPubsubBroadcastMessage sends a pubsub message to twitch broadcast topic
func SendPubsubBroadcastMessage(message *models.SongRequestMessage, token *models.TokenData) {
	srMessage, _ := json.Marshal(message)

	postData, _ := json.Marshal(&models.PostData{
		ContentType: "application/json",
		Message:     string(srMessage),
		Targets:     []string{"broadcast"},
	})

	r, _ := http.NewRequest("POST",
		fmt.Sprintf("https://api.twitch.tv/extensions/message/%s", token.ChannelID),
		bytes.NewReader(postData))

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.Token))
	r.Header.Add("Client-Id", "cm5nkhrq18yy02yy9tp108lx745vcx")
	r.Header.Add("Content-Type", "application/json")

	c := &http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		log.Printf("error %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		log.Printf("Error from twitch API: expected 204 got %d", resp.StatusCode)
	}
}
