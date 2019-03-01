package main

import (
	"net"
	"os"
	"fmt"
	"strings"
	"time"
)

const SERVER_PORT string = "8000"
const SERVER_IP_L string = "0.0.0.0"

var LOG bool = false

func ErrorHandle(err error){
    if err  != nil {
		fmt.Println("We have a error: " , err)
		fmt.Println("Exitting") 
		os.Exit(0)
    }
}

func PrintErrorIfExists(err error){
    if err != nil {
        fmt.Println("We have a error: ",err)
    } 
}

func main() {
	//Log check
	if len(os.Args) > 1 && (strings.ToLower(os.Args[1]) == "log"){ LOG = true}
	
	//LOG
	if LOG { fmt.Println("Starting server.") }

	//Listener TCP
	listener, errorCheck := net.Listen("tcp", SERVER_IP_L+":"+SERVER_PORT)
	ErrorHandle(errorCheck)
	
	//LOG
	if LOG { fmt.Println("Server running, Listening ALL interfaces.") }
	
	for{
		//Accept connection on port
		conn, errorCheck := listener.Accept()
		PrintErrorIfExists(errorCheck)
		
		//Scheduling of connection close
		defer conn.Close()
		if LOG { defer fmt.Println("Server Closed") }
		//Get time
		timeNow := time.Now()

		//LOG
		if LOG { fmt.Println("Client connected in: " + timeNow.String()) }
		
		//Send to client
		conn.Write([]byte(timeNow.String() + "\n"))
	}
}