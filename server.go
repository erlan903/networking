// #Server#
package main

import (
	"fmt"
	"net"
	"time"
)

// IP is specified as localhost (127.0.0.1)
var ip = net.ParseIP("127.0.0.1")
var udpAddr = net.UDPAddr{
	IP:   ip,
	Port: 80,
}

func main() {
	//listening the given port and IP
	udp_conn, err := net.ListenUDP("udp", &udpAddr)
	if err != nil {
		fmt.Println("Connection failed.", err)
		return
	} else {
		var bufferArray [128]byte
		bufferSlice := bufferArray[:]
		for {
			if err != nil {
				fmt.Println("Error occured: ", err)
			}
			//reading data from udp connection
			n, addr, error_from_udp := udp_conn.ReadFromUDP(bufferSlice)
			fmt.Println("Packet.", "Current time: ", time.Now())
			if error_from_udp != nil {
				fmt.Println("Error from UDP: ", error_from_udp)
			}
			fmt.Println(n, "Source Address: ", addr, "Error Code: ", err)
			fmt.Println("Data: ", bufferSlice)
		}

	}

}
