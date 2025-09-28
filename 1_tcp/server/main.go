package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Соединение не установлено")
			return
		}
		message = strings.TrimSpace(message)
		fmt.Printf("Получено сообщение: %s\n", message)
		conn.Write([]byte("Эхо: " + message + "\n"))
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("TCP server listening on :8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Получили ошибку:", err)
			continue
		}
		go handleConnection(conn) // каждый клиент — в отдельной горутине
	}
}
