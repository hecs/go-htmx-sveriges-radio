package main

import (
	"fmt"
	"net/http"
	"os"
)

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

var programs, programsErr = fetchPrograms()

func programHandler(w http.ResponseWriter, r *http.Request) {
	programId := r.PathValue("pId")
	setHeaders(w)
	episodes, episodesErr := fetchEpisodes(programId)
	if episodesErr != nil {
		fmt.Println("Error fetching episodes:", episodesErr)
		http.Error(w, episodesErr.Error(), http.StatusInternalServerError)
		return
	}
	program_episodes(episodes).Render(r.Context(), w)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	query := r.URL.Query().Get("searchQuery")
	filtered := filterPrograms(programs.Programs, query)
	program_list_item(filtered).Render(r.Context(), w)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	index(programs.Programs, programs.Copyright).Render(r.Context(), w)
}

func main() {
	if programsErr != nil {
		fmt.Println("Error fetching data:", programsErr)
		os.Exit(1)
	}

	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("GET /program/{pId}", programHandler)
	http.HandleFunc("GET /search", searchHandler)
	http.HandleFunc("GET /{$}", indexHandler)
	port := ":8123"
	println("Listening on port http://localhost" + port)
	http.ListenAndServe(port, nil)
}
