package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	// "os"
	// "bufio"
	"net/http"
	// "net"
	"encoding/json"
	"errors"

)

type message struct {

}

var serverToUser chan message




// func handCon(w http.ResponseWriter, r *http.Request){
// 	upgrader := websocket.Upgrader{}

// 	conn, err := upgrader.Upgrade(w,r,nil)
// 	if err!=nil{
// 		log.Fatal("There was an error in the connection.")
// 	}

// 	defer conn.Close()

// 	userToServer := make(chan message)

// 	go readFromConn(conn, userToServer)
// 	// go writeToConn(conn)
	
// }
// // func writeToConn(conn *websocket.Conn){

// // }

// func readFromConn(conn *websocket.Conn, userToServer chan message){

// 	// scanner := bufio.NewScanner(conn)
// 	var data string

// 	//can use a non blocking event loop to taken out put and be able to write to the output at the same time
// 	for {
// 		// scanner.Scan()
// 		conn.ReadJSON(&data)	
// 		fmt.Println(data)	

// 	}

// }

func messageToServer(w http.ResponseWriter, r *http.Request){

	fmt.Println("Endpoint hit")

	// json.Marshal("{test:'test'}")
	// fmt.Println(json.Marshal("{test:'test'}"))

	// var message []byte
	// body,err := r.GetBody()
	// if err!=nil {
	// 	log.Println("Error: ",err)
	// }
	// body.Read(message)
	// fmt.Println(message)

	// var jsonMessage string
	// json.Unmarshal(message,jsonMessage)

	var e string
	var unmarshalErr *json.UnmarshalTypeError
	
	// var temp []byte
	temp := make([]byte, 1024)
	num,err := r.Body.Read(temp)
	// fmt.Println(temp," ",num," ",err)

	err = json.Unmarshal(temp, e)
	// fmt.Println(e)

	// decoder := json.NewDecoder(r.Body)
	// // decoder.DisallowUnknownFields()
	// err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}


	fmt.Println(e)
}

func checkChat(w http.ResponseWriter, r *http.Request){

}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func main()  {
	
	myRouter := mux.NewRouter().StrictSlash(true)

	// myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/messageToServer", messageToServer).Methods("POST")
	// myRouter.HandleFunc("/checkChat", checkChat).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080",myRouter))

	// temp := bufio.NewScanner(os.Stdin)
	// temp.Scan()
}


//probably just use http requests cuz i dont know how to use websockets


//have the server range over a channel that users would inject there message into
//will need a channel to send data from server to user and a channel to send data from user to server


// need to be able to send messages
// need to be able to be able to create new chats
// need to be able to add user to the chat
// need to get updates from a the server to update messages
//