package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var programs, programsErr = fetchPrograms()
var tmpl, templateErr = template.ParseFiles("programs.html", "partial_program.html")

func main() {
	if programsErr != nil {
		fmt.Println("Error fetching data:", programsErr)
		os.Exit(1)
	}
	if templateErr != nil {
		fmt.Println("Error parsing templates:", templateErr)
		os.Exit(1)
	}

	http.HandleFunc("/audio", programAudioHandler)
	http.HandleFunc("/program/", programHandler)
	http.HandleFunc("/search", searchHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "programs.html", programs)
	if err != nil {
		fmt.Println("Error rendering template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("searchQuery")
	filtered := filterPrograms(programs.Programs, query)
	fmt.Printf("#### filter length\n%s\n", len(filtered))

	err := tmpl.ExecuteTemplate(w, "program_list_item", filtered)
	if err != nil {
		fmt.Println("Error parsing template program_list_item:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func programAudioHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "program_audio", nil)
	if err != nil {
		fmt.Println("Error parsing template program_audio:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
		err := tmpl.ExecuteTemplate(w, "program_episodes", episodes)
		if err != nil {
			fmt.Println("Error parsing template:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
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
