package client

import (
    "contest-check-in/protocol"
    "fmt"
    "log"
)

func NewClientWorker (IP string, Port int) {
    client, err := NewClient(IP, Port)
    if err != nil {
        panic(err.Error())
    }
    fmt.Printf("%s connected\n", client.Address)
    go client.Handle()

    var msg string
    for {
        msg = ""
        fmt.Scan(&msg)
        switch msg {
        case "quit":
            client.Conn.Close()
            return
        case "hello":
            msg := protocol.MessageProtocol{
                Action: protocol.ActionShakeHand,
                Payload: []byte {
                    1, 2 ,3 , 4, 5,
                },
            }
            err := client.Send(msg)
            if err != nil {
                log.Printf("send message error: %s", err.Error())
            }
        }
    }
}