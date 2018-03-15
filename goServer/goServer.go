package main

import (
	"bufio"
	"fmt"
	"net"
)

var ConnMap map[string]*net.TCPConn

func main() {
	var tcpAddr *net.TCPAddr
	ConnMap = make(map[string]*net.TCPConn)
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "localhost:7788")

	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}

		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		ConnMap[tcpConn.RemoteAddr().String()] = tcpConn
		go tcpPipe(tcpConn)
	}
}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected : " + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		fmt.Println(conn.RemoteAddr().String() + ":" + string(message))
		// msg := time.Now().String() + "\n"
		// b := []byte(msg)
		// conn.Write(b)
		broadcastMessage(conn.RemoteAddr().String() + ":" + string(message))
	}
}

func broadcastMessage(message string) {
	b := []byte(message)

	for _, conn := range ConnMap {
		conn.Write(b)
	}
}
