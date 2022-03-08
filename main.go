package main

import (
	"fmt"
	"github/hideaki10/tcp-server-demo1/packet"
)

func handlePacket(framePayload []byte) (ackFramePayload []byte, err error) {
	var p packet.Packet
	p, err = packet.Decode(framePayload)
	if err != nil {
		fmt.Println("handleConn : packet decode error: ", err)
		return
	}
}

func main() {

}
