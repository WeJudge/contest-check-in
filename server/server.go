package server

import (
    "bufio"
    "contest-check-in/protocol"
    "fmt"
    "log"
    "net"
)

type Server struct {
    IP          string
    Port        int
    Address     string
    Conn        net.Listener
}

type ClientConnection struct {
    Address     string
    Conn        net.Conn
    TeamId      string  // 队伍ID
    TeamName    string  // 队名
    Members     string  // 队伍成员
    UserId      string  // 比赛账号内部Id
}

func NewServer (IP string, Port int) (*Server, error) {
    conn, err := net.Listen("tcp", fmt.Sprintf("%s:%d", IP, Port))
    if err !=nil {
        return nil, err
    }
    return &Server{
        IP:      IP,
        Port:    Port,
        Address: fmt.Sprintf("%s:%d", IP, Port),
        Conn:    conn,
    }, nil
}


func handleConnection (conn net.Conn, AccessKey string) {
    clientKey := conn.RemoteAddr().String()
    log.Printf("client [%s] connected\n", clientKey)

    // TODO  校验Token，AccessKey
    Connections.Store(clientKey, conn)

    input := bufio.NewScanner(conn)
    input.Split(protocol.BytesSplitter)
    for input.Scan() {
        pack, err := protocol.Unmarshal(input.Bytes())
        if err != nil {
            log.Println("receive a wrong package, ignore.")
            continue
        }
        log.Printf("receive a valid package, action: %d\n", pack.Action)
    }
    // 移除并清理
    log.Printf("client [%s] disconnect\n", clientKey)
    Connections.Delete(clientKey)
    _ = conn.Close()
}