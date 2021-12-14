package main

import (
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8001")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//buf := make([]byte, 256)
	for {
		io.Copy(os.Stdout, conn)
	}
}
