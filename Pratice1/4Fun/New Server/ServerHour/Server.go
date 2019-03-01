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
const SERVER_IP_L string = "127.0.0.1"

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
	timeNow := time.Now()
	conn.Write([]byte(timeNow.String()))
  	// run loop forever (or until ctrl-c)
  	//for {
		// will listen for message to process ending in newline (\n)
		//message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		//fmt.Print("Message Received:", string(message))
		// sample process for string received
		//newmessage := strings.ToUpper(message)
		// send new string back to client
		
	//}
}