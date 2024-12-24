package main

import (
	"bufio"
	"bytes"
	"net"
	"os"
	"testing"
	"time"
)

// Тест 1: Успешное подключение к TCP-серверу
func TestSuccessfulConnection(t *testing.T) {
	// Запускаем временный TCP-сервер
	ln, err := net.Listen("tcp", ":0") // Используем свободный порт
	if err != nil {
		t.Fatalf("Failed to start test server: %v", err)
	}
	defer ln.Close()

	serverAddress := ln.Addr().String()

	// Горутин для обработки входящих соединений
	go func() {
		conn, err := ln.Accept()
		if err != nil {
			t.Errorf("Failed to accept connection: %v", err)
			return
		}
		defer conn.Close()

		// Читаем данные от клиента
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			t.Logf("Server received: %s", message)
			conn.Write([]byte("Echo: " + message))
		}
	}()

	// Подключаемся к серверу как клиент
	client, err := net.DialTimeout("tcp", serverAddress, 2*time.Second)
	if err != nil {
		t.Fatalf("Failed to connect to test server: %v", err)
	}
	defer client.Close()

	// Отправляем данные на сервер
	message := "Hello, Server\n"
	_, err = client.Write([]byte(message))
	if err != nil {
		t.Fatalf("Failed to send data: %v", err)
	}

	// Читаем ответ от сервера
	response := make([]byte, 1024)
	n, err := client.Read(response)
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	expectedResponse := "Echo: " + message
	if string(response[:n]) != expectedResponse {
		t.Errorf("Expected %q, but got %q", expectedResponse, string(response[:n]))
	}
}

// Тест 2: Таймаут подключения
func TestConnectionTimeout(t *testing.T) {
	// Подключаемся к несуществующему адресу
	_, err := net.DialTimeout("tcp", "192.0.2.0:9999", 1*time.Second) // Специальный тестовый IP
	if err == nil {
		t.Fatalf("Expected timeout error, but connection succeeded")
	}

	if !os.IsTimeout(err) && !bytes.Contains([]byte(err.Error()), []byte("timeout")) {
		t.Errorf("Expected timeout error, but got: %v", err)
	}
}

// Тест 3: Отключение клиента по Ctrl+D (EOF)
func TestClientEOF(t *testing.T) {
	// Запускаем временный TCP-сервер
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Failed to start test server: %v", err)
	}
	defer ln.Close()

	serverAddress := ln.Addr().String()

	// Горутин для обработки входящих соединений
	go func() {
		conn, err := ln.Accept()
		if err != nil {
			t.Errorf("Failed to accept connection: %v", err)
			return
		}
		defer conn.Close()

		// Ждём закрытия соединения
		_, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			t.Log("Connection closed by client")
		}
	}()

	// Подключаемся к серверу как клиент
	client, err := net.DialTimeout("tcp", serverAddress, 2*time.Second)
	if err != nil {
		t.Fatalf("Failed to connect to test server: %v", err)
	}
	defer client.Close()

	// Закрываем соединение (имитация Ctrl+D)
	err = client.(*net.TCPConn).CloseWrite()
	if err != nil {
		t.Fatalf("Failed to close client write connection: %v", err)
	}

	time.Sleep(1 * time.Second) // Ждём немного, чтобы сервер отреагировал
	t.Log("Client EOF successfully sent")
}

// Тест 4: Неправильный порт
func TestInvalidPort(t *testing.T) {
	_, err := net.DialTimeout("tcp", "localhost:99999", 1*time.Second)
	if err == nil {
		t.Fatalf("Expected error due to invalid port, but connection succeeded")
	}
}

// Тест 5: Сервер закрывает соединение
func TestServerClosesConnection(t *testing.T) {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Failed to start test server: %v", err)
	}
	defer ln.Close()

	serverAddress := ln.Addr().String()

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			t.Errorf("Failed to accept connection: %v", err)
			return
		}
		conn.Close() // Немедленно закрываем соединение
	}()

	client, err := net.DialTimeout("tcp", serverAddress, 2*time.Second)
	if err != nil {
		t.Fatalf("Failed to connect to test server: %v", err)
	}

	// Пытаемся читать из закрытого соединения
	buffer := make([]byte, 1024)
	_, err = client.Read(buffer)
	if err == nil {
		t.Errorf("Expected error when reading from closed server connection")
	}
	t.Log("Successfully detected server closed connection")
}
