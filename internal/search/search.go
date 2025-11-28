package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type spotifySearchResponse struct {
	Tracks struct {
		Items []struct {
			Name         string `json:"name"`
			URI          string `json:"uri"`
			ExternalURLs struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Artists []struct {
				Name string `json:"name"`
			} `json:"artists"`
		} `json:"items"`
	} `json:"tracks"`
}

func SearchTrack(token string, query string) (string, string, error) {
	if strings.TrimSpace(query) == "" {
		return "", "", errors.New("empty song query")
	}

	encoded := url.QueryEscape(query) // does this lose control-> lose+control
	apiURL := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=track&limit=1", encoded)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	var result spotifySearchResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", "", err
	}

	if len(result.Tracks.Items) == 0 {
		return "", "", errors.New("no tracks found")
	}

	track := result.Tracks.Items[0]

	trackName := track.Name + " - " + track.Artists[0].Name
	trackURL := track.URI

	return trackName, trackURL, nil
}
