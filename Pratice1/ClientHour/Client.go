package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {

	conn, _ := net.Dial("tcp", os.Args[1]+":8000")
	defer conn.Close()

    message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print(message)

}