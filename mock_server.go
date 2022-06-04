package boxscore

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// GetMockServer initializes the mock responses when calling data api
func GetMockServer(t *testing.T) *httptest.Server {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if strings.Contains(req.URL.Path, "scoreboard.json") {
			writeContent(t, &w, "20211225_scoreboard.json")
		}
		if strings.Contains(req.URL.Path, "boxscore.json") {
			writeContent(t, &w, "20211225_0022100489_boxscore.json")
		}
		if strings.Contains(req.URL.Path, "players.json") {
			writeContent(t, &w, "2021_players.json")
		}
		if strings.Contains(req.URL.Path, "teams.json") {
			writeContent(t, &w, "2021_teams.json")
		}
		if strings.Contains(req.URL.Path, "profile.json") {
			writeContent(t, &w, "203507_profile.json")
		}
	}))
	return server
}

func writeContent(t *testing.T, w *http.ResponseWriter, filename string) {
	content, err := ioutil.ReadFile("mocks/" + filename)
	if err != nil {
		t.Fatalf("Could not read file: " + filename)
	}
	(*w).WriteHeader(http.StatusOK)
	(*w).Write(content)
}
