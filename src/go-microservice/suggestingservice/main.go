package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	host    = "https://ws.audioscrobbler.com/2.0/?method="
	path1   = "track.getsimilar&artist="
	path2   = "&track="
	format  = "json"
	api_key = "dbfe4bcd35b4bd186ed92c62a8cb0790"
)

func main() {
	r := chi.NewRouter()

	r.Get("/{artist}/{track}", getSuggestion)
	http.ListenAndServe(":3003", r)
}

func getSuggestion(w http.ResponseWriter, r *http.Request) {
	artist := chi.URLParam(r, "artist")
	track := chi.URLParam(r, "track")
	url := fmt.Sprintf("%strack.getsimilar&artist=%s&track=%s&api_key=%s&format=%s", host, artist, track, api_key, format)
	resp, err := http.Get(url)
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	defer resp.Body.Close()
	suggestionTrack := &SuggestionTrack{}
	if err := json.NewDecoder(resp.Body).Decode(suggestionTrack); err != nil {
		w.Write([]byte("error"))
		return
	}
	fmt.Println("Final suggestion Track:", SerializeToJSONString(suggestionTrack))
	w.Write([]byte("got suggestion Track"))
}

// SerializeToJSONString serializes a structure into a JSON format and applies formatting.
func SerializeToJSONString(v interface{}) string {
	binVal, err := json.MarshalIndent(v, "", " ")

	if err != nil {
		return ""
	}

	return string(binVal)
}

type SuggestionTrack struct {
	Similartracks Similartracks `json:"similartracks"`
}

type Similartracks struct {
	Track []Track `json:"track"`
	Attr  Attr    `json:"@attr"`
}

type Track struct {
	Name       string  `json:"name"`
	Playcount  int     `json:"playcount"`
	Mbid       string  `json:"mbid,omitempty"`
	Match      float64 `json:"match"`
	URL        string  `json:"url"`
	Streamable struct {
		Text      string `json:"#text"`
		Fulltrack string `json:"fulltrack"`
	} `json:"streamable"`
	Duration int `json:"duration,omitempty"`
	Artist   struct {
		Name string `json:"name"`
		Mbid string `json:"mbid"`
		URL  string `json:"url"`
	} `json:"artist"`
	Image []struct {
		Text string `json:"#text"`
		Size string `json:"size"`
	} `json:"image"`
}

type Attr struct {
	Artist string `json:"artist"`
}
