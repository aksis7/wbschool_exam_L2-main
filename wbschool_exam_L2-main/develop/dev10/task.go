package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

// Константа по умолчанию для таймаута
const defaultTimeout = 10 * time.Second

func main() {
	// Парсинг аргументов командной строки
	timeoutArg := flag.Duration("timeout", defaultTimeout, "Timeout for connection (e.g., 10s)")
	flag.Parse()

	// Проверка аргументов
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: go-telnet [--timeout=10s] host port")
		os.Exit(1)
	}

	host := args[0]
	port := args[1]
	address := net.JoinHostPort(host, port)

	// Установка TCP соединения
	conn, err := net.DialTimeout("tcp", address, *timeoutArg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to", address)

	// Канал для завершения программы
	done := make(chan struct{})

	// Чтение данных из сокета и вывод в STDOUT
	go func() {
		io.Copy(os.Stdout, conn)
		fmt.Println("\nConnection closed by remote host")
		done <- struct{}{}
	}()

	// Чтение данных из STDIN и отправка в сокет
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			_, err := conn.Write([]byte(scanner.Text() + "\n"))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error writing to connection: %v\n", err)
				break
			}
		}
		// Завершаем при EOF (Ctrl+D)
		done <- struct{}{}
	}()

	// Ожидаем завершения любой из горутин
	<-done
	fmt.Println("Connection closed")
}
