package main

import (
	"net"
)

type Client struct {
	conn        net.Conn
	server      *Server
	listeners   []func([]byte)
	seqId       int32
	remoteSeqId int32
}

func newClient(conn net.Conn, server *Server) {
	c := Client{conn, server, make([]func([]byte), 0), 0, 0}

	c.handshake()

	c.loop()
}

func (c Client) handshake() {
	// nothing for now??
	//c.Send(consts.SERVER_HI)
	//c.expect(consts.CLIENT_HI)
}

func (c Client) loop() {
	for {

	}
}

func (c Client) Send(b []byte) {
	c.conn.Write(b)
}

func (c Client) expect(b byte) {
	buf := make([]byte, 1)
	_, err := c.conn.Read(buf)

	if err != nil {
		panic(err)
	}

	if buf[0] != b {
		panic("Bad handshake")
	}
}
