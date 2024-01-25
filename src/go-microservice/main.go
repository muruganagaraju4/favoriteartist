package main

import (
   "io"
   "net/http"

   "github.com/go-chi/chi"
)

func main() {
   r := chi.NewRouter()

   r.Get("/", homePage)
   r.Get("/{country}", getTopTrack)
   r.Get("/artist/{artist}", getArtist)
   r.Get("/{artist}/{track}", getSuggestion)
   http.ListenAndServe(":3000", r)
}

func getTopTrack(w http.ResponseWriter, r *http.Request) {
   resp, err := http.Get("http://localhost:3001/" + chi.URLParam(r, "country"))
   if err != nil {
      w.Write([]byte("error while making a call to top track service"))
      return
   }
   defer resp.Body.Close()
   bytes, err := io.ReadAll(resp.Body)
   if err != nil {
      w.Write([]byte("error while reading the bytes"))
      return
   }
   a := string(bytes)
   w.Write([]byte(a))
}

func getArtist(w http.ResponseWriter, r *http.Request) {
   resp, err := http.Get("http://localhost:3002/artist/" + chi.URLParam(r, "artist"))
   if err != nil {
      w.Write([]byte("error while making a call to get artist"))
      return
   }
   defer resp.Body.Close()
   bytes, err := io.ReadAll(resp.Body)
   if err != nil {
      w.Write([]byte("error while reading the bytes"))
      return
   }
   b := string(bytes)

   w.Write([]byte(b))
}

func getSuggestion(w http.ResponseWriter, r *http.Request) {
   resp, err := http.Get("http://localhost:3003/" + chi.URLParam(r, "artist") + "/" + chi.URLParam(r, "track"))
   if err != nil {
      w.Write([]byte("error while making a call to get suggestion"))
      return
   }
   defer resp.Body.Close()
   bytes, err := io.ReadAll(resp.Body)
   if err != nil {
      w.Write([]byte("error while reading the bytes"))
      return
   }
   c := string(bytes)
   w.Write([]byte(c))
}

func homePage(w http.ResponseWriter, r *http.Request) {
   str := `
   ex:
   use localhost:3000/ Home page
   use localhost:3001/india to get top track
   use localhost:3002/artist/Cher to get artist info
   use localhost:3003/Cher/believe to get suggesting track
   `
   w.Write([]byte(str))
}