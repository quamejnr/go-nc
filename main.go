package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
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
	fmt.Printf("connected to %s...\tPress Ctrl+C to quit\n", os.Args[1])

	go func() {
		for {
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
			buf := make([]byte, 4096)
			_, err = conn.Read(buf)
			if err != nil {
				fmt.Println("error reading request:", err)
			}
			fmt.Printf("%s\n\n", string(buf))
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	_ = <-sigChan
  fmt.Println("exiting...")

}
