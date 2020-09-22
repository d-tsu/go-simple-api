package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Beer is a beer.
type Beer struct {
	Name string `json:"name"`
}

func beerHandler(w http.ResponseWriter, r *http.Request) {
	beer := Beer{
		Name: "kirin",
	}
	res, _ := json.Marshal(beer)
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// get
		w.Write(res)
	case http.MethodPost:
		// post
		w.Write(res)
	default:
		fmt.Println("error")
	}
}

func main() {
	http.HandleFunc("/beer", beerHandler)
	http.ListenAndServe(":80", nil)
}
