package main

import (
	"net/http"

	"github.com/ReynerioSamos/AWT_Q2/middleware/handlers"
	"github.com/ReynerioSamos/AWT_Q2/middleware/middleware"
)

func main() {
	http.Handle("/hello", middleware.Log(&handlers.Hello{}))
	http.Handle("/world", middleware.Log(&handlers.World{}))

	http.ListenAndServe(":8080", nil)
}