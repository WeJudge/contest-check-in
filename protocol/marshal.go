package protocol

import "encoding/binary"

func (proto *MessageProtocol) Marshal () []byte {
    proto.PayloadLength = uint32(len(proto.Payload))
    rel := make([]byte, 0)
    buf16 := make([]byte, 2)
    buf32 := make([]byte, 4)
    binary.BigEndian.PutUint16(buf16, MagicCode)
    rel = append(rel, buf16...)
    binary.BigEndian.PutUint16(buf16, proto.Action)
    rel = append(rel, buf16...)
    binary.BigEndian.PutUint32(buf32, proto.PayloadLength)
    rel = append(rel, buf32...)
    rel = append(rel, proto.Payload...)
    return rel
}