package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Marshaler interface {
    MarshalJSON() ([]byte, error)
}
type JSONTime time.Time

func (t JSONTime)MarshalJSON() ([]byte, error) {
    //do your serializing here
    stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
    return []byte(stamp), nil
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{
    CheckOrigin: customUpgrader,
}
// Just for local envs. It should not return always true on production applications.
// https://pkg.go.dev/github.com/gorilla/websocket?utm_source=godoc#hdr-Origin_Considerations
var customUpgrader = func(r *http.Request) bool { 
    return true
}


type Message struct {
    Sender  string `json:"sender"`
    Content string `json:"content"`
    Time JSONTime `json:"time"`
}

var messages = []Message{
    {
        Sender  : "markelca", 
        Content : "foo", 
        Time    : JSONTime(time.Now()),
    }, 
    {
        Sender  : "johnhoo", 
        Content : "rustacean", 
        Time    : JSONTime(time.Now()),
    },
}

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
        msg.Time = JSONTime(time.Now())
        log.Printf("debug: %v", msg)
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


