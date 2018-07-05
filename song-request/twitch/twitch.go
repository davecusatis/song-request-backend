package twitch

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/davecusatis/song-request-backend/song-request/models"
)

type TwitchClient struct {
	ClientID string
	Client   *http.Client
}

// NewTwitchClient returns an instance of our Twitch client
func NewTwitchClient(client *http.Client) *TwitchClient {
	return &TwitchClient{
		ClientID: "cm5nkhrq18yy02yy9tp108lx745vcx",
		Client:   client,
	}
}

func (c *TwitchClient) GetLogin(userID string) string {
	r, _ := http.NewRequest("GET", fmt.Sprintf("https://api.twitch.tv/helix/users?id=%s", userID), nil)
	r.Header.Add("Client-Id", c.ClientID)
	r.Header.Add("Content-Type", "application/json")
	resp, err := c.Client.Do(r)
	if err != nil {
		log.Printf("Error getting user data for %s", userID)
	}
	defer resp.Body.Close()

	var respBody models.TwitchUserResponse
	json.NewDecoder(r.Body).Decode(&respBody)

	return respBody.Data.DisplayName
}
