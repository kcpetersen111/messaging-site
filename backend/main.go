package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	// "os"
	// "bufio"
	"net/http"
	// "net"

)

type message struct {

}

var serverToUser chan message




func handCon(w http.ResponseWriter, r *http.Request){
	upgrader := websocket.Upgrader{}

	conn, err := upgrader.Upgrade(w,r,nil)
	if err!=nil{
		log.Fatal("There was an error in the connection.")
	}

	defer conn.Close()

	userToServer := make(chan message)

	go readFromConn(conn, userToServer)
	go writeToConn(conn)
	
}
func writeToConn(conn *websocket.Conn){

}

func readFromConn(conn *websocket.Conn, userToServer chan message){

	// scanner := bufio.NewScanner(conn)
	var data string

	for {
		// scanner.Scan()
		conn.ReadJSON(&data)	
		fmt.Println(data)	

	}

}

func main()  {
	
	serverToUser = make(chan message)

	http.HandleFunc("/", handCon)
    log.Fatal(http.ListenAndServe("localhost:8080", nil))

	// temp := bufio.NewScanner(os.Stdin)
	// temp.Scan()
}



//have the server range over a channel that users would inject there message into
//will need a channel to send data from server to user and a channel to send data from user to server


// need to be able to send messages
// need to be able to be able to create new chats
// need to be able to add user to the chat
// need to get updates from a the server to update messages
//