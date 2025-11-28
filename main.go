package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"go-cli-player/internal/auth"
	"go-cli-player/internal/opener"
	"go-cli-player/internal/search"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	if len(os.Args) < 3 {
		fmt.Println("Usage: go-cli play \"song name\"")
		return
	}

	command := os.Args[1]
	query := strings.Join(os.Args[2:], " ")

	if command != "play" {
		fmt.Println("Unknown command:", command)
		return
	}

	token, err := auth.GetAccessToken()
	if err != nil {
		fmt.Println("Error getting token:", err)
		return
	}

	trackName, trackURL, err := search.SearchTrack(token, query)
	if err != nil {
		fmt.Println("Error searching track:", err)
		return
	}

	fmt.Println("Playing:", trackName)

	err = opener.OpenURL(trackURL)
	if err != nil {
		fmt.Println("Error opening Spotify:", err)
		return
	}
}
