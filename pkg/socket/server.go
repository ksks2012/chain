package socket

import (
	"log"
	"net"

	"github.com/block-chain/pkg/setting"
)

func StartSocketServer(cfgSetting setting.SocketSettingS) {
	log.Printf("Usage: %s ip-addr\n", cfgSetting.Host)

	name := cfgSetting.Host
	addr := net.ParseIP(name)
	if addr == nil {
		log.Println("Invalid address")
	} else {
		log.Println("The address is ", addr.String())
	}

	netListen, err := net.Listen("tcp", cfgSetting.Host+":"+cfgSetting.Port)
	if err != nil {
		log.Println("Error listening:", err.Error())
	}
	for {
		connection, err := netListen.Accept()
		if err != nil {
			continue
		}
		log.Println(connection.RemoteAddr().String(), " tcp connection success")
		handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	buffer := make([]byte, 2048)

	for {
		n, err := connection.Read(buffer)
		if err != nil {
			log.Println(connection.RemoteAddr().String(), " connectionion error: ", err)
			return
		}
		log.Println(connection.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))

	}
}
