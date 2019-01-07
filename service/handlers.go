package service

import (
	"fmt"
	"net/http"
)

//GetScore returns player's score
func GetScore(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "20")
}
