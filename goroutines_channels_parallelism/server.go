package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("⏰ Time server is running on localhost:9000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Printf("🟢 Connection opened from %s\n", conn.RemoteAddr())
	defer func() {
		fmt.Printf("🔴 Connection closed from %s\n", conn.RemoteAddr())
		conn.Close()
	}()

	for {
		currentTime := time.Now().Format("15:04:05")
		_, err := conn.Write([]byte("Time from server: " + currentTime + "\n"))
		if err != nil {
			return 
		}

		time.Sleep(2 * time.Second)
	}
}
