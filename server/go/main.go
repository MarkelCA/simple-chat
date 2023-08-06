package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{
    CheckOrigin: customUpgrader,
}
var allowedHosts = map[string]bool{
    "127.0.0.1": true,
    "localhost": true,
}
// Just for local envs. It should not return always true on production applications.
// https://pkg.go.dev/github.com/gorilla/websocket?utm_source=godoc#hdr-Origin_Considerations
var customUpgrader = func(r *http.Request) bool { 
    return true
}


type Message struct {
    Sender  string `json:"sender"`
    Content string `json:"content"`
}

var messages = []Message{{"markelca", "foo"}, {"johnhoo", "rustacean"}}

func main() {
    http.HandleFunc("/list", listMessages)
    http.HandleFunc("/add", handleWebSocket)
    go handleMessages()
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func listMessages(w http.ResponseWriter, r *http.Request)  {
    w.Header().Add("Access-Control-Allow-Origin", "*")
    messagesBytes, jsonErr := json.Marshal(messages)
    if jsonErr != nil {
        log.Println("JSON marshal failed:", jsonErr)
        fmt.Fprintf(w, "Error listing the messages")
        return
    }
    fmt.Fprint(w, string(messagesBytes))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)

    if err != nil {
        log.Println("Error upgrading connection:", err)
        return
    }
    defer conn.Close()

    clients[conn] = true

    for {
        var msg Message
        err := conn.ReadJSON(&msg)
        log.Printf("Message received: %v", msg)
        if err != nil {
            log.Println("Error reading message:", err)
            delete(clients, conn)
            break
        }
        broadcast <- msg
    }
}

func handleMessages() {
    for {
        msg := <-broadcast
        messages = append(messages, msg)
        for conn := range clients {
            err := conn.WriteJSON(msg)
            if err != nil {
                log.Println("Error writing message:", err)
                conn.Close()
                delete(clients, conn)
            }
        }
    }
}


