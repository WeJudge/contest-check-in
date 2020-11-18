package protocol

import (
    "encoding/binary"
    "fmt"
)

func Unmarshal (data []byte)  (*MessageProtocol,error) {
    magicIndex := findMagic(data)
    // 如果捕获到了协议包头
    if magicIndex > -1 {
        proto := MessageProtocol{}
        proto.Action = binary.BigEndian.Uint16(data[magicIndex + 2: magicIndex + 4])
        proto.PayloadLength = binary.BigEndian.Uint32(data[magicIndex + 4: magicIndex + 8])
        proto.Payload = data[magicIndex + 8:]
        return &proto, nil
    }
    return nil, fmt.Errorf("wrong protocol")
}

func BytesSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
    // 表示已经扫描到结尾了
    if atEOF && len(data) == 0 {
        return 0, nil, nil
    }
    if len(data) >= 8 {  // 收到的数据大于等于8字节
        magicIndex := findMagic(data)
        // 如果捕获到了协议包头
        if magicIndex > -1 {
            // 获取packageLength
            pLen := findPackageLength(data, magicIndex)
            // 下次前进magicIndex+pLen个字节
            return magicIndex + pLen, data[magicIndex:pLen], nil
        }
    }
    // 如果已经到了末尾
    if atEOF {
        return len(data), data, nil
    }
    // 表示现在不能分割
    return 0, nil, nil
}


func findMagic (data []byte) int {
    if len(data) < 2 { return 0 }
    length := len(data) - 1
    lc := byte(MagicCode & 0xFF)
    hc := byte(MagicCode >> 8)
    for i := 0; i < length; i++ {
        if data[i] == hc {
            if data[i + 1] == lc {
                return i
            }
        }
    }
    return -1
}

func findPackageLength (data []byte, magicIndex int) int {
    if magicIndex + 8 <= len(data) {        // 如果数据合法
        return int(binary.BigEndian.Uint32(data[magicIndex: magicIndex + 4])    )
    }
    return 0
}