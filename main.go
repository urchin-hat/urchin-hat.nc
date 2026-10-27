package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:2323", "TCP listen address")
	flag.Parse()

	listener, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Printf("listening on %s", listener.Addr())

	if err := serve(listener); err != nil {
		log.Fatal(err)
	}
}

func serve(listener net.Listener) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	if err := conn.SetWriteDeadline(time.Now().Add(5 * time.Second)); err != nil {
		log.Printf("set write deadline: %v", err)
		return
	}
	if _, err := io.WriteString(conn, "hello,world\n"); err != nil {
		log.Printf("write greeting: %v", err)
	}
}
