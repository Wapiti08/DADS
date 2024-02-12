package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	op := flag.String("type", "", "Server (s) or client (c) ?")
	address := flag.String("addr",":8000", "address? host:port ")
	flag.Parse()

	switch strings.ToUpper(*op) {
	case "S":
		runUDPServer(*address)
	case "C":
		runUDPClient(*address)
	}
}

func runUDPClient(address string) error {
	conn, err := net.Dial("udp", address)
	if err != nil {
		log.Fatal("connecting to %s fail", address)
	}
	defer conn.Close()
	// read file and return its bytes type
	file, err := os.Open("inputfile.csv")
	if err!=nil{
		fmt.Println(err)
	}
	filebytes, err := io.ReadAll(file)
	defer file.Close() 
	// define the encoding length
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(filebytes)))
	base64.StdEncoding.Encode(dst, filebytes)
	// encode with fixed length
	log.Println("Sending ", len(dst), " bytes")
	// write data to server
	_, err = conn.Write(dst)
	return err

}

func runUDPServer(address string) error {
	pc, err := net.ListenPacket("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	buffer := make([]byte, 4096)
	fmt.Println("Listening ....")

	n, _, err := pc.ReadFrom(buffer)
	if err != nil {
		log.Fatal(err)
	}
	dst := make([]byte, base64.StdEncoding.DecodedLen(n))
	// control the size of real context part
	_, err = base64.StdEncoding.Decode(dst, buffer[:n])
	if err != nil {
		log.Fatal(err)
	}
	file, _ := os.Create("outputfile.csv")
	file.Write(dst)
	return file.Close()
}