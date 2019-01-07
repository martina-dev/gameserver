package service

import (
	"fmt"
	"log"
	"net/http"
)

//NewServer function takes in a port number as a string
func NewServer(port string) {
	fmt.Println("Starting server at port:", port)
	handler := http.HandlerFunc(GetScore)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
