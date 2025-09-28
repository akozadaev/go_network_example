package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8081")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("UDP server listening on :8081")

	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Получили ошибку:", err)
			continue
		}
		msg := string(buffer[:n])
		fmt.Printf("From %s: %s", clientAddr, msg)

		// Отправляем ответ
		conn.WriteToUDP([]byte("Привет, привет!"), clientAddr)
	}
}
