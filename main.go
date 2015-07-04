package main

import (
	"fmt"
	"github.com/moowiz/gophysx"
	"net"
	"time"
)

type Clock struct{}

func (c Clock) Now() time.Time {
	return time.Now()
}

type Server struct {
	system *gophysx.System
}

func main() {
	//server := Server{gophysx.Init(Clock{})}

	addr := net.UDPAddr{
		Port: 10037,
	}
	for {
		conn, err := net.ListenUDP("udp", &addr)

		if err != nil {
			panic(err)
		}

		//go newClient(conn, &server)
		icp := InitICP(*conn)

		buf, err := icp.Read()
		fmt.Println(string(buf))
		_, err = icp.Write([]byte("back at you\n"))
		icp.Close()

		fmt.Println("DONE MAN")
	}
}
