package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)


func main() {
	url := "https://eowaf55nxz7tnbw.m.pipedream.net"
	resp, err := http.Get(url)
	inspectResponse(resp, err)

	// encode data to json
	data, err := json.Marshal(struct {
		X int
		Y float32
		}{X:4, Y:3.8})

	if err != nil {
		log.Fatal("Error occured while marshaling json ", err)
	}
	// convert data to I/O reader interface
	resp, err = http.Post(url, "application/json", bytes.NewReader(data))
	inspectResponse(resp, err)
	}



func inspectResponse(resp *http.Response, err error) {
	if err != nil {
		log.Fatal("Error occurred while marshaling json ", err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error occurred while trying to read http response body", err)
	}
	log.Println(string(b))
}
