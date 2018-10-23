package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", createUser).Methods("POST")
	router.HandleFunc("/session", login).Methods("POST")
	router.HandleFunc("/session", logout).Methods("DELETE")
	router.HandleFunc("/songs", createSong).Methods("POST")
	router.HandleFunc("/songs", getTrendingSongs).Methods("GET")
	router.HandleFunc("/songs/{id}/comments", postComment).Methods("POST")
	router.HandleFunc("/songs/{id}/comments/{id}", deleteComment).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func createUser(w http.ResponseWriter, r *http.Request)       {}
func login(w http.ResponseWriter, r *http.Request)            {}
func logout(w http.ResponseWriter, r *http.Request)           {}
func createSong(w http.ResponseWriter, r *http.Request)       {}
func getTrendingSongs(w http.ResponseWriter, r *http.Request) {}
func postComment(w http.ResponseWriter, r *http.Request)      {}
func deleteComment(w http.ResponseWriter, r *http.Request)    {}

type User struct {
	ID             string `json:"id,omitempty"`
	Username       string `json:"string,omitempty"`
	Email          string `json:"email,omitempty"`
	PasswordDigest string `json:"password_digest,omitempty"`
	SessionToken   string `json:"session_token,omitempty"`
}

type Song struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	ArtistID   string `json:"artist_id,omitempty"`
	ArtistName string `json:"artist_name,omitempty"`
	URL        string `json:"url,omitempty"`
}

type Artist struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"id,omitempty"`
}

type Comment struct {
	ID     string `json:"id,omitempty"`
	UserID string `json:"user_id,omitempty"`
	SongID string `json:"song_id,omitempty"`
	Body   string `json:"body,omitempty"`
}
