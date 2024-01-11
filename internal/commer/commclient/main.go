package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
)

func main() {
	rand.Seed(123)

	// initialize random name for connected clients
	name := fmt.Sprintf("Anonymous%d", rand.Intn(400))
	fmt.Println("Starting communicating with multiple clients ...")
	fmt.Println("The connected client is:")
	fmt.Scanln(&name)

	fmt.Printf("Hello %s, connecting to the communication chat system .... \n", name)
	// listen on service
	conn, err := net.Dial("tcp", "127.0.0.1:2300")
	if err != nil {
		log.Fatal("Could not connect to the server", err)
	}
	fmt.Println("Connected to the communication server")
	name += ":"
	// this close executes afer the whole function finishes
	defer conn.Close()
	// define scanner on conn to avoid block
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	// build another scanner for stdin input data sending to conn
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		n, _ := fmt.Fprintf(conn, name+msg+"\n")
		fmt.Println(n)
	}

}
