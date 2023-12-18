package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

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
