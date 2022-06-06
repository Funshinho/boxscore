package boxscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetScoreboardData(t *testing.T) {
	server := GetMockServer(t)
	defer server.Close()
	client := NewApiClient(server.URL)

	result, _ := client.GetScoreboardData("2021")
	assert := assert.New(t)
	assert.Equal(5, int(result["numGames"].(float64)))
	assert.Equal("BOS", result["games"].([]interface{})[1].(map[string]interface{})["vTeam"].(map[string]interface{})["triCode"])
	assert.Equal("MIL", result["games"].([]interface{})[1].(map[string]interface{})["hTeam"].(map[string]interface{})["triCode"])
}

func TestGetBoxscoreData(t *testing.T) {
	server := GetMockServer(t)
	defer server.Close()
	client := NewApiClient(server.URL)

	result, _ := client.GetBoxscoreData("20211225", "0022100489")
	assert := assert.New(t)
	assert.Equal("Tatum", result["stats"].(map[string]interface{})["activePlayers"].([]interface{})[1].(map[string]interface{})["lastName"])
	assert.Equal("25", result["stats"].(map[string]interface{})["activePlayers"].([]interface{})[1].(map[string]interface{})["points"])

	assert.Equal("Antetokounmpo", result["stats"].(map[string]interface{})["activePlayers"].([]interface{})[13].(map[string]interface{})["lastName"])
	assert.Equal("36", result["stats"].(map[string]interface{})["activePlayers"].([]interface{})[13].(map[string]interface{})["points"])
}

func TestGetPlayersData(t *testing.T) {
	server := GetMockServer(t)
	defer server.Close()
	client := NewApiClient(server.URL)

	result, _ := client.GetPlayersData("2021")
	assert := assert.New(t)
	assert.Equal("Precious", result["league"].(map[string]interface{})["standard"].([]interface{})[0].(map[string]interface{})["firstName"])
	assert.Equal("Achiuwa", result["league"].(map[string]interface{})["standard"].([]interface{})[0].(map[string]interface{})["lastName"])

	assert.Equal("Bam", result["league"].(map[string]interface{})["standard"].([]interface{})[2].(map[string]interface{})["firstName"])
	assert.Equal("Adebayo", result["league"].(map[string]interface{})["standard"].([]interface{})[2].(map[string]interface{})["lastName"])
}

func TestGetPlayerStatsData(t *testing.T) {
	server := GetMockServer(t)
	defer server.Close()
	client := NewApiClient(server.URL)

	result, _ := client.GetPlayerStatsData("2021", "203507")
	assert := assert.New(t)
	stats := result["league"].(map[string]interface{})["standard"].(map[string]interface{})["stats"].(map[string]interface{})
	seasons := stats["regularSeason"].(map[string]interface{})["season"].([]interface{})
	assert.Equal("29.9", seasons[0].(map[string]interface{})["total"].(map[string]interface{})["ppg"])
	assert.Equal("11.6", seasons[0].(map[string]interface{})["total"].(map[string]interface{})["rpg"])
	assert.Equal("5.8", seasons[0].(map[string]interface{})["total"].(map[string]interface{})["apg"])

	assert.Equal("28.1", seasons[1].(map[string]interface{})["total"].(map[string]interface{})["ppg"])
	assert.Equal("11", seasons[1].(map[string]interface{})["total"].(map[string]interface{})["rpg"])
	assert.Equal("5.9", seasons[1].(map[string]interface{})["total"].(map[string]interface{})["apg"])
}

func TestGetTeamsData(t *testing.T) {
	server := GetMockServer(t)
	defer server.Close()
	client := NewApiClient(server.URL)

	result, _ := client.GetTeamsData("2021")
	assert := assert.New(t)
	assert.Equal("Atlanta Hawks", result["league"].(map[string]interface{})["standard"].([]interface{})[0].(map[string]interface{})["fullName"])
	assert.Equal("ATL", result["league"].(map[string]interface{})["standard"].([]interface{})[0].(map[string]interface{})["tricode"])

	assert.Equal("Miami Heat", result["league"].(map[string]interface{})["standard"].([]interface{})[15].(map[string]interface{})["fullName"])
	assert.Equal("MIA", result["league"].(map[string]interface{})["standard"].([]interface{})[15].(map[string]interface{})["tricode"])
}
