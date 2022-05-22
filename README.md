# Boxscore

Go client to fetch NBA player info and statistics from https://data.nba.net

## Usage

### Installation

```shell
go get github.com/Funshinho/boxscore/v1
```

### Importing

```go
 import "github.com/Funshinho/boxscore/v1"
```

### Get boxscores

```go
boxscores := boxscore.GetBoxscores("20220101")
```
