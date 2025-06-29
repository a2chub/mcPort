package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

// Config represents a placeholder for server configuration.
type Config struct {
	PortRange struct {
		Start int `json:"start"`
		End   int `json:"end"`
	} `json:"port_range"`
	StorePath string `json:"store_path"`
}

func allocateHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]int{"port": 20010}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func releaseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func main() {
	configPath := flag.String("config", "mcport.yaml", "path to config file")
	addr := flag.String("addr", "127.0.0.1:8080", "listen address")
	flag.Parse()

	// TODO: load configuration from file
	fmt.Printf("using config: %s\n", *configPath)

	http.HandleFunc("/allocate", allocateHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/release", releaseHandler)
	http.HandleFunc("/services", servicesHandler)

	log.Printf("mcportd listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
