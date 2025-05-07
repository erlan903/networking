// #Client#
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	// dial function call to connect a server
	connection, err := net.Dial("udp", "127.0.0.1:80")

	if err != nil {
		fmt.Println("Connection failed.")
		return
	} else {
		for {
			fmt.Println("Connection succeed.", time.Now(), connection.RemoteAddr())
			t := time.Now()
			time_data := t.String()
			//sending data over udp
			connection.Write([]byte(time_data))
			time.Sleep(3000 * time.Millisecond)
		}

	}
}
