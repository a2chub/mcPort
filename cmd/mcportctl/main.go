package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	serverAddr := flag.String("server", "http://127.0.0.1:8080", "mcportd address")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("usage: mcportctl <command> [options]")
		fmt.Println("commands: allocate, register, release, list")
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case "allocate":
		resp, err := http.Post(*serverAddr+"/allocate", "application/json", bytes.NewBuffer([]byte("{}")))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		io.Copy(os.Stdout, resp.Body)
	case "register":
		// placeholder: send empty registration
		resp, err := http.Post(*serverAddr+"/register", "application/json", bytes.NewBuffer([]byte("{}")))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		io.Copy(os.Stdout, resp.Body)
	case "release":
		resp, err := http.Post(*serverAddr+"/release", "application/json", bytes.NewBuffer([]byte("{}")))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		io.Copy(os.Stdout, resp.Body)
	case "list":
		resp, err := http.Get(*serverAddr + "/services")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		var services interface{}
		json.NewDecoder(resp.Body).Decode(&services)
		b, _ := json.MarshalIndent(services, "", "  ")
		fmt.Println(string(b))
	default:
		fmt.Println("unknown command:", cmd)
		os.Exit(1)
	}
}
