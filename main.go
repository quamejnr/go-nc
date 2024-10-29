package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("provide host address")
		return
	}

	conn, err := net.Dial("tcp", os.Args[1])
	if conn == nil && err != nil {
    fmt.Println("error connecting to host:", err)
		return
	}
	defer conn.Close()
	fmt.Printf("connected to %s...\n", os.Args[1])
	reader := bufio.NewReader(os.Stdin)
	var request []byte

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("error reading input:", err)
			return
		}

		request = append(request, line...)

		if bytes.Compare(line, []byte("\n")) == 0 {
			break
		}
	}

	_, err = conn.Write(request)
	if err != nil {
		fmt.Println("error sending request:", err)
		return
	}
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("error reading request:", err)
	}
	fmt.Println(string(buf))
}
