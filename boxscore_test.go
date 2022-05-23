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

	player1 := boxscore.StatsList[1]
	assert.Equal("1628369", player1.PlayerID)
	assert.Equal("Jayson", player1.FirstName)
	assert.Equal("Tatum", player1.LastName)
	assert.Equal(25, player1.Points)
	assert.Equal(9, player1.Rebounds)
	assert.Equal(4, player1.Assists)
	assert.Equal(2, player1.Steals)
	assert.Equal(2, player1.Blocks)
	assert.Equal(4, player1.Turnovers)
	assert.Equal(3, player1.Fouls)
	assert.Equal(7, player1.FGM)
	assert.Equal(20, player1.FGA)
	assert.Equal(4, player1.TPM)
	assert.Equal(10, player1.TPA)
	assert.Equal(7, player1.FTM)
	assert.Equal(8, player1.FTA)

	player2 := boxscore.StatsList[13]
	assert.Equal("203507", player2.PlayerID)
	assert.Equal("Giannis", player2.FirstName)
	assert.Equal("Antetokounmpo", player2.LastName)
	assert.Equal(36, player2.Points)
	assert.Equal(12, player2.Rebounds)
	assert.Equal(5, player2.Assists)
	assert.Equal(0, player2.Steals)
	assert.Equal(2, player2.Blocks)
	assert.Equal(2, player2.Turnovers)
	assert.Equal(4, player2.Fouls)
	assert.Equal(13, player2.FGM)
	assert.Equal(23, player2.FGA)
	assert.Equal(0, player2.TPM)
	assert.Equal(2, player2.TPA)
	assert.Equal(10, player2.FTM)
	assert.Equal(15, player2.FTA)
}
