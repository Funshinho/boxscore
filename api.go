package boxscore

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
func (c ApiClient) GetScoreboardData(date string) map[string]interface{} {
	endpoint := c.Url + "/" + date + "/scoreboard.json"
	scoreboard := get(endpoint)
	return scoreboard
}

// GetBoxscoreData returns the boxscore from a given date and game identifier
func (c ApiClient) GetBoxscoreData(date string, gameId string) map[string]interface{} {
	endpoint := NBA_DATA_ENDPOINT + "/" + date + "/" + gameId + "_boxscore.json"
	boxscore := get(endpoint)
	return boxscore
}

// GetPlayersData returns the list of players for a given season (year)
func (c ApiClient) GetPlayersData(year string) map[string]interface{} {
	endpoint := c.Url + "/" + year + "/players.json"
	players := get(endpoint)
	return players
}

// GetPlayerStatsData returns the list of statistics of a player until the given season (year)
func (c ApiClient) GetPlayerStatsData(year string, playerId string) map[string]interface{} {
	endpoint := c.Url + "/" + year + "/players/" + playerId + "_profile.json"
	stats := get(endpoint)
	return stats
}

// GetTeamsData returns the list of players for a given season (year)
func (c ApiClient) GetTeamsData(year string) map[string]interface{} {
	endpoint := c.Url + "/" + year + "/teams.json"
	teams := get(endpoint)
	return teams
}

func get(endpoint string) map[string]interface{} {
	request, _ := http.NewRequest(http.MethodGet, endpoint, nil)
	request.Header.Add("Accept", "application/json")
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return nil
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	defer response.Body.Close()
	data := make(map[string]interface{})
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Println(err)
		return nil
	}
	return data
}
