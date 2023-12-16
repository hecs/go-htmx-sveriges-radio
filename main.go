package main

import (
	"context"
	"encoding/json"
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

	http.HandleFunc("/program/", programHandler)
	http.HandleFunc("/search", searchHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index(programs.Programs, programs.Copyright).Render(context.Background(), w)
	})
	port := ":8080"
	println("Listening on port http://localhost" + port)
	http.ListenAndServe(port, nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("searchQuery")
	filtered := filterPrograms(programs.Programs, query)

	program_list_item(filtered).Render(context.Background(), w)
}

func programHandler(w http.ResponseWriter, r *http.Request) {
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

}

func fetchPrograms() (*PageData, error) {
	resp, err := http.Get("http://api.sr.se/api/v2/programs?format=json&size=20&programcategoryid=14")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pageData := PageData{}
	err = json.NewDecoder(resp.Body).Decode(&pageData)

	return &pageData, err
}

func fetchEpisodes(programId string) ([]Episode, error) {
	resp, err := http.Get("http://api.sr.se/api/v2/episodes/index?programid=" + programId + "&format=json&audioquality=hi")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	episodesResponse := EpisodesResponse{}
	err = json.NewDecoder(resp.Body).Decode(&episodesResponse)

	return episodesResponse.Episodes, err
}

func filterPrograms(programs []Program, query string) []Program {
	query = strings.ToLower(query)
	var filteredPrograms []Program
	if len(query) == 0 {
		filteredPrograms = programs
	} else {
		for _, p := range programs {
			if strings.Contains(strings.ToLower(p.Description), query) || strings.Contains(strings.ToLower(p.Name), query) {
				filteredPrograms = append(filteredPrograms, p)
			}
		}
	}

	return filteredPrograms
}
