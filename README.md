# Go Spotify CLI

A simple cross-platform CLI tool written in Go that allows you to search a song and open it directly in the Spotify Desktop App.


The tool uses Spotify’s **Client Credentials Flow** (no OAuth login required) and works with **free** Spotify accounts.

---

## Features

- Search any song using Spotify Web API  
- Opens the track directly in Spotify Desktop App  
- Cross-platform added:
  - Linux → `xdg-open`
  - macOS → `open`
  - Windows → `rundll32`
- No OAuth 

---

## Setup

### 1. Create a `.env` file in the project root
SPOTIFY_CLIENT_ID=your_client_id_here
SPOTIFY_CLIENT_SECRET=your_client_secret_here

### 2. Install dotenv loader

go get github.com/joho/godotenv

### 3. Run the CLI

go run main.go play "song name"
