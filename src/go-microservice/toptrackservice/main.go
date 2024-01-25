package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	host     = "https://ws.audioscrobbler.com/2.0/?method="
	endPoint = "geo.gettoptracks&country="
	format   = "json"
	api_key  = "dbfe4bcd35b4bd186ed92c62a8cb0790"
)

func main() {
	r := chi.NewRouter()

	r.Get("/{country}", getTopTrack)
	http.ListenAndServe(":3001", r)
}

func getTopTrack(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s%s%s&api_key=%s&format=%s", host, endPoint, chi.URLParam(r, "country"), api_key, format)
	resp, err := http.Get(url)
	if err != nil {
		w.Write([]byte("error while getting the top tracks"))
		return
	}
	defer resp.Body.Close()
	tracks := &TopTracks{}
	if err := json.NewDecoder(resp.Body).Decode(tracks); err != nil {
		w.Write([]byte("error while decoding to object"))
		return
	}
	fmt.Println("Top Track:", SerializeToJSONString(tracks))
	w.Write([]byte("got Top Tracker List"))
}

// SerializeToJSONString serializes a structure into a JSON format and applies formatting.
func SerializeToJSONString(v interface{}) string {
	binVal, err := json.MarshalIndent(v, "", " ")

	if err != nil {
		return ""
	}

	return string(binVal)
}

// TopTracks - contains all the top track
type TopTracks struct {
	Tracks Tracks `json:"tracks"`
}

type Tracks struct {
	Track []Track `json:"track"`
	Attr  Attr    `json:"@attr"`
}

type Track struct {
	Name       string `json:"name"`
	Duration   string `json:"duration"`
	Listeners  string `json:"listeners"`
	Mbid       string `json:"mbid"`
	URL        string `json:"url"`
	Streamable struct {
		Text      string `json:"#text"`
		Fulltrack string `json:"fulltrack"`
	} `json:"streamable"`
	Artist struct {
		Name string `json:"name"`
		Mbid string `json:"mbid"`
		URL  string `json:"url"`
	} `json:"artist"`
	Image []struct {
		Text string `json:"#text"`
		Size string `json:"size"`
	} `json:"image"`
	Attr struct {
		Rank string `json:"rank"`
	} `json:"@attr"`
}

type Attr struct {
	Country    string `json:"country"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total      string `json:"total"`
}
