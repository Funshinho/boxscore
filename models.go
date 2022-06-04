package boxscore

type Position string

const (
	Forward Position = "F"
	Center  Position = "C"
	Guard   Position = "G"
)

// Player represents the team information
type Team struct {
	ID      string
	Name    string
	Tricode string
}

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

type AverageStats struct {
	PlayerID         string
	PointsPerGame    float64
	ReboundsPerGame  float64
	AssistsPerGame   float64
	BlocksPerGame    float64
	StealsPerGame    float64
	TurnoversPerGame float64
	FGP              float64 // Field goal percentage
	TPP              float64 // Three points percentage
	FTP              float64 // Free throw percentage
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
