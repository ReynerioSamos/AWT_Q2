package handlers

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (rcv Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}