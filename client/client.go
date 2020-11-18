package client

import (
    "contest-check-in/protocol"
    "fmt"
    "net"
    "time"
)

type CheckInClient struct {
    Address     string
    Conn        net.Conn
}



func NewClient (IP string, Port int) *CheckInClient {
    client := CheckInClient {
        Address:  fmt.Sprintf("%s:%d", IP, Port),
    }
    return &client
}

func (client *CheckInClient) Connect() error {
    var err error
    client.Conn, err = net.DialTimeout("tcp", client.Address, 3 * time.Second)
    if err != nil {
        return err
    }
    return nil
}

func (client *CheckInClient) Handle(message chan string) {
    for {
        data := make([]byte, 255)
        msg, err := client.Conn.Read(data)
        if msg == 0 || err != nil {
            fmt.Println("disconnect")
            message <- "close"
            break
        }
        fmt.Println(string(data[0:msg]))
    }
}

func (client *CheckInClient) Send(msg protocol.MessageProtocol) error {
    rel := msg.Marshal()
    _, err := client.Conn.Write(rel)
    return err
}
