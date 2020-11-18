package client

import (
    "contest-check-in/protocol"
    "fmt"
    "net"
)

type CheckInClient struct {
    Address     string
    Conn        net.Conn
}


func NewClient (IP string, Port int) (*CheckInClient, error) {
    var err error
    client := CheckInClient {
        Address:  fmt.Sprintf("%s:%d", IP, Port),
    }
    client.Conn, err = net.Dial("tcp", client.Address)
    if err != nil {
        return nil, err
    }
    return &client, nil
}

func (client *CheckInClient) Handle() {
    for {
        data := make([]byte, 255)
        msg, err := client.Conn.Read(data)
        if msg == 0 || err != nil {
            fmt.Println("disconnect")
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
