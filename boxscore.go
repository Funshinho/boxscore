package boxscore

import "strconv"

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

func GetBoxscore(date string, homeTeam string, visitorTeam string) Boxscore {
	games := GetGames(date)
	var gameId string
	for _, g := range games {
		if g.HomeTeam == homeTeam && g.VistorTeam == visitorTeam {
			gameId = g.ID
			break
		}
	}
	boxscoreData := GetBoxscoreData(date, gameId)

	statsArray := boxscoreData["stats"].(map[string]interface{})["activePlayers"].([]interface{})
	var stats []Stats
	for _, s := range statsArray {
		points, _ := strconv.Atoi(s.(map[string]interface{})["points"].(string))
		rebounds, _ := strconv.Atoi(s.(map[string]interface{})["totReb"].(string))
		assists, _ := strconv.Atoi(s.(map[string]interface{})["assists"].(string))
		playerStats := Stats{
			PlayerID:  s.(map[string]interface{})["personId"].(string),
			FirstName: s.(map[string]interface{})["firstName"].(string),
			LastName:  s.(map[string]interface{})["lastName"].(string),
			TeamID:    s.(map[string]interface{})["teamId"].(string),
			Points:    points,
			Rebounds:  rebounds,
			Assists:   assists,
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
