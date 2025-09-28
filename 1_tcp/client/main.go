package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Подключено к серверу. Введите сообщение (нажмите Enter для отправки):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		conn.Write([]byte(text + "\n"))

		response := make([]byte, 1024)
		n, _ := conn.Read(response)
		fmt.Printf("Сервер: %s", string(response[:n]))
	}
}
