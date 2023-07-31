package main

import (
	"encoding/json"
	"fmt"
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
    http.HandleFunc("/list", listMessages)
    http.HandleFunc("/add", addMessage)
    http.ListenAndServe(":8080", nil)
}

func listMessages(w http.ResponseWriter, r *http.Request)  {
    messagesBytes, jsonErr := json.Marshal(messages)
    if jsonErr != nil {
        log.Println("JSON marshal failed:", jsonErr)
        return
    }
    fmt.Fprint(w, string(messagesBytes))
}


func addMessage(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade failed: ", err)
        return
    }

    for {
        _, input, err := conn.ReadMessage()
        if err != nil {
            log.Println("read failed:", err)
            break
        }
        log.Printf("Received message: %s", string(input))

        var data Message
        jsonErr := json.Unmarshal(input, &data)
        if jsonErr != nil {
            log.Println("Error: {}", jsonErr)
            return
        }
        messages = append(messages, data)
        log.Println(messages)
    }

}
