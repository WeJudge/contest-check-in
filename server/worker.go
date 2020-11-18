package server

import (
    "bufio"
    "contest-check-in/protocol"
    "contest-check-in/server/network"
    "fmt"
    "log"
    "net"
    "sync"
)

var Connections sync.Map

// 运行服务器
// AccessKey WeJudge比赛服务的数据Key，用于校验token
func RunServerWorker (IP string, Port int, AccessKey string) {
    server, err := network.NewServer(IP, Port)
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