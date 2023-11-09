package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		log.Fatalf("Host and port is empty\n")
	}

	socket := args[0] + ":" + args[1]
	conn, err := net.DialTimeout("tcp", socket, *timeout)
	if err != nil {
		log.Fatalf("Cant connect to %s, error: %v", socket, err.Error())
	}

	defer func() {
		if err = conn.Close(); err != nil {
			fmt.Printf("Cant close connection: %v", err.Error())
		}
		log.Println("Connection is closed")
	}()

	go func() {
		reader := bufio.NewReader(conn)
		for {
			msg, err := reader.ReadString('\n')
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Printf("Error! %v", err.Error())
				continue
			}
			fmt.Printf("Got new message:%s", msg)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		_, err = fmt.Fprintf(conn, msg+"\n")
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
