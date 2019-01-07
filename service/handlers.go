package service

import (
	"fmt"
	"net/http"
)

//GetPlayer returns player's score
func GetPlayer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	fmt.Fprintf(w, GetPlayerScore(player))
}

//GetPlayerScore matches name to player and returns player's score
func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}
