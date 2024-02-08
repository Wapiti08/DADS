package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// define the option to choose server or client
	op := flag.String("type", "", "Server (s) or client (c) ?")
	// define the address for hosting service
	address := flag.String("addr", ":8000", "address? host:port")
	flag.Parse()

	switch strings.ToUpper(*op) {
	case "S":
		runServer(*address)
	case "C":
		runClient(*address)
	}
}

func runServer(address string) error {
	net.Listen("tcp", address)
}

func runClient(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	defer conn.Close()
	// read data from standard input
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println("Writing ", scanner.Text())
		conn.Write(append(scanner.Bytes(), '\r'))

		// read input data
		fmt.Println("What message would you like to send?")
		buffer := make([]byte, 1024)
		// read data from conn to buffer
		_, err := conn.Read(buffer)
		// check err conditions
		if err != nil && err != io.EOF {
			log.Fatal(err)
		} else if err == io.EOF {
			fmt.Println("connection is close")
		}

		fmt.Println(string(buffer))
	}

	return scanner.Err()
}