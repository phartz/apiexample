package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type State struct {
	State string `json:"state"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/state", GetState)
	router.HandleFunc("/remotestate/{remoteIp}", RemoteState)

	port := ":3000"
	if len(os.Args) > 1 {
		port = fmt.Sprintf(":%s", os.Args[1])
	}

	log.Fatal(http.ListenAndServe(port, router))
}

func GetState(w http.ResponseWriter, r *http.Request) {
	state := State{State: "ok"}
	json.NewEncoder(w).Encode(state)
}

func RemoteState(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var state State

	remoteIP := vars["remoteIp"]

	err := getJSON(fmt.Sprintf("http://%s/state", remoteIP), &state)
	if err != nil {
		state.State = err.Error()
	}

	json.NewEncoder(w).Encode(state)
}

func getJSON(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}

	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
