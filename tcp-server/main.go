package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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

		go handle(conn)
	}
}

func handle(conn net.Conn) {

	//timeout after 15seconds
	err := conn.SetDeadline(time.Now().Add(15 * time.Second))
	if err != nil {
		log.Println("Connection Timed Out")
	}

	scanner := bufio.NewScanner(conn)

	//prints out connection
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "Recieved: %s\n", ln)
	}

	defer conn.Close()

	// fmt.Println("Heres code")

}
