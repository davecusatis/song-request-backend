package models

// TestSonglist returns a bogus songlist for test purposes
func TestSonglist() []Song {
	return []Song{{
		Title:  "ttfaf",
		Artist: "dragonforce",
		Genre:  "bad",
		Game:   "gh3",
	}, {
		Title:  "rip the rev",
		Artist: "avenged sevenfold",
		Genre:  "bad",
		Game:   "rb2",
	}, {
		Title:  "ogap",
		Artist: "dragonforce",
		Genre:  "bad",
		Game:   "gh3 dlc",
	}}
}

// TestPlaylist returns a bogus playlist for test purposes
func TestPlaylist() []Song {
	return []Song{{
		Title:  "ttfaf",
		Artist: "dragonforce",
		Genre:  "bad",
		Game:   "gh3",
	}}
}