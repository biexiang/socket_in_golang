package main

import (
	"./util"
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

//Server 端
func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8899")
	listener, err := net.ListenTCP("tcp", addr)
	herror(err)
	defer listener.Close()

	for {
		var conn net.Conn
		conn, err = listener.AcceptTCP()
		herror(err)
		defer conn.Close()

		go handleMany(conn)

	}

}

func handleMany(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := util.Decode(reader)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(msg)

		//Encode 写
		msg = "Repeat Your Word : " + msg
		newmsg, err := util.Encode(msg)
		herror(err)
		conn.Write(newmsg)
	}

}
