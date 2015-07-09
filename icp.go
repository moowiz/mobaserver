package main

import (
	"net"
)

type ICPConn interface {
	HasMessages() bool
	GetNextMessage() (Message, error)
	SendMessage(Message) error
	ReadByte() (byte, *net.UDPAddr, error)
	WriteByte(byte, *net.UDPAddr) error
}

type icpConn struct {
	udp     *net.UDPConn
	addr    *net.UDPAddr
	sendSeq byte
	recvSeq byte
}

func seqMoreRecent(a, b, max byte) bool {
	return (a > b) && (a-b <= max/2) || (b > a) && (b-a > max/2)
}

func MakeICP(udp *net.UDPConn) ICPConn {
	return &icpConn{udp, nil, 0, 0}
}

func (this *icpConn) HasMessages() bool {
	return false
}

func (this *icpConn) GetNextMessage() (msg Message, err error) {
	return nil, nil
}

func (this *icpConn) SendMessage(Message) error {
	return nil
}

func (this *icpConn) ReadByte() (byte, *net.UDPAddr, error) {
	buf := make([]byte, 1)
	n, addr, err := this.udp.ReadFromUDP(buf)

	if err != nil {
		return 0, nil, err
	}
	if n != 1 {
		panic("oh gosh n != 1")
	}

	return buf[0], addr, nil
}

func (this *icpConn) WriteByte(b byte, addr *net.UDPAddr) error {
	n, oobn, err := this.udp.WriteMsgUDP([]byte{b}, nil, addr)
	if err != nil {
		return err
	}
	if n != 1 || oobn != 0 {
		panic("sent the wrong number of bytes???")
	}
	return nil
}
