package socket_client

import (
	"testing"

	"github.com/block-chain/pkg/setting"
	"github.com/block-chain/pkg/socket"
)

func TestSocketClient(t *testing.T) {
	cfg := setting.SocketSettingS{
		Host: "127.0.0.1",
		Port: "1024",
	}
	socket.StartSocketClient(cfg)
}
