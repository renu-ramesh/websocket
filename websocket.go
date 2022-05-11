package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sacOO7/gowebsocket"
)

var (
	Reset = "\033[0m"
	Cyan  = "\033[36m"
)

// Webapi Response Message format
type ServerMessage struct {
}

var WEB_SOCKET_URL = "wss://democmsapi.cqg.com"

func main() {

	var ch = make(chan bool)
	socket := gowebsocket.New(WEB_SOCKET_URL)

	socket.OnConnected = func(socket gowebsocket.Socket) {
		log.Println("Connected to server")
	}

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Println("Recieved connect error ", err)
	}

	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		log.Println("Recieved message " + message)
	}

	socket.OnBinaryMessage = func(data []byte, socket gowebsocket.Socket) {
		reply := ServerMessage{}
		_ = json.Unmarshal(data, reply)
		fmt.Println("Inside on biary : ")
		b, err := json.MarshalIndent(reply, "", "  ")
		fmt.Println(Cyan + string(b) + Reset)

		if err != nil {
			panic(err)
		}
		ch <- true
	}

	socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Recieved ping " + data)

	}

	socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Recieved pong " + data)

	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		log.Println("Disconnected from server ")
	}

	socket.Connect()

	var data []byte //API request data
	socket.SendBinary(data)
	<-ch

	close(ch)

}
