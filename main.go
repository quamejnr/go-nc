package main

import (
	"bufio"
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
	if err != nil {
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

				if string(line) == "\n" || string(line) == "\r\n" {
					break
				}
			}

			_, err = conn.Write(request)
			if err != nil {
				fmt.Println("error sending request:", err)
				return
			}

			// read response from connection
			buf := make([]byte, 4096)
			_, err = conn.Read(buf)
			if err != nil {
				fmt.Println("error reading response:", err)
			}
			fmt.Printf("%s\n\n", string(buf))
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	fmt.Println("exiting...")

}
