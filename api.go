package boxscore

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const NBA_DATA_ENDPOINT = "https://data.nba.net/data/10s/prod/v1"

// GetScoreboardData returns the scoreboard from a given date
func GetScoreboardData(date string) map[string]interface{} {
	endpoint := NBA_DATA_ENDPOINT + "/" + date + "/scoreboard.json"
	scoreboard := get(endpoint)
	return scoreboard
}

// GetBoxscoreData returns the boxscore from a given date and game identifier
func GetBoxscoreData(date string, gameId string) map[string]interface{} {
	endpoint := NBA_DATA_ENDPOINT + "/" + date + "/" + gameId + "_boxscore.json"
	boxscore := get(endpoint)
	return boxscore
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
