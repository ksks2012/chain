package socket

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/block-chain/pkg/setting"
)

func StartSocketClient(cfgSetting setting.SocketSettingS) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", cfgSetting.Host+":"+cfgSetting.Port)
	if err != nil {
		log.Printf("Fatal error: %s", err.Error())
		os.Exit(1)
	}

	connection, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Printf("Fatal error: %s", err.Error())
		os.Exit(1)
	}
	defer connection.Close()
	fmt.Println("connectionect success")
	sender(connection)
}

func sender(connection net.Conn) {
	words := "hello world!"
	connection.Write([]byte(words))
	log.Printf("send over")
}
