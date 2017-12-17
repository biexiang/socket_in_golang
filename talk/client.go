package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func herror(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

var semaphore chan bool

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:3344")
	conn, err := net.DialTCP("tcp", nil, addr)
	herror(err)
	defer conn.Close()

	go handleOne(conn)

	var msg string
	for {
		reader := bufio.NewReader(os.Stdin)
		msg, err = reader.ReadString('\n')
		herror(err)

		if msg == "quit" {
			break
		}
		conn.Write([]byte(msg))
	}
}

func handleOne(conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		herror(err)
		if msg != "EOF" {
			fmt.Println(msg)
		}
	}
}
