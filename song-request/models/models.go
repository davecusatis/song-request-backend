package models

import (
	"github.com/dgrijalva/jwt-go"
)

type PostData struct {
	ContentType string   `json:"content_type"`
	Message     string   `json:"message"`
	Targets     []string `json:"targets"`
}

type Song struct {
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Genre       string `json:"genre"`
	Game        string `json:"game"`
	RequestedBy string `json:"requestedBy,omitempty"`
}

type SongRequestMessage struct {
	Token       *TokenData
	MessageType string      `json:"type"`
	Data        MessageData `json:"data"`
}

type TokenData struct {
	Token     string
	UserID    string
	ChannelID string
	Role      string
}

type PubsubPerms struct {
	Send   []string `json:"send"`
	Listen []string `json:"listen"`
}
type SRClaims struct {
	OpaqueUserID string      `json:"opaque_user_id"`
	UserID       string      `json:"user_id"`
	ChannelID    string      `json:"channel_id"`
	Role         string      `json:"role"`
	PubsubPerms  PubsubPerms `json:"pubsub_perms"`
	jwt.StandardClaims
}

type MessageData struct {
	Songlist []Song `json:"songlist,omitempty"`
	Playlist []Song `json:"playlist,omitempty"`
}
