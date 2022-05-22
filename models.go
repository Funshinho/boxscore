package boxscore

type Position string

const (
	Forward Position = "F"
	Center           = "C"
	Guard            = "G"
)

type Player struct {
	ID        string
	TeamID    string
	FirstName string
	LastName  string
	Position  Position
}

type Stats struct {
	PlayerID  string
	FirstName string
	LastName  string
	TeamID    string
	Points    int
	Rebounds  int
	Assists   int
}

type Boxscore struct {
	HomeTeam   string
	VistorTeam string
	StatsList  []Stats
}

type Game struct {
	ID         string
	HomeTeam   string
	VistorTeam string
}
