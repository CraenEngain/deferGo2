package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	// Парсинг флагов и аргументов
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	if len(flag.Args()) != 2 {
		fmt.Println("Usage: go-telnet --timeout=10s host port")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	address := net.JoinHostPort(host, port)

	// Установка TCP-соединения с таймаутом
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Connected to %s\n", address)

	// Каналы для сигнализации о завершении работы
	done := make(chan struct{})

	// Чтение из соединения и запись в STDOUT
	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from connection: %v\n", err)
		}
		done <- struct{}{}
	}()

	// Чтение из STDIN и запись в соединение
	go func() {
		if _, err := io.Copy(conn, os.Stdin); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to connection: %v\n", err)
		}
		done <- struct{}{}
	}()

	// Ожидание завершения работы
	<-done
	fmt.Println("Connection closed")
}
