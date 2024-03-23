package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	accept()
}

func accept() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listening on :3000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("accepted incomming connection")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	go func() {
		for i := 0; i > -1; i++ {
			time.Sleep(time.Second * 5)
			if i%2 == 0 {
				_, err := conn.Write([]byte("message from server"))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("message sent")
			} else if i%30 == 0 {
				conn.Close()
				return
			}
		}
	}()

	for {
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		} else if err == io.EOF {
			return
		}

		fmt.Println(string(buffer))
	}
}
