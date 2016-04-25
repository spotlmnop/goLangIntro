package main

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"encoding/json"
	"github.com/spotlmnop/goLangIntro/models"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func createUser(c web.C, w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var u models.User
	decoder.Decode(&u)
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(u)
}

func main() {
	goji.Get("/hello/:name", hello)
	goji.Post("/user", createUser)
	goji.Serve()
}
