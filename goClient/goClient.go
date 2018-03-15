package main

import (
	"bufio"
	"fmt"
	"net"
)

var quitSemaphore chan bool

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "localhost:7788")

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected!")

	go onMessageReceived(conn)

	// b := []byte("time\n")
	// conn.Write(b)

	// <-quitSemaphore

	for {
		var msg string
		fmt.Scanln(&msg)
		if msg == "quit" {
			break
		}

		b := []byte(msg + "\n")
		conn.Write(b)
	}
}

func onMessageReceived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			quitSemaphore <- true
			break
		}
		// time.Sleep(time.Second)
		// b := []byte(msg)
		// conn.Write(b)
	}
}
