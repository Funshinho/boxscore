# Boxscore

Go client to fetch NBA players and teams info, as well as game statistics from https://data.nba.net

## Usage

### Installation

```shell
go get github.com/Funshinho/boxscore
```

### Importing

```go
 import "github.com/Funshinho/boxscore"
```

### Usage

```go
boxscores := boxscore.GetBoxscores("20220101")
players := boxscore.GetPlayers("2021")
teams := boxscore.GetTeams("2021")
```
