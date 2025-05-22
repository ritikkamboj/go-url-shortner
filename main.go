package main

import (
	"crypto/md5"
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

func generateShortURL(OriginalUrl string) {
	hasher := md5.New()
	hasher.Write([]byte(OriginalUrl))
	fmt.Println("hasher is :", hasher)

}

func main() {
	original_url := "https://github.com/ritikkamboj"
	generateShortURL(original_url)

}
