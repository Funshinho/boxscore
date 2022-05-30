package boxscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetScoreboardData(t *testing.T) {
	server := GetMockServer(t)
	defer server.Close()

	result := GetScoreboardData("20211225")
	assert := assert.New(t)
	assert.Equal(5, int(result["numGames"].(float64)))
	assert.Equal("BOS", result["games"].([]interface{})[1].(map[string]interface{})["vTeam"].(map[string]interface{})["triCode"])
	assert.Equal("MIL", result["games"].([]interface{})[1].(map[string]interface{})["hTeam"].(map[string]interface{})["triCode"])
}

func TestGetBoxscoreData(t *testing.T) {
	server := GetMockServer(t)
	defer server.Close()

	result := GetBoxscoreData("20211225", "0022100489")
	assert := assert.New(t)
	assert.Equal("Tatum", result["stats"].(map[string]interface{})["activePlayers"].([]interface{})[1].(map[string]interface{})["lastName"])
	assert.Equal("25", result["stats"].(map[string]interface{})["activePlayers"].([]interface{})[1].(map[string]interface{})["points"])

	assert.Equal("Antetokounmpo", result["stats"].(map[string]interface{})["activePlayers"].([]interface{})[13].(map[string]interface{})["lastName"])
	assert.Equal("36", result["stats"].(map[string]interface{})["activePlayers"].([]interface{})[13].(map[string]interface{})["points"])
}

func TestGetPlayersData(t *testing.T) {
	server := GetMockServer(t)
	defer server.Close()

	result := GetPlayersData("2021")
	assert := assert.New(t)
	assert.Equal("Precious", result["league"].(map[string]interface{})["standard"].([]interface{})[0].(map[string]interface{})["firstName"])
	assert.Equal("Achiuwa", result["league"].(map[string]interface{})["standard"].([]interface{})[0].(map[string]interface{})["lastName"])

	assert.Equal("Bam", result["league"].(map[string]interface{})["standard"].([]interface{})[2].(map[string]interface{})["firstName"])
	assert.Equal("Adebayo", result["league"].(map[string]interface{})["standard"].([]interface{})[2].(map[string]interface{})["lastName"])
}

func TestGetTeamsData(t *testing.T) {
	server := GetMockServer(t)
	defer server.Close()

	result := GetTeamsData("2021")
	assert := assert.New(t)
	assert.Equal("Atlanta Hawks", result["league"].(map[string]interface{})["standard"].([]interface{})[0].(map[string]interface{})["fullName"])
	assert.Equal("ATL", result["league"].(map[string]interface{})["standard"].([]interface{})[0].(map[string]interface{})["tricode"])

	assert.Equal("Miami Heat", result["league"].(map[string]interface{})["standard"].([]interface{})[15].(map[string]interface{})["fullName"])
	assert.Equal("MIA", result["league"].(map[string]interface{})["standard"].([]interface{})[15].(map[string]interface{})["tricode"])
}
