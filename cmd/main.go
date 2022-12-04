package main

import (
	"net/http"

	"example.com/Todo/api"
)

func main() {
	r := api.Initialize()
	// Wrap our server with our gzip handler to gzip compress all responses.
	http.ListenAndServe(":8000", r.Router)
}
