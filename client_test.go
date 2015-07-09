package main

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"net"
	"testing"
)

type Pipe struct {
	buffer []byte
	id     int
}

var id int

func NewPipe() *Pipe {
	p := Pipe{make([]byte, 0), id}
	id++
	return &p
}

func (this *Pipe) Write(buf []byte) (n int, err error) {
	cpy := make([]byte, len(buf))
	copy(cpy, buf)
	this.buffer = append(this.buffer, cpy...)
	return len(buf), nil
}

func (this *Pipe) Read(b []byte) (n int, err error) {
	fmt.Println("reading", this, "into", b)
	if len(b) > 0 {
		copy(b, this.buffer)
		this.buffer = this.buffer[len(b):]
		return len(b), nil
	}
	return 0, nil
}

type FakeICP struct {
	// client means messages to the client
	clientPr io.Reader
	clientPw io.Writer
	// server means messages to the server
	serverPr io.Reader
	serverPw io.Writer
}

func (this FakeICP) HasMessages() bool {
	return true
}

func (this FakeICP) GetNextMessage() (Message, error) {
	return nil, nil
}

func (this FakeICP) SendMessage(Message) error {
	return nil
}

func (this FakeICP) ReadByte() (byte, *net.UDPAddr, error) {
	buf := make([]byte, 1)
	fmt.Println("read byte", this.serverPr)
	this.serverPr.Read(buf)
	return buf[0], nil, nil
}

func (this FakeICP) WriteByte(b byte, addr *net.UDPAddr) error {
	fmt.Println("writing", b, "to", this.clientPw)
	this.serverPw.Write([]byte{b})
	fmt.Println("Now is", this.clientPw)
	return nil
}

func (this FakeICP) sendToServer(buf []byte) {
	this.clientPw.Write(buf)
}

func (this FakeICP) readFromServer(numExpected int) []byte {
	fmt.Println("before make")
	buf := make([]byte, numExpected)
	fmt.Println("before read")
	this.clientPr.Read(buf)
	fmt.Println("before return")
	return buf
}

type FakeService struct {
}

func makeFakeICP() FakeICP {
	clientToServer, serverToClient := NewPipe(), NewPipe()
	icp := FakeICP{serverToClient, clientToServer, clientToServer, serverToClient}
	return icp
}

func TestClientHandshakeValid(t *testing.T) {
	Convey("Server should handshake correctly", t, func() {
		icp := makeFakeICP()
		fmt.Println("Sending")
		icp.sendToServer([]byte{CLIENT_HI})
		fmt.Println("New client")
		newClient(icp)

		fmt.Println("Sent")
		So(icp.readFromServer(1), ShouldResemble, []byte{SERVER_HI})
	})

}
