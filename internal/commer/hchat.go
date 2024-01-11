package commer

import (
	"distributed_anomaly_detection_system/internal/hlogger"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var logger = hlogger.GetInstance()

func Run(connection string) error {

	// create listener to keep monitor connection
	l, err := net.Listen("tcp", connection)
	if err != nil {
		logger.Println("Error connecting to chat client", err)
		return err
	}
	// create the room for chat
	r := CreateRoom("AnoamalDetection")

	// handle SIGINT and SIGTERM
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		// block until receive some data
		<-ch
		// flush everything if received signal
		l.Close()
		fmt.Println("Closing tcp connection")
		close(r.Quit)
		// check whether there is existed clients
		if r.ClCount() > 0 {
			<-r.Msgch
		}
		// exit the current program
		os.Exit(0)
	}()

	// process the connection
	for {
		// accept connection
		conn, err := l.Accept()
		if err != nil {
			logger.Println("Error accepting connection from client", err)
			break
		}
		// handle connection --- add client
		go handleConnection(r, conn)
	}

	return err
}

func handleConnection(r *room, c net.Conn) {
	// print out the address of request
	logger.Println("Received request from client", c.RemoteAddr())
	r.AddClient(c)
}
