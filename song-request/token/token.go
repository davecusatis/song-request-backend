package token

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/davecusatis/song-request-backend/song-request/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func ExtractAndValidateTokenFromHeader(header http.Header) (*models.TokenData, error) {
	if authHeaders, ok := header["Authorization"]; ok {
		for _, header := range authHeaders {
			if strings.Contains(header, "Bearer") {
				tokenStr := strings.Split(header, " ")[1]
				secret, _ := base64.StdEncoding.DecodeString("")
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
					}, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("Unable to get token")
}

func ValidateToken(token string) bool {
	return true
}
