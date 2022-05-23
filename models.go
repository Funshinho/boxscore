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
	Blocks    int
	Steals    int
	Turnovers int
	Fouls     int
	FGM       int // Field goal made
	FGA       int // Field goal attempted
	TPM       int // Three points made
	TPA       int // Three points attempted
	FTM       int // Free throw made
	FTA       int // Free throw attempted
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
