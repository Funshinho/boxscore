package boxscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGames(t *testing.T) {
	server := GetMockServer(t)
	defer server.Close()

	games := GetGames("20211225")
	assert := assert.New(t)
	assert.Equal(5, len(games))

	assert.Equal("0022100488", games[0].ID)
	assert.Equal("NYK", games[0].HomeTeam)
	assert.Equal("ATL", games[0].VistorTeam)

	assert.Equal("0022100489", games[1].ID)
	assert.Equal("MIL", games[1].HomeTeam)
	assert.Equal("BOS", games[1].VistorTeam)
}

func TestGetBoxscores(t *testing.T) {
	server := GetMockServer(t)
	defer server.Close()

	boxscores := GetBoxscores("20211225")
	assert := assert.New(t)

	assert.Equal(5, len(boxscores))

	assert.Equal("NYK", boxscores[0].HomeTeam)
	assert.Equal("ATL", boxscores[0].VistorTeam)
	assert.Equal("MIL", boxscores[1].HomeTeam)
	assert.Equal("BOS", boxscores[1].VistorTeam)

}

func TestGetBoxscore(t *testing.T) {
	server := GetMockServer(t)
	defer server.Close()

	boxscore := GetBoxscore("20211225", "MIL", "BOS")
	assert := assert.New(t)

	assert.Equal("MIL", boxscore.HomeTeam)
	assert.Equal("BOS", boxscore.VistorTeam)

	assert.Equal("1628369", boxscore.StatsList[1].PlayerID)
	assert.Equal("Jayson", boxscore.StatsList[1].FirstName)
	assert.Equal("Tatum", boxscore.StatsList[1].LastName)
	assert.Equal(25, boxscore.StatsList[1].Points)
	assert.Equal(9, boxscore.StatsList[1].Rebounds)
	assert.Equal(4, boxscore.StatsList[1].Assists)

	assert.Equal("203507", boxscore.StatsList[13].PlayerID)
	assert.Equal("Giannis", boxscore.StatsList[13].FirstName)
	assert.Equal("Antetokounmpo", boxscore.StatsList[13].LastName)
	assert.Equal(36, boxscore.StatsList[13].Points)
	assert.Equal(12, boxscore.StatsList[13].Rebounds)
	assert.Equal(5, boxscore.StatsList[13].Assists)
}
