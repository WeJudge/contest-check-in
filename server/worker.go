package server

import (
    "fmt"
    "sync"
)

var Connections sync.Map

// 运行服务器
// AccessKey WeJudge比赛服务的数据Key，用于校验token
func RunServerWorker (IP string, Port int, AccessKey string) {
    server, err := NewServer(IP, Port)
    if err != nil {
        panic(err.Error())
    }
    fmt.Printf("Server now is listening at %s\n", fmt.Sprintf("%s:%d", IP, Port))
    for {
        conn, err := server.Conn.Accept()
        if err != nil {
            fmt.Printf("listener accept error: %s\n", err.Error())
            continue
        }
        go handleConnection(conn, AccessKey)
    }
}
