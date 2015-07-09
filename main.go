package main

import (
	"fmt"
	"net"
	"time"
)

var _ = fmt.Print // Debugging; remove later

type Clock struct{}

func (c Clock) Now() time.Time {
	return time.Now()
}

/*
 * A service is a long running process on the server which gets notified
 * whenever a client is added to the server.
 */
type Server struct {
	port     int32
	services []Service
}

// Messages are serialized as follows
// 123 (start byte)
// ServiceType
// The rest of the message. This is left up to the service to decide the format

const (
	SERVER_HI  = 189
	CLIENT_HI  = 198
	MSG_HEADER = 123
)

type Message interface {
	GetServiceType() ServiceType
	GetContents() []byte
}

func InitServer() *Server {
	return &Server{}
}

func (this *Server) Handle(c *net.UDPConn) {
	client := newClient(MakeICP(c))

	for _, service := range this.services {
		client.AddService(service)
	}

	go client.Start()
}

func (this *Server) AddService(s Service) {
	s.Init()
	this.services = append(this.services, s)
}

func main() {
	server := InitServer()
	physxService := &PhysxService{nil}

	server.AddService(physxService)

	addr := net.UDPAddr{
		Port: 10038,
	}

	for {
		conn, err := net.ListenUDP("udp", &addr)
		fmt.Println(conn)

		if err != nil {
			panic(err)
		}

		go func(c *net.UDPConn) {
			server.Handle(c)
		}(conn)
	}
}
