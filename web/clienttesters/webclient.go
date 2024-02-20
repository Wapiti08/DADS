package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)


func main() {
	// url := "https://eowaf55nxz7tnbw.m.pipedream.net"
	
	// achieve Get method
	// resp, err := http.Get(url)
	// inspectResponse(resp, err)

	// achieve Post method
	// // encode data to json
	// data, err := json.Marshal(struct {
	// 	X int
	// 	Y float32
	// 	}{X:4, Y:3.8})

	// if err != nil {
	// 	log.Fatal("Error occured while marshaling json ", err)
	// }
	// // convert data to I/O reader interface
	// resp, err = http.Post(url, "application/json", bytes.NewReader(data))
	// inspectResponse(resp, err)

	// achieve new request method with custom parameters
	// client := http.Client{
	// 	Timeout: 3 * time.Minute,
	// }

	// client.Get(url)

	// req, err := http.NewRequest(http.MethodPut, url, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// req.Header.Add("x-testheader", "learning go header")
	// req.Header.Set("User-Agent", "Go learning HTTP/1.1")

	// resp, err := client.Do(req)
	// inspectResponse(resp, err)

	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// define the return type
	v := struct {
		IP string `json:"ip"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(v.IP)
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
