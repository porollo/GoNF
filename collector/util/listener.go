package util

import (
	"fmt"
	"net"
)

func listener() {
	addr, _ := net.ResolveUDPAddr("udp", ":9999")
	sock, _ := net.ListenUDP("udp", addr)
	sock.SetReadBuffer(1048576)

	i := 0
	for {
		i++
		buf := make([]byte, 1024)
		rlen, _, err := sock.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(buf[0:rlen])
		fmt.Println(i)
	}
}