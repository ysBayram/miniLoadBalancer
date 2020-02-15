package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

var (
	counter    int
	listenAddr = "localhost:3441"
	servers    = []string{
		"localhost:5001",
		"localhost:5002",
		"localhost:5003",
	}
)

func main() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
		}

		backend := chooseBackend()
		fmt.Printf("counter = %v backend = %v\n", counter, backend)
		go func() {
			err := proxy(backend, conn)
			if err != nil {
				log.Printf("Warning: proxying failed: %v", err)
			}
		}()

	}
}

func proxy(backend string, con net.Conn) error {
	bc, err := net.Dial("tcp", backend)
	if err != nil {
		return fmt.Errorf("failed to connect to backend %v : %v", backend, err)
	}

	go io.Copy(bc, con)

	go io.Copy(con, bc)

	return nil
}

func chooseBackend() string {
	//Round robin
	s := servers[counter%len(servers)]
	counter++
	return s
}
