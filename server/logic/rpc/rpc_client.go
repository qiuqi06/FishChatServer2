package rpc

import (
	"github.com/golang/glog"
	"github.com/oikomi/FishChatServer2/server/logic/rpc/client"
)

type RPCClient struct {
	Register *client.RegisterRPCCli
	Manager  *client.ManagerRPCCli
}

func NewRPCClient() (c *RPCClient, err error) {
	register, err := client.NewRegisterRPCCli()
	if err != nil {
		glog.Error(err)
		return
	}
	manager, err := client.NewManagerRPCCli()
	if err != nil {
		glog.Error(err)
		return
	}
	c = &RPCClient{
		Register: register,
		Manager:  manager,
	}
	return
}
