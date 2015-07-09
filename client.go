package main

import (
	"container/list"
	"fmt"
)

type Client struct {
	icpConn  ICPConn
	services *list.List
	messages map[ServiceType][]Message
}

func newClient(icp ICPConn) *Client {
	fmt.Println("Getting byte")
	byt, addr, err := icp.ReadByte()
	if err != nil {
		panic(err)
	}

	if byt != CLIENT_HI {
		panic("Invalid client hello")
	}

	fmt.Println("Sending byte!!")
	err = icp.WriteByte(SERVER_HI, addr)
	if err != nil {
		panic(err)
	}

	return &Client{icp, list.New(), make(map[ServiceType][]Message)}
}

func (c *Client) AddService(s Service) {
	c.services.PushBack(s)
}

func (c *Client) getMessagesForService(s Service) []Message {
	typ := s.GetServiceType()
	msgs := c.messages[typ]
	delete(c.messages, typ)
	return msgs
}

func (c *Client) handshake() bool {
	// nothing for now??
	return true
}

func (c *Client) Start() {
	c.loop()
}

func (c *Client) readMessages() {
	for c.icpConn.HasMessages() {
		msg, err := c.icpConn.GetNextMessage()

		if err != nil {
			// TODO: log error
			continue
		}

		if buf, ok := c.messages[msg.GetServiceType()]; ok {
			buf = append(buf, msg)
		} else {
			c.messages[msg.GetServiceType()] = []Message{msg}
		}
	}
}

func (c *Client) loop() {
	for {
		c.readMessages()

		for e := c.services.Front(); e != nil; e = e.Next() {
			s := e.Value.(Service)
			if s.ProcessMessages(c, c.getMessagesForService(s)) {
				c.services.Remove(e)
			}
		}
	}
}
