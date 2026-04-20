package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func download_page(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(io.LimitReader(r.Body, 2048)) // cap at 2 KB
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}
	id := strings.TrimSpace(string(body))
	command := exec.Command("yt-dlp",
		"--format", "bestaudio",
		"--extract-audio",
		"--audio-format", "opus",
		"--audio-quality", "0",
		"-o", "%(title)s.%(ext)s",
		"--replace-in-metadata", "title", `"`, "\u201C",
		"--replace-in-metadata", "title", `"`, "\u201D",
		"--replace-in-metadata", "title", `\s*\([^)]*(audio|Audio|AUDIO|lyric|Lyric|lyrics|Lyrics|visualizer|Visualizer|video|Video)[^)]*\)\s*`, "",
		"--replace-in-metadata", "title", `\s*\[[^\]]*(audio|Audio|AUDIO|lyric|Lyric|lyrics|Lyrics|visualizer|Visualizer|video|Video)[^\]]*\]\s*`, "",
		id) // Regex then "" replaces it with "", or nothing btw
	go func() {
		err := command.Run()
		if err != nil {
			log.Println("yt-dlp failed:", err)
		}
	}()
	fmt.Printf("We are downloading %s ! \n", id)
	w.Header().Set("Content-Type", "text/plain") // send the data as plain text
	w.Write([]byte("ok"))                        // we are okay
}

func main() {
	http.HandleFunc("/download", download_page)
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", nil))
}
