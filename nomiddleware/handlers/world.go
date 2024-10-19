package handlers

import (
	"fmt"
	"net/http"
)

type World struct{}

func (rcv World) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, " World!")
}