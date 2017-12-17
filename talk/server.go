package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

//Max 100 client
var mapConn = make(map[string]net.Conn, 100)

func herror(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:3344")
	listener, err := net.ListenTCP("tcp", addr)
	herror(err)
	defer listener.Close()
	for {
		var conn net.Conn
		conn, err = listener.AcceptTCP()
		herror(err)

		mapConn[conn.RemoteAddr().String()] = conn
		go handleMany(conn)

	}

}

func handleMany(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		herror(err)

		if msg != "" {
			fmt.Println(msg)
			msg = conn.RemoteAddr().String() + "    said:  " + msg
			send2Group(msg)
		}

	}

}

func IsClosed(conn net.Conn) bool {
	_, err := conn.Read([]byte{})
	return err == io.EOF
}

func send2Group(msg string) {
	for _, conn := range mapConn {
		if !IsClosed(conn) {
			conn.Write([]byte(msg))
		}
	}
}
