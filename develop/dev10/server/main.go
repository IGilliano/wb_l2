package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Cant start server: %v", err.Error())
	}

	defer func() {
		if err = listener.Close(); err != nil {
			log.Printf("Cant close connection: %v", err.Error())
		}
	}()

	log.Println("Listen and serve on port 8080")

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')

		if err == io.EOF {
			return
		}

		if err != nil {
			log.Printf("Cant read message: %v", err.Error())
		}

		fmt.Printf("Got new message:%s", msg)

		_, err = conn.Write([]byte("->" + msg))
		if err != nil {
			log.Printf("Cant write message socket: %v", err.Error())
		}
	}
}
