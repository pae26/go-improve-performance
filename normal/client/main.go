package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	conn, e := net.Dial("tcp", "localhost:8888")
	if e != nil {
		panic(e)
	}
	request, e := http.NewRequest("GET", "http://localhost:8888", nil)
	if e != nil {
		panic(e)
	}
	request.Write(conn)
	response, e := http.ReadResponse(bufio.NewReader(conn), request)
	if e != nil {
		panic(e)
	}
	dump, e := httputil.DumpResponse(response, true)
	if e != nil {
		panic(e)
	}
	fmt.Println(string(dump))
}
