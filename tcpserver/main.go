package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp4"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err == io.EOF {
			fmt.Println("client close connection")
			break
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		result := strconv.Itoa(rand.Intn(100)) + "\n"
		go conn.Write([]byte(result))
	}
}

type handler struct {
}

func newHandler() *handler {
	return &handler{}
}

func (h *handler) handleRequest(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())
	for {
		// Make a buffer to hold incoming data.
		buf := make([]byte, 1024)
		// Read the incoming connection into the buffer.
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			if err == io.EOF {
				fmt.Println("client close connection")
				break
			}
		}
		// Send a response back to person contacting us.
		msg := fmt.Sprintf("received msg: %s\n", buf[:len])
		conn.Write([]byte(msg))
	}
}

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	// close listener
	defer l.Close()
	rand.Seed(time.Now().Unix())

	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	//h := newHandler()
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			return
		}
		// Handle connections in a new goroutine.
		go handleConnection(conn)
	}
}
