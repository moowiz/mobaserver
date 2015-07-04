package main

import (
	"errors"
	"net"
	"time"
)

type ICPConn struct {
	udp     net.UDPConn
	addr    *net.UDPAddr
	buf     []byte
	sendSeq byte
	recvSeq byte
}

func seqMoreRecent(a, b, max byte) bool {
	return (a > b) && (a-b <= max/2) || (b > a) && (b-a > max/2)
}

func InitICP(udp net.UDPConn) *ICPConn {
	return &ICPConn{udp, nil, make([]byte, 20), 0, 0}
}

func (this *ICPConn) Read() ([]byte, error) {
	buf := make([]byte, 128)
	_, addr, err := this.udp.ReadFromUDP(buf)
	this.addr = addr
	return buf, err
}

func (this *ICPConn) Write(b []byte) (int, error) {
	if this.addr == nil {
		return -1, errors.New("Need to receive a packet before you can send one")
	}
	return this.udp.WriteToUDP(b, this.addr)
}

func (this *ICPConn) Close() {
	this.udp.Close()
}

func (this *ICPConn) LocalAddr() net.Addr {
	return this.udp.LocalAddr()
}

func (this *ICPConn) RemoteAddr() net.Addr {
	return this.udp.RemoteAddr()
}

func (this *ICPConn) SetDeadline(t time.Time) error {
	return this.udp.SetDeadline(t)
}

func (this *ICPConn) SetReadDeadline(t time.Time) error {
	return this.udp.SetReadDeadline(t)
}

func (this *ICPConn) SetWriteDeadline(t time.Time) error {
	return this.udp.SetWriteDeadline(t)
}
