package main

import (
	"bufio"
	"fmt"
	"io"
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

		scanner := bufio.NewScanner(conn)

		for scanner.Scan() {

			input := scanner.Text()

			fmt.Println("" + input)

			fmt.Fprintf(conn, "Sent: %s\n", input)

			defer conn.Close()

			// io.WriteString(conn, "Connected.......")

		}

		defer conn.Close()

		//missed this.
		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected.")

		conn.Close()

	}

}
