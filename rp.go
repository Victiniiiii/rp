package rp

import (
	"github.com/Victiniiiii/rp/ipc"
	"github.com/Victiniiiii/rp/rpc"
)

func NewClient(ID string) (*rpc.Client, error) {
	c := &rpc.Client{ClientID: ID}

	i, err := ipc.NewIPC()
	if err != nil {
		return nil, err
	}

	c.IPC = i

	if err := c.Login(); err != nil {
		return nil, err
	}

	return c, nil
}
