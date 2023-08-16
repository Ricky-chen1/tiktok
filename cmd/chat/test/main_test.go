package test

import (
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/kitex_gen/chat/messageservice"
)

var conn messageservice.Client

func TestMain(m *testing.M) {
	// 连接服务器
	c, err := messageservice.NewClient("chat", client.WithHostPorts("0.0.0.0:8889"))
	if err != nil {
		panic(err)
	}

	conn = c
	m.Run()
}
