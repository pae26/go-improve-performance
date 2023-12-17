package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	listener, e := net.Listen("tcp", "localhost:8888")
	if e != nil {
		panic(e)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, e := listener.Accept()
		if e != nil {
			panic(e)
		}
		go func() {
			fmt.Printf("Accept %v", conn.RemoteAddr())
			request, e := http.ReadRequest(bufio.NewReader(conn))
			if e != nil {
				panic(e)
			}
			dump, e := httputil.DumpRequest(request, true)
			if e != nil {
				panic(e)
			}
			fmt.Println(string(dump))
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       io.NopCloser(strings.NewReader("Hello World\n")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}
