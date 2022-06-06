package boxscore

import (
	"strconv"
	"strings"
)

type Client struct {
	ApiClient ApiClient
}

func NewClient(url ...string) Client {
	if url == nil {
		return Client{ApiClient: NewApiClient()}
	}
	return Client{ApiClient: NewApiClient(url[0])}
}

// GetGames returns the list of games that was played from a given date
func (c Client) GetGames(date string) []Game {
	scoreboard := c.ApiClient.GetScoreboardData(date)

	gamesArray := scoreboard["games"].([]interface{})
	var games []Game
	for _, g := range gamesArray {
		game := Game{
			ID:         g.(map[string]interface{})["gameId"].(string),
			HomeTeam:   g.(map[string]interface{})["hTeam"].(map[string]interface{})["triCode"].(string),
			VistorTeam: g.(map[string]interface{})["vTeam"].(map[string]interface{})["triCode"].(string),
		}
		games = append(games, game)
	}
	return games
}

// GetBoxscores returns the list of boxscores that was played from a given date
func (c Client) GetBoxscores(date string) []Boxscore {
	games := c.GetGames(date)
	var boxscores []Boxscore
	for _, g := range games {
		boxscores = append(boxscores, getBoxscoreFromGameId(date, g.ID, g.HomeTeam, g.VistorTeam))
	}
	return boxscores
}

// GetBoxscore returns the boxscore of a game between two teams that was played from a given date
func (c Client) GetBoxscore(date string, homeTeam string, visitorTeam string) Boxscore {
	games := c.GetGames(date)
	var gameId string
	for _, g := range games {
		if g.HomeTeam == homeTeam && g.VistorTeam == visitorTeam {
			gameId = g.ID
			break
		}
	}
	return getBoxscoreFromGameId(date, gameId, homeTeam, visitorTeam)
}

func getBoxscoreFromGameId(date string, gameId string, homeTeam string, visitorTeam string) Boxscore {
	client := NewApiClient()
	boxscoreData := client.GetBoxscoreData(date, gameId)

	statsArray := boxscoreData["stats"].(map[string]interface{})["activePlayers"].([]interface{})
	var stats []Stats
	for _, s := range statsArray {
		points, _ := strconv.Atoi(s.(map[string]interface{})["points"].(string))
		rebounds, _ := strconv.Atoi(s.(map[string]interface{})["totReb"].(string))
		assists, _ := strconv.Atoi(s.(map[string]interface{})["assists"].(string))
		blocks, _ := strconv.Atoi(s.(map[string]interface{})["blocks"].(string))
		steals, _ := strconv.Atoi(s.(map[string]interface{})["steals"].(string))
		turnovers, _ := strconv.Atoi(s.(map[string]interface{})["turnovers"].(string))
		fouls, _ := strconv.Atoi(s.(map[string]interface{})["pFouls"].(string))
		fgm, _ := strconv.Atoi(s.(map[string]interface{})["fgm"].(string))
		fga, _ := strconv.Atoi(s.(map[string]interface{})["fga"].(string))
		tpm, _ := strconv.Atoi(s.(map[string]interface{})["tpm"].(string))
		tpa, _ := strconv.Atoi(s.(map[string]interface{})["tpa"].(string))
		ftm, _ := strconv.Atoi(s.(map[string]interface{})["ftm"].(string))
		fta, _ := strconv.Atoi(s.(map[string]interface{})["fta"].(string))
		playerStats := Stats{
			PlayerID:  s.(map[string]interface{})["personId"].(string),
			FirstName: s.(map[string]interface{})["firstName"].(string),
			LastName:  s.(map[string]interface{})["lastName"].(string),
			TeamID:    s.(map[string]interface{})["teamId"].(string),
			Points:    points,
			Rebounds:  rebounds,
			Assists:   assists,
			Blocks:    blocks,
			Steals:    steals,
			Turnovers: turnovers,
			Fouls:     fouls,
			FGM:       fgm,
			FGA:       fga,
			TPM:       tpm,
			TPA:       tpa,
			FTM:       ftm,
			FTA:       fta,
		}
		stats = append(stats, playerStats)
	}
	boxscore := Boxscore{
		HomeTeam:   homeTeam,
		VistorTeam: visitorTeam,
		StatsList:  stats,
	}
	return boxscore
}

// GetPlayers returns the list of players in a given season (ex 2021 for season 2021/2022)
func (c Client) GetPlayers(year string) []Player {
	players := c.ApiClient.GetPlayersData(year)
	return mapPlayers(players)
}

func mapPlayers(playersData map[string]interface{}) []Player {
	playersArray := playersData["league"].(map[string]interface{})["standard"].([]interface{})
	var players []Player
	for _, p := range playersArray {
		var position Position
		positionData := p.(map[string]interface{})["pos"].(string)
		if strings.HasPrefix(positionData, "G") {
			position = Guard
		} else if strings.HasPrefix(positionData, "F") {
			position = Forward
		} else {
			position = Center
		}

		player := Player{
			ID:        p.(map[string]interface{})["personId"].(string),
			FirstName: p.(map[string]interface{})["firstName"].(string),
			LastName:  p.(map[string]interface{})["lastName"].(string),
			TeamID:    p.(map[string]interface{})["teamId"].(string),
			Position:  position,
		}
		players = append(players, player)
	}
	return players
}

// GetPlayerStats returns the average stats of a player for the given season (ex 2021 for season 2021/2022)
func (c Client) GetPlayerStats(year string, playerId string) AverageStats {
	playerStatsData := c.ApiClient.GetPlayerStatsData(year, playerId)
	stats := playerStatsData["league"].(map[string]interface{})["standard"].(map[string]interface{})["stats"].(map[string]interface{})
	return mapStats(stats, year, playerId)
}

func mapStats(stats map[string]interface{}, year string, playerId string) AverageStats {
	seasons := stats["regularSeason"].(map[string]interface{})["season"].([]interface{})
	var season map[string]interface{}
	yearf, _ := strconv.ParseFloat(year, 64)
	for _, s := range seasons {
		if s.(map[string]interface{})["seasonYear"].(float64) == yearf {
			season = s.(map[string]interface{})
			break
		}
	}
	if season == nil {
		return AverageStats{
			PlayerID:         playerId,
			PointsPerGame:    0,
			ReboundsPerGame:  0,
			AssistsPerGame:   0,
			BlocksPerGame:    0,
			StealsPerGame:    0,
			TurnoversPerGame: 0,
			FGP:              0,
			TPP:              0,
			FTP:              0,
		}
	}
	ppg, _ := strconv.ParseFloat(season["total"].(map[string]interface{})["ppg"].(string), 64)
	rpg, _ := strconv.ParseFloat(season["total"].(map[string]interface{})["rpg"].(string), 64)
	apg, _ := strconv.ParseFloat(season["total"].(map[string]interface{})["apg"].(string), 64)
	bpg, _ := strconv.ParseFloat(season["total"].(map[string]interface{})["bpg"].(string), 64)
	spg, _ := strconv.ParseFloat(season["total"].(map[string]interface{})["spg"].(string), 64)
	topg, _ := strconv.ParseFloat(season["total"].(map[string]interface{})["topg"].(string), 64)
	fgp, _ := strconv.ParseFloat(season["total"].(map[string]interface{})["fgp"].(string), 64)
	tpp, _ := strconv.ParseFloat(season["total"].(map[string]interface{})["tpp"].(string), 64)
	ftp, _ := strconv.ParseFloat(season["total"].(map[string]interface{})["ftp"].(string), 64)
	playerStats := AverageStats{
		PlayerID:         playerId,
		PointsPerGame:    ppg,
		ReboundsPerGame:  rpg,
		AssistsPerGame:   apg,
		BlocksPerGame:    bpg,
		StealsPerGame:    spg,
		TurnoversPerGame: topg,
		FGP:              fgp,
		TPP:              tpp,
		FTP:              ftp,
	}
	return playerStats
}

// GetTeams returns the list of players in a given season (ex 2021 for season 2021/2022)
func (c Client) GetTeams(year string) []Team {
	teams := c.ApiClient.GetTeamsData(year)
	return mapTeams(teams)
}

func mapTeams(teamsData map[string]interface{}) []Team {
	teamsArray := teamsData["league"].(map[string]interface{})["standard"].([]interface{})
	var teams []Team
	for _, t := range teamsArray {
		team := Team{
			ID:      t.(map[string]interface{})["teamId"].(string),
			Name:    t.(map[string]interface{})["fullName"].(string),
			Tricode: t.(map[string]interface{})["tricode"].(string),
		}
		teams = append(teams, team)
	}
	return teams
}
