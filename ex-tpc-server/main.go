package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {

		conn, err := li.Accept()

		if err != nil {
			log.Println(err)
		}

		// io.WriteString(conn, "Connected.......")

		// conn.Close()

		go handle(conn)

	}

}

func handle(conn net.Conn) {

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {

		input := scanner.Text()

		fmt.Println("" + input)

		fmt.Fprintf(conn, "Sent: %s\n", input)

	}

	defer conn.Close()

}
