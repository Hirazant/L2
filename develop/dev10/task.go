package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func connect(ip string, port string) {
	// Подключаемся к сокету
	address := ip + ":" + port
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("connection failed")
	}
	fmt.Println("Connection start")
	for {

		// Чтение входных данных от stdin
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		// Отправляем в socket
		fmt.Fprintf(conn, text+"\n")
		// Прослушиваем ответ
		message, _ := bufio.NewReader(conn).ReadString('\n')

		if message == "" {
			fmt.Println("connect close")
			break
		}

		fmt.Print("Message from server: " + message)

	}
}

func main() {

	timeout := flag.Int("t", 10, "timeout")

	flag.Parse()

	if flag.NArg() != 2 {
		log.Fatalln("Invalid input(port or ip)")
	}

	args := flag.Args()

	ip := args[0]
	port := args[1]

	fmt.Println("Wait to connection")
	time.Sleep(time.Second * time.Duration(*timeout))

	connect(ip, port)
}
