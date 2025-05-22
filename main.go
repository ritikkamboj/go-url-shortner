package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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

func main() {
	original_url := "https://github.com/ritikkamboj"
	data1 := generateShortURL(original_url)
	fmt.Println("data1 is :", data1)

}
