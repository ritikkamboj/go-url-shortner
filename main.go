package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json : "id"`
	OriginalUrl  string    `json : "original_url"`
	ShortUrl     string    `json : "shot_url"`
	CreationDate time.Time `json : "creation_time"`
}

var urlDB = make(map[string]URL)

func generateShortURL(OriginalUrl string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalUrl))
	// this above gives us bites slice
	fmt.Println("hasher is :", hasher)
	data := hasher.Sum(nil)
	fmt.Println("data is :", data)

	hash := hex.EncodeToString(data)
	fmt.Println("hash is :", hash)
	return hash[:8]

}

// creating a function string to URL

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}

	return url, nil
}

func createURL(OriginalUrl string) string {
	shortURL := generateShortURL(OriginalUrl)
	id := shortURL
	urlDB[id] = URL{
		ID:           id,
		OriginalUrl:  OriginalUrl,
		ShortUrl:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL

}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("jai baabe ki")
	fmt.Fprintf(w, "Hello World ")
}

func main() {
	original_url := "https://github.com/ritikkamboj"
	data1 := generateShortURL(original_url)
	fmt.Println("data1 is :", data1)

	// Register the handler function to regiter all request to root URL

	http.HandleFunc("/", handler)
	// starting the HTTP server on some port

	fmt.Println("server is going to start on 3000....")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println("There is error on starting server ..", err)
	}

}
