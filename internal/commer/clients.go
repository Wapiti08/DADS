package commer

import (
	"bufio"
	"io"
)

type client struct {
	*bufio.Reader
	*bufio.Writer
	// write channel
	wc chan string
}

func StartClient(msgCh chan<- string, cn io.ReadWriteCloser, quit chan struct{}) (chan<- string, <-chan struct{}) {
	// create a new client with pointer type
	c := new(client)
	// create reader channel by loading reader from connection channel
	c.Reader = bufio.NewReader(cn)
	c.Writer = bufio.NewWriter(cn)

	c.wc = make(chan string)
	done := make(chan struct{})

	go func() {
		scanner := bufio.NewScanner(c.Reader)
		for scanner.Scan() {
			logger.Println()
			msgCh <- scanner.Text()
		}
		done <- struct{}{}
	}()

	// build the writer
	c.writerMonitor()

	// build the part to process stop/done/quit
	go func() {
		select {
		case <-quit:
			cn.Close()
		// exit the select when done
		case <-done:
		}
	}()
	return c.wc, done
}

func (c *client) writerMonitor() {
	go func() {
		for s := range c.wc {
			c.WriteString(s + "\n")
			c.Flush()
		}
	}()
}
