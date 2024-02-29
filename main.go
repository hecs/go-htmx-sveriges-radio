package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

var programs, programsErr = fetchPrograms()

func main() {
	if programsErr != nil {
		fmt.Println("Error fetching data:", programsErr)
		os.Exit(1)
	}

	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("GET /program/{pId}", HTML(programHandler))
	http.HandleFunc("GET /search", HTML(searchHandler))
	http.HandleFunc("GET /{$}", HTML(indexHandler))
	port := ":8123"
	println("Listening on port http://localhost" + port)
	http.ListenAndServe(port, nil)
}

func programHandler(w http.ResponseWriter, r *http.Request) {
	programId := r.PathValue("pId")
	episodes, episodesErr := fetchEpisodes(programId)
	if episodesErr != nil {
		fmt.Println("Error fetching episodes:", episodesErr)
		http.Error(w, episodesErr.Error(), http.StatusInternalServerError)
		return
	}
	if r.Header.Get("HX-Request") == "" {
		single := []Program{getProgram(programs.Programs, programId)}
		if len(single) == 1 {
			episodes := program_episodes(episodes)
			index(program_list_item(single), episodes, programs.Copyright).Render(r.Context(), w)
		}
	} else {
		program_episodes(episodes).Render(r.Context(), w)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("searchQuery")
	filtered := filterPrograms(programs.Programs, query)
	program_list_item(filtered).Render(r.Context(), w)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index(program_list_item(programs.Programs), templ.NopComponent, programs.Copyright).Render(r.Context(), w)
}

func HTML(HandleFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Content-Security-Policy", "default-src 'none'; script-src 'self' 'unsafe-inline' unpkg.com; media-src http://sverigesradio.se *.sr.se; style-src 'self' 'unsafe-inline'; image-src 'self' data:; connect-src 'self'; frame-src 'self'; frame-ancestors 'self'; block-all-mixed-content; upgrade-insecure-requests;")
		HandleFunc(w, r)
	}
}
