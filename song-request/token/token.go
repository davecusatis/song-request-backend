package token

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/davecusatis/song-request-backend/song-request/models"
	jwt "github.com/dgrijalva/jwt-go"
)

const secret = ""

// ExtractAndValidateTokenFromHeader extracts data and validates it against the secret
func ExtractAndValidateTokenFromHeader(header http.Header) (*models.TokenData, error) {
	if authHeaders, ok := header["Authorization"]; ok {
		for _, header := range authHeaders {
			if strings.Contains(header, "Bearer") {
				tokenStr := strings.Split(header, " ")[1]
				secret, _ := base64.StdEncoding.DecodeString(secret)
				token, err := jwt.ParseWithClaims(tokenStr, &models.SRClaims{}, func(token *jwt.Token) (interface{}, error) {
					return []byte(secret), nil
				})
				if err != nil {
					return nil, fmt.Errorf("Invalid secret")
				}

				if claims, ok := token.Claims.(*models.SRClaims); ok && token.Valid {
					return &models.TokenData{
						Token:     tokenStr,
						UserID:    claims.UserID,
						ChannelID: claims.ChannelID,
						Role:      claims.Role,
					}, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("Unable to get token")
}

func CreateServerToken(data *models.TokenData) *models.TokenData {
	exp := time.Now().UTC().Add(time.Minute * time.Duration(5)).Second()
	claims := models.SRClaims{
		OpaqueUserID: data.OpaqueUserID,
		UserID:       data.UserID,
		ChannelID:    data.ChannelID,
		Role:         "external",
		PubsubPerms: models.PubsubPerms{
			Send:   []string{"*"},
			Listen: []string{"*"},
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(exp),
		},
	}
	secret, _ := base64.StdEncoding.DecodeString(secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secret)
	if err != nil {
		log.Printf("Error signing token: %s", err)
		return nil
	}
	return &models.TokenData{
		UserID:    data.UserID,
		ChannelID: data.ChannelID,
		Role:      "external",
		Token:     ss,
	}
}
