package boxscore

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

const NBA_DATA_ENDPOINT = "https://data.nba.net/data/10s/prod/v1"

type ApiClient struct {
	Url string
}

func NewApiClient(url ...string) ApiClient {
	if url == nil {
		return ApiClient{Url: NBA_DATA_ENDPOINT}
	}
	return ApiClient{Url: url[0]}
}

// GetScoreboardData returns the scoreboard from a given date
func (c ApiClient) GetScoreboardData(date string) (map[string]interface{}, error) {
	endpoint := c.Url + "/" + date + "/scoreboard.json"
	scoreboard := get(endpoint)
	if scoreboard == nil {
		message := "An error occurred when fetching scoreboard"
		log.Println(message)
		return nil, errors.New(message)
	}
	return scoreboard, nil
}

// GetBoxscoreData returns the boxscore from a given date and game identifier
func (c ApiClient) GetBoxscoreData(date string, gameId string) (map[string]interface{}, error) {
	endpoint := NBA_DATA_ENDPOINT + "/" + date + "/" + gameId + "_boxscore.json"
	boxscore := get(endpoint)
	if boxscore == nil {
		message := "An error occurred when fetching boxscore"
		log.Println(message)
		return nil, errors.New(message)
	}
	return boxscore, nil
}

// GetPlayersData returns the list of players for a given season (year)
func (c ApiClient) GetPlayersData(year string) (map[string]interface{}, error) {
	endpoint := c.Url + "/" + year + "/players.json"
	players := get(endpoint)
	if players == nil {
		message := "An error occurred when fetching players"
		log.Println(message)
		return nil, errors.New(message)
	}
	return players, nil
}

// GetPlayerStatsData returns the list of statistics of a player until the given season (year)
func (c ApiClient) GetPlayerStatsData(year string, playerId string) (map[string]interface{}, error) {
	endpoint := c.Url + "/" + year + "/players/" + playerId + "_profile.json"
	stats := get(endpoint)
	if stats == nil {
		message := "An error occurred when fetching player stats"
		log.Println(message)
		return nil, errors.New(message)
	}
	return stats, nil
}

// GetTeamsData returns the list of players for a given season (year)
func (c ApiClient) GetTeamsData(year string) (map[string]interface{}, error) {
	endpoint := c.Url + "/" + year + "/teams.json"
	teams := get(endpoint)
	if teams == nil {
		message := "An error occurred when fetching teams"
		log.Println(message)
		return nil, errors.New(message)
	}
	return teams, nil
}

func get(endpoint string) map[string]interface{} {
	request, _ := http.NewRequest(http.MethodGet, endpoint, nil)
	request.Header.Add("Accept", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}

	response, err := client.Do(request)
	if response.StatusCode != http.StatusOK || err != nil {
		log.Println(err)
		return nil
	}
	defer response.Body.Close()

	data := make(map[string]interface{})
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		return nil
	}
	return data
}
