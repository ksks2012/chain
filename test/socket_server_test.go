package socket_server

import (
	"testing"

	"github.com/block-chain/pkg/setting"
	"github.com/block-chain/pkg/socket"
)

func TestSocketServer(t *testing.T) {
	cfg := setting.SocketSettingS{
		Host: "127.0.0.1",
		Port: "1024",
	}
	socket.StartSocketServer(cfg)
}
