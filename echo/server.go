package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	tcpaddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:4455")
	tcplistener, err := net.ListenTCP("tcp", tcpaddr)
	herror(err)

	defer tcplistener.Close()

	for {
		var tcpconn net.Conn
		tcpconn, err = tcplistener.AcceptTCP()
		herror(err)

		go connHandler(tcpconn)
	}
}

func herror(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func connHandler(conn net.Conn) {
	defer conn.Close()
	fmt.Println(conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Println(msg)

		nmsg := time.Now().String() + "\n"
		conn.Write([]byte(nmsg))
	}
}
