package token

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func ExtractTokenFromHeader(header http.Header) (string, error) {
	if authHeaders, ok := header["Authorization"]; ok {
		for _, header := range authHeaders {
			if strings.Contains(header, "Bearer") {
				log.Printf("%v", strings.Split(header, " ")[1])
				return "", nil
			}
		}
	}
	return "", fmt.Errorf("Header not present")
}

func ValidateToken(token string) bool {
	return true
}
