package main

import (
	"net"
	"os"
	"fmt"
	//"bufio"
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
	if len(os.Args) > 1 && (strings.ToLower(os.Args[1]) == "log"){ 
		fmt.Println("yes")
		LOG = true
		}else{
			fmt.Println("no")
		}
	//LOG
	if LOG { fmt.Println("Starting server.") }

	//Listener TCP
	listener, errorCheck := net.Listen("tcp", SERVER_IP_L+":"+SERVER_PORT)
	ErrorHandle(errorCheck)
	
	//LOG
	if LOG { fmt.Println("Server running, Listening ALL interfaces.") }

  	//Accept connection on port
  	conn, errorCheck := listener.Accept()
	ErrorHandle(errorCheck)
	defer conn.Close()

	//Send time
	timeNow := time.Now()
	conn.Write([]byte(timeNow.String() + "\n"))

	
}