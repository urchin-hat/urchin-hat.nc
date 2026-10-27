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
	remote := conn.RemoteAddr().String()
	if err := conn.SetWriteDeadline(time.Now().Add(5 * time.Second)); err != nil {
		log.Printf("access remote=%s bytes=0 status=error error=%q", remote, err)
		return
	}
	n, err := io.WriteString(conn, "hello,world")
	if err != nil {
		log.Printf("access remote=%s bytes=%d status=error error=%q", remote, n, err)
		return
	}
	log.Printf("access remote=%s bytes=%d status=ok", remote, n)
}
