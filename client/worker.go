package client

import (
    "contest-check-in/protocol"
    "fmt"
    "log"
    "time"
)

var cMsg chan string

func doConnect(client *CheckInClient, status int) error {
    if status == 1 {
        log.Println("connection intercept, retry after 3s.")
    } else if status == 2 {
        log.Println("connection failed, retry after 3s.")
    }
    time.Sleep(3 * time.Second)
    err := client.Connect()
    if err != nil {
        return err
    }
    log.Printf("%s connected\n", client.Address)
    go client.Handle(cMsg)
    return nil
}

func NewClientWorker (IP string, Port int) {

    cMsg = make(chan string)

    client := NewClient(IP, Port)
    err := doConnect(client, 0)
    if err != nil {
        cMsg <- "fail"
        log.Println(err.Error())
    }
    go func() {
        for {
            select {
                case msg := <- cMsg:
                    if msg == "close" {
                        err := doConnect(client, 1)
                        if err != nil {
                            cMsg <- "fail"
                        }
                    } else if msg == "fail" {
                        err := doConnect(client, 2)
                        if err != nil {
                            cMsg <- "fail"
                        }
                    }
            }
        }
    }()



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