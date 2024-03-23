package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	var (
		timeout int
		host    string
		port    string
	)

	flag.IntVar(&timeout, "t", 10, "timeout before connecting to the host")
	flag.StringVar(&host, "h", "", "hostname")
	flag.StringVar(&port, "p", "", "port")
	flag.Parse()

	if timeout < 0 {
		fmt.Println("invalid timeout")
		return
	} else if host == "" || port == "" {
		fmt.Println("invalid host or port")
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	<-time.After(time.Duration(timeout) * time.Second)

	// connection
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		log.Fatal(err)
	}

	go readFromSocket(ctx, conn)
	go writeToSocket(ctx, conn)

	<-ctx.Done()
	if err := conn.Close(); err != nil {
		log.Fatal(err)
	}
}

func readFromSocket(ctx context.Context, conn net.Conn) {
	buffer := make([]byte, 1024)

	for {
		select {
		case <-ctx.Done():
			return

		default:
			_, err := conn.Read(buffer)

			if err != nil && !errors.Is(err, io.EOF) && !errors.Is(err, net.ErrClosed) {
				log.Fatal(err)

			} else if err == io.EOF {
				fmt.Println("server closed connection")
				return
			}

			fmt.Println(string(buffer))
		}
	}
}

func writeToSocket(ctx context.Context, conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		select {
		case <-ctx.Done():
			return

		default:

			readFromStdin := scanner.Scan()
			if !readFromStdin {
				os.Exit(0)
			}

			_, err := conn.Write([]byte(scanner.Text()))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
