package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	//ch := make(chan int) // make a channel
	// two channels of the same type may be compared using ==,
	// the comparison is true if both are references to the same
	// channel data structure

	// a channel has two principal operations, send and receive. known as communications
	// a send statement transmits a value from one goroutine, through the channel, to another goroutine
	// executing a corresponding receive expression.

	//ch <- x // a send statement
	//x = <- ch // a receive expression in an assignment statement
	//<-ch // a receive statement and the result is discarded

	// channels have a third operation called Close which set's a flag indicating that no more values
	// will ever be sent on this channel

	//close(ch)


	// Unbuffered Channels
	// a channel created with a simple call to make (with one argument to make) is called an unbuffered channel
	// make takes an optional second argument as an int and it is called the channels capacity

	//ch = make(chan int)
	//ch = make(chan int, 0) // only non zero arguments bring a capacity so this is unbuffered
	//ch = make(chan int, 3) // buffered channel with capacity 3

	// Communication over an unbuffered channel causes the sending and receiving goroutines to synchronize
	// because of this, unbuffered channels are sometimes called synchronous channels

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // ignoring errors
		log.Panicln("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}