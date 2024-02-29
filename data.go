package main

import (
	"encoding/json"
	"net/http"
	"strconv"
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
	err := fetch("http://api.sr.se/api/v2/episodes?programid="+programId+"&format=json&audioquality=hi", &episodesResponse)
	return episodesResponse.Episodes, err
}

func includes(s string, query string) bool {
	return strings.Contains(strings.ToLower(s), query)
}

func filterPrograms(programs []Program, query string) []Program {
	if len(query) == 0 {
		return programs
	}
	queryLow := strings.ToLower(query)
	var filteredPrograms []Program
	for _, p := range programs {
		if includes(p.Description, queryLow) || includes(p.Name, queryLow) {
			filteredPrograms = append(filteredPrograms, p)
		}
	}
	return filteredPrograms
}

func getProgram(programs []Program, programId string) Program {
	id, err := strconv.Atoi(programId)
	if err != nil {
		return Program{}
	}
	for _, p := range programs {
		if p.ID == id {
			return p
		}
	}
	return Program{}
}
