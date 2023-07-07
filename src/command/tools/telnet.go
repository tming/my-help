package tools

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Telnet(host string, port string) error {
	fmt.Printf("ready telnet to %s:%s\n", host, port)
	// create tcp connection
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Dial error:", err)
		return err
	}
	defer conn.Close()
	fmt.Printf("Connected to %s:%s.\n", host, port)

	scanner := bufio.NewScanner(conn)
	writer := bufio.NewWriter(conn)

	// read from stdin，and write to tcp connection
	go func() {
		inputscanner := bufio.NewScanner(os.Stdin)
		for {
			inputscanner.Scan()
			text := inputscanner.Text()
			// fmt.Fprintf(os.Stderr, "got:%s\n", text)
			if text == "quit" {
				os.Exit(0)
			}
			writer.WriteString(text + "\r\n")
			writer.Flush()
		}
	}()

	// read from tcp connection
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// check scanner
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
		return err
	}

	return nil
}
