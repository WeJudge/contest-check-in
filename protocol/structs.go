package protocol

/*
    Message Protocol Definition
    Magic | Action | Payload-Length | Payload (json)
      2   |   2   |        4        | ...
 */
type MessageProtocol struct {
    Action          uint16
    PayloadLength   uint32
    Payload         []byte
}

const (
    MagicCode = 0xC540
    ActionShakeHand = 0x0001
)

