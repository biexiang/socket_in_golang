package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func herror(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

var quitSemaphore chan bool

func main() {
	connaddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:4455")
	conn, err := net.DialTCP("tcp", nil, connaddr)
	herror(err)

	defer conn.Close()

	go goHandleConn(conn)
	msg := []byte("Hello Fucking World \n")
	conn.Write(msg)
	<-quitSemaphore

}

func goHandleConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			herror(err)
			quitSemaphore <- true
			break
		}
		fmt.Println(msg)
		time.Sleep(time.Second)

		conn.Write([]byte(msg))
	}

}
