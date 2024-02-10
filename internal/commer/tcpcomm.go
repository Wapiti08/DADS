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
	"time"
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
	
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Fail to listen: ", address)
	}
	log.Println("Listening ...")
	defer l.Close()
	for {
		// the listener containers the built-in multiple channels
		conn, err := l.Accept()
		// check error
		if err != nil {
			return err
		}
		// handle connection
		go handleConnection(conn)
	}

}

func handleConnection(c net.Conn) {
	defer c.Close()

	// define reader and writer, reader receives info from client, writer sends info to client
	reader := bufio.NewReader(c)
	writer := bufio.NewWriter(c)

	// set reader method with line by line
	for {
		c.SetDeadline(time.Now().Add(5 * time.Second))
		// read the separater for rows (pre-defined)
		line, err := reader.ReadString('\r')
	
		if err != nil && err != io.EOF {
			log.Println(err)
			return

		} else if err == io.EOF {
			log.Println("Connection closed")
			return
		}

		fmt.Printf("Received %s from address %s \n", line[:len(line)-1], c.RemoteAddr)
		// flush the message
		writer.WriteString("Message received ...")
		writer.Flush()
	}

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
			log.Println("connection is close")
		}

		fmt.Println(string(buffer))
	}

	return scanner.Err()
}