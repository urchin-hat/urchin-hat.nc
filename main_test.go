package main

import (
	"errors"
	"io"
	"net"
	"testing"
	"time"
)

func TestServe(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { listener.Close() })
	done := make(chan error, 1)
	go func() { done <- serve(listener) }()

	// Each connection receives the greeting and EOF; the server stays available.
	for i := 0; i < 3; i++ {
		conn, err := net.DialTimeout("tcp", listener.Addr().String(), time.Second)
		if err != nil {
			t.Fatal(err)
		}
		if err := conn.SetReadDeadline(time.Now().Add(3 * time.Second)); err != nil {
			conn.Close()
			t.Fatal(err)
		}
		got, err := io.ReadAll(conn)
		conn.Close()
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != "hello,world" {
			t.Fatalf("greeting = %q, want %q", got, "hello,world")
		}
	}

	listener.Close()
	select {
	case err := <-done:
		if !errors.Is(err, net.ErrClosed) {
			t.Fatalf("serve error = %v, want net.ErrClosed", err)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("serve did not stop after listener closed")
	}
}
