package boxscore

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// GetMockServer initializes the mock responses when calling data api
func GetMockServer(t *testing.T) *httptest.Server {

	mux := http.NewServeMux()
	mux.HandleFunc("scoreboard.json", func(w http.ResponseWriter, req *http.Request) {
		writeContent(t, &w, "20211225_scoreboard.json")
	})
	mux.HandleFunc("boxscore.json", func(w http.ResponseWriter, req *http.Request) {
		writeContent(t, &w, "20211225_0022100489_boxscore.json")
	})
	mux.HandleFunc("players.json", func(w http.ResponseWriter, req *http.Request) {
		writeContent(t, &w, "2021_players.json")
	})
	mux.HandleFunc("teams.json", func(w http.ResponseWriter, req *http.Request) {
		writeContent(t, &w, "2021_teams.json")
	})

	server := httptest.NewServer(mux)
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
