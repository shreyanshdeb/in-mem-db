package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var data map[string]string

func init() {
	data = make(map[string]string)
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic("Error")
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic("Error")
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		wrds := strings.Split(strings.TrimSpace(ln), " ")
		fmt.Fprintln(conn, wrds)
		switch wrds[0] {
		case "GET":
			if val, ok := data[wrds[1]]; ok && len(wrds) > 1{
				fmt.Fprintln(conn, val)
			} else {
				fmt.Fprintf(conn, "No key \"%s\" found.\n", wrds[1])
			}
			break
		case "SET":
			if len(wrds) > 2 {
				data[wrds[1]] = wrds[2]
			}
			break
		case "DEL":
			if _, ok := data[wrds[1]]; ok {
				delete(data, wrds[1])
			} else {
				fmt.Fprintf(conn, "No key \"%s\" found.\n", wrds[1])
			}
			break
		default:
			fmt.Fprintln(conn, "Incorrect Command")
		}
	}
	defer conn.Close()
}
