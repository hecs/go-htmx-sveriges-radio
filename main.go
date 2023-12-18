package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var programs, programsErr = fetchPrograms()

func main() {
	if programsErr != nil {
		fmt.Println("Error fetching data:", programsErr)
		os.Exit(1)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/program/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")

		if len(parts) == 3 {
			programId := parts[2]
			episodes, episodesErr := fetchEpisodes(programId)

			if episodesErr != nil {
				fmt.Println("Error fetching episodes:", episodesErr)
				http.Error(w, episodesErr.Error(), http.StatusInternalServerError)
				return
			}
			program_episodes(episodes).Render(context.Background(), w)
		} else {
			fmt.Println("Couldn't parse program url id:", parts)
			http.NotFound(w, r)
		}

	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("searchQuery")
		filtered := filterPrograms(programs.Programs, query)

		program_list_item(filtered).Render(context.Background(), w)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index(programs.Programs, programs.Copyright).Render(context.Background(), w)
	})
	port := ":8123"
	println("Listening on port http://localhost" + port)
	http.ListenAndServe(port, nil)
}
