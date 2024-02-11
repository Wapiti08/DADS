package main

import (
	"bufio"
	"flag"
	"fmt"
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

	// scan the standard input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What messages would you like to send?")
	for scanner.Scan() {
		
		fmt.Println("Writing ", scanner.Text())
		conn.Write(scanner.Bytes())
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		
		if err!=nil{
			log.Fatal(err)
		}

		fmt.Println(string(buffer))
		// give notice for new message
		fmt.Println("What message would you like to send")
	}
	return nil
}

func runUDPServer(address string) error {

}