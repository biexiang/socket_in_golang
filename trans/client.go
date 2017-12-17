package main

import (
	//不推荐相对路径
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

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8899")
	conn, err := net.DialTCP("tcp", nil, addr)
	herror(err)
	defer conn.Close()

	go receiveData(conn)

	var input string
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err = reader.ReadString('\n')
		herror(err)
		//数据包头加上数据长度
		complete, err := util.Encode(input)
		herror(err)

		conn.Write(complete)

	}

}

func receiveData(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := util.Decode(reader)
		herror(err)
		fmt.Println(msg)
	}
}
