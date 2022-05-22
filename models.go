package boxscore

type Position string

const (
	Forward Position = "F"
	Center           = "C"
	Guard            = "G"
)

// Player represents the player information
type Player struct {
	ID        string
	TeamID    string
	FirstName string
	LastName  string
	Position  Position
}

// Stats represents the statistics of a player
type Stats struct {
	PlayerID  string
	FirstName string
	LastName  string
	TeamID    string
	Points    int
	Rebounds  int
	Assists   int
}

// Boxscore represents the boxscore of a game
type Boxscore struct {
	HomeTeam   string
	VistorTeam string
	StatsList  []Stats
}

// Game represents the teams that were against
type Game struct {
	ID         string
	HomeTeam   string
	VistorTeam string
}
