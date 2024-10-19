package main

import (
	"fmt"
	"net/http"

	"github.com/ReynerioSamos/AWT_Q2/nomiddleware/handlers"
)

func main() {
	http.Handle("/hello", &handlers.Hello{})
	http.Handle("/world", &handlers.World{})

	http.ListenAndServe(":8080", nil)
}