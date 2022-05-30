package boxscore

import (
	"strconv"
	"strings"
)

// GetGames returns the list of games that was played from a given date
func GetGames(date string) []Game {
	scoreboard := GetScoreboardData(date)

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
func GetBoxscores(date string) []Boxscore {
	games := GetGames(date)
	var boxscores []Boxscore
	for _, g := range games {
		boxscores = append(boxscores, getBoxscoreFromGameId(date, g.ID, g.HomeTeam, g.VistorTeam))
	}
	return boxscores
}

// GetBoxscore returns the boxscore of a game between two teams that was played from a given date
func GetBoxscore(date string, homeTeam string, visitorTeam string) Boxscore {
	games := GetGames(date)
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
	boxscoreData := GetBoxscoreData(date, gameId)

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
func GetPlayers(year string) []Player {
	players := GetPlayersData(year)
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

// GetTeams returns the list of players in a given season (ex 2021 for season 2021/2022)
func GetTeams(year string) []Team {
	teams := GetTeamsData(year)
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
