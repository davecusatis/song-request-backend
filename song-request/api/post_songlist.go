package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/davecusatis/song-request-backend/song-request/models"
	"github.com/davecusatis/song-request-backend/song-request/token"
)

// PostSonglist is how broadcasters send us songlists
func (a *API) PostSonglist(w http.ResponseWriter, req *http.Request) {
	// validate token
	token, err := token.ExtractAndValidateTokenFromHeader(req.Header)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprintf("error %s", err)))
		return
	}

	if token.Role != "broadcaster" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized: not broadcaster"))
		return
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error reading data"))
		return
	}

	var newSonglist []models.Song
	err = json.Unmarshal(body, &newSonglist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error reading data"))
		return
	}
	// verify songlist
	// save songlist
	err = saveSonglist(a.S3Uploader, token.UserID, newSonglist)
	if err != nil {
		log.Printf("error uploading file %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error sending data to s3"))
		return
	}

	// blast message to clients
	// a.Aggregator.MessageChan <- &models.SongRequestMessage{
	// 	MessageType: "songlistUpdated",
	// 	Data:        models.MessageData{
	// 	// Songlist: newSonglist,
	// 	},
	// 	Token: token,
	// }

	w.Write([]byte("OK"))
}

func saveSonglist(s *s3manager.Uploader, userID string, songlist []models.Song) error {
	userFilename := fmt.Sprintf("%s.txt", userID)
	f, err := os.Create(userFilename)
	defer os.Remove(userFilename)

	if err != nil {
		return fmt.Errorf("Error creating file %s", err)
	}
	for _, song := range songlist {
		f.WriteString(fmt.Sprintf("%s , %s\n", song.Title, song.Artist))
	}
	f.Close()

	f, err = os.Open(userFilename)
	if err != nil {
		return fmt.Errorf("Error opening file %s", err)
	}
	// Upload the file to s3.
	_, err = s.Upload(&s3manager.UploadInput{
		Bucket:      aws.String("song-request-distro"),
		Key:         aws.String(userFilename),
		ContentType: aws.String("text/plain"),
		ACL:         aws.String("public-read"),
		Body:        f,
	})
	if err != nil {
		log.Printf("error: %s", err)
		return fmt.Errorf("failed to upload file, %v", err)
	}
	return nil
}
