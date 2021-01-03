// this file is a server that lets clients chat with each other.
//for the test : go build chat.go
// go build client.go
//./chat &
// in two different terminals, launch ./client and chatting
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client chat messages
)

//The broadcaster listens on the global entering and leaving channels for announcements of arriving and departing clients. 
//When it receives one of these events, it updates the "clients"
func broadcaster() {
	// create a local variable type Map "clients" to record the current set of all connected clients
	clients := make(map[client]bool)
	
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			//TODO
            for clientIterator, _ := range clients {
                if clients[clientIterator] == true {
                    clientIterator <- msg
                }
            }

		case cli := <-entering:
			//TODO
			clients[cli] = true
		
		//if a client leaving ?
		case cli := <-leaving:
			//TODO
			clients[cli] = false
		}
	}
}

//!-broadcaster

//!+handleConn
//The handleConn function creates a new outgoing message channel for its client 
//and announces the arrival of this client to the broadcaster over the "entering" channel
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	//send chat message via os.stdin
	//TODO
	scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        messages <- who + " said : " + scanner.Text()
    }

	//ctrl-D leaves os.stdin and leaves chatting
	//TODO
	messages <- who + " has left"
    leaving <- ch

	//Close the connection
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main goroutine is to listen for and accept incoming network connections from clients. For each one,it creates a new handleConn goroutine.
func main() {
	//create listener
	//TODO
	server, _ := net.Listen("tcp", ":8000")
	//create broadcaster
	go broadcaster()
	//Handle connection
	for {
		//1: listener should accept the incoming network connection
		//TODO
		conn, err := server.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		//2.create connection using handleConn method
		//TODO
		go handleConn(conn)
	}
}
