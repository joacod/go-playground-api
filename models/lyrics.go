package models

// SongLyrics : Song Lyrics model
type SongLyrics struct {
	Artist    string `json:"Artist"`
	SongTitle string `json:"SongTitle"`
	Lyrics    string `json:"Lyrics"`
}
