package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"playground/models"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetLyrics godoc
// @Summary Get lyrics from an Artist and a Song
// @Description Get lyrics from the public api https://api.lyrics.ovh/v1/
// @Tags lyrics
// @Accept json
// @Produce json
// @Param  artist path string true "Artist"
// @Param  song path string true "Song"
// @Success 200 {object} models.SongLyrics
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /lyrics/{artist}/{song} [get]
func (c *Controller) GetLyrics(ctx *gin.Context) {
	artist := ctx.Params.ByName("artist")
	title := ctx.Params.ByName("song")

	apiURL := "https://api.lyrics.ovh/v1/%s/%s"
	url := fmt.Sprintf(apiURL, artist, title)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		ctx.String(resp.StatusCode, "Lyrics not found.")
	} else {
		var song models.SongLyrics
		json.NewDecoder(resp.Body).Decode(&song)

		if song.Lyrics == "" {
			ctx.String(http.StatusNotFound, "Lyrics not found.")
		} else {
			song.Artist = strings.Title(artist)
			song.SongTitle = strings.Title(title)
			ctx.JSON(http.StatusOK, song)
		}
	}

	return
}
