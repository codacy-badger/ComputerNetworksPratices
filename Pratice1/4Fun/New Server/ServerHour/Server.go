package main

import (
	"net"
	"os"
	"fmt"
	"strings"
	"time"
)

//serverPort: set a default port to connect, usually is (8000).
const serverPort string = "8000"
//serverIpL: set a ip to listen all interfaces.
const serverIPL string = "0.0.0.0"
"LOG: flag set via command line and defines whether the program should print logs for the user."
var LOG bool

"ErrorHandle(err error): recives a error var, and check if exist error"
"if exists: print a error mensage and exit the program"
func ErrorHandle(err error){
    if err  != nil {
		fmt.Println("We have a error: " , err)
		fmt.Println("Exitting") 
		os.Exit(0)
    }
}

"PrintErrorIfExists(err error): recives a error var, and check if exist error"
"if exists: print a error mensage and NOT exit the program"
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
	listener, errorCheck := net.Listen("tcp", serverIPL+":"+serverPort)
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