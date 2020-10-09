package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("fuck you")
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net error", err)
		return
	}
	go broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stdout, "you got something wrong %v", err)
			continue
		}
		go handleConn(conn)
	}
}

type ChannelClient chan<- string //send only

var (
	EnterChannelEnvent   = make(chan ChannelClient)
	LeaveChannelEnvent   = make(chan ChannelClient)
	MessageChannelEnvent = make(chan string)
)

func broadcast() {
	ChannelClientS := make(map[ChannelClient]bool)
	for {
		select {
		case msg := <-MessageChannelEnvent:
			for chanclient := range ChannelClientS {
				chanclient <- msg
			}
		case chanclient := <-EnterChannelEnvent:
			ChannelClientS[chanclient] = true
		case chanclient := <-LeaveChannelEnvent:
			delete(ChannelClientS, chanclient)
			close(chanclient)
		}

	}
}
func handleConn(Conn net.Conn) {
	ch := make(chan string)
	go clientWrite(Conn, ch)
	who := Conn.RemoteAddr().String()
	ch <- "you are " + who
	MessageChannelEnvent <- who + " has joined us"
	EnterChannelEnvent <- ch
	input := bufio.NewScanner(Conn)
	for input.Scan() {
		MessageChannelEnvent <- who + ":" + input.Text()
	}
	LeaveChannelEnvent <- ch
	MessageChannelEnvent <- who + " has left"
	Conn.Close()
}

func clientWrite(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
