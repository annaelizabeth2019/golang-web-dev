package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("l err:", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Conn err:", err)
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
			if ln == "" {
				fmt.Println("It's the end of the headers as we know it")
				break
			}
		}

		fmt.Println("code got here")
		io.WriteString(conn, "I see you connected")
		conn.Close()
	}
}
