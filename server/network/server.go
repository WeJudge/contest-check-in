package network

import (
    "fmt"
    "net"
)

type Server struct {
    IP          string
    Port        int
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
        Conn:    conn,
    }, nil
}