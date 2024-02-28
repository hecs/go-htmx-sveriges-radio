package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func fetch(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func fetchPrograms() (*PageData, error) {
	pageData := PageData{}
	err := fetch("http://api.sr.se/api/v2/programs?format=json&size=20&programcategoryid=14", &pageData)
	return &pageData, err
}

func fetchEpisodes(programId string) ([]Episode, error) {
	episodesResponse := EpisodesResponse{}
	err := fetch("http://api.sr.se/api/v2/episodes/index?programid="+programId+"&format=json&audioquality=hi", &episodesResponse)
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
