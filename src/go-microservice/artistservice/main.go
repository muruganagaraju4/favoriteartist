package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	host     = "https://ws.audioscrobbler.com/2.0/?method="
	endPoint = "artist.getinfo&artist="
	format   = "json"
	api_key  = "dbfe4bcd35b4bd186ed92c62a8cb0790"
)

func main() {
	r := chi.NewRouter()

	r.Get("/artist/{artist}", getArtist)
	http.ListenAndServe(":3002", r)
}

func getArtist(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s%s%s&api_key=%s&format=%s", host, endPoint, chi.URLParam(r, "artist"), api_key, format)
	resp, err := http.Get(url)
	if err != nil {
		w.Write([]byte("error while getting the artist info"))
		return
	}
	defer resp.Body.Close()
	artistInfo := &ArtistInfo{}
	if err := json.NewDecoder(resp.Body).Decode(artistInfo); err != nil {
		w.Write([]byte("error while decoding to object"))
		return
	}
	fmt.Println("Final artist list:", SerializeToJSONString(artistInfo))
	w.Write([]byte("got artist info"))
}

// SerializeToJSONString serializes a structure into a JSON format and applies formatting.
func SerializeToJSONString(v interface{}) string {
	binVal, err := json.MarshalIndent(v, "", " ")

	if err != nil {
		return ""
	}

	return string(binVal)
}

type ArtistInfo struct {
	Artist Artist `json:"artist"`
}

type Artist struct {
	Name       string  `json:"name"`
	Mbid       string  `json:"mbid"`
	URL        string  `json:"url"`
	Image      []Image `json:"image"`
	Streamable string  `json:"streamable"`
	Ontour     string  `json:"ontour"`
	Stats      Stats   `json:"stats"`
	Similar    struct {
		Artist []struct {
			Name  string `json:"name"`
			URL   string `json:"url"`
			Image []struct {
				Text string `json:"#text"`
				Size string `json:"size"`
			} `json:"image"`
		} `json:"artist"`
	} `json:"similar"`
	Tags struct {
		Tag []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"tag"`
	} `json:"tags"`
	Bio struct {
		Links struct {
			Link struct {
				Text string `json:"#text"`
				Rel  string `json:"rel"`
				Href string `json:"href"`
			} `json:"link"`
		} `json:"links"`
		Published string `json:"published"`
		Summary   string `json:"summary"`
		Content   string `json:"content"`
	} `json:"bio"`
}

type Image struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}

type Stats struct {
	Listeners string `json:"listeners"`
	Playcount string `json:"playcount"`
}
