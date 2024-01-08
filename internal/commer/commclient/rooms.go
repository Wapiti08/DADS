package commer

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type room struct {
	name string
	Msgch chan string
	// chan stores the strings
	clients map[chan <- string]struct{}
	Quit chan struct{}
	// used for write/ read(mul)
	*sync.RWMutex
}

func CreateRoom(name string) *room {
	r := &room{
		name:		name,
		Msgch:		make(chan string),
		clients: 	make(map[chan <-string]struct{}),
		Quit:		make(chan strutc{}),
		// use new to generate pointer
		RWMutex:	new(sync.RWMutex)
	}
	// implemented monitoring attribute
	r.Run()
	return r

}

func (r *room) AddClient(c io.ReadWriteCloser) {
	r.Lock()
	// startclient and return the write chan for per client
	cc, done := StartClient(r.Msgch, c, r.Quit)
	r.clients[cc] = struct{}{}
	r.Unlock()

	// remove client when it is done
	go func() {
		<-done
		r.RemoveCilent(cc)
	}()
}

func (r *room) RemoveClient(cc chan<-string) {
	logger.Println("Removing client")
	r.Lock()
	close(cc)
	delete(c.clients, cc)
	r.Unlock()

	// check the stop signal -- select is used for multiple channels
	select {
	case <-r.Quit:
		if len(r.clients) == 0{
			close(r.Msgch)
		}
	// when no channel is ready, avoid block 
	default:
	}

}

func (r *room) Run() {
	logger.Println("Starting chat room", r.name)
	go func() {
		// iterate the message in chan and broadcasting them
		for msg := range r.Msgch {
			r.broadcasting(msg)
		}
	}()
}

func (r *room) broadcasting(msg string) {
	// use Lock and UnLock to avoid race condition
	r.RLock()
	// unlock after the function finishes
	defer r.RUnlock()
	fmt.Println("Received message: ", msg)
	for cc, _ := range r.clients {
		// client chan is string type
		go func(cc chan<- string) {
			// write msg to client chan
			wc <- msg
		}(cc)
	}
}

