package main

import (
	"github.com/moowiz/gophysx"
)

type Service interface {
	Init()
	ProcessMessages(*Client, []Message) bool
	GetServiceType() ServiceType
}

type ServiceType byte

const (
	PhysxServiceType = 1
)

type PhysxService struct {
	system *gophysx.System
}

func (this *PhysxService) GetServiceType() ServiceType {
	return PhysxServiceType
}

func (this *PhysxService) Init() {
	this.system = gophysx.Init(Clock{})
}

func (this *PhysxService) ProcessMessages(c *Client, msgs []Message) bool {
	return true
}
