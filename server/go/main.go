package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var messages = []Message{{"markelca", "foo"}, {"johnhoo", "rustacean"}}

type Message struct {
    Sender  string `json:"sender"`
    Content string `json:"content"`
}


func main() {
    http.HandleFunc("/", messageHandler)
    http.ListenAndServe(":8080", nil)
}

func messageHandler(w http.ResponseWriter, r *http.Request)  {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade failed: ", err)
        return
    }
    defer conn.Close()

    for {
        mt, input, err := conn.ReadMessage()
        if err != nil {
            log.Println("read failed:", err)
            break
        }
        log.Printf("Received message: %s", string(input))

        addMessage(input)
        sendResponse(conn, mt)
    }
}

func addMessage(input []byte) {
    var data Message
    err := json.Unmarshal(input, &data)
    if err != nil {
        log.Println("Error: {}", err)
        return
    }
    messages = append(messages, data)
}

func sendResponse(conn *websocket.Conn, mt int) {
    messagesBytes, err := json.Marshal(messages)
    if err != nil {
        log.Println("JSON marshal failed:", err)
        return
    }

    writeErr := conn.WriteMessage(mt, messagesBytes)
    if writeErr != nil {
        log.Println("write failed:", err)
        return
    }
}
