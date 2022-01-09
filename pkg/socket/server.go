package socket

import (
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"

	"github.com/block-chain/pkg/setting"
)

var (
	waitGroup sync.WaitGroup = sync.WaitGroup{}
)

func StartSocketServer(cfgSetting setting.SocketSettingS) {
	stop_chan := make(chan os.Signal)
	signal.Notify(stop_chan, os.Interrupt)

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

	go func() {
		<-stop_chan
		log.Printf("Get Stop Command. Now Stoping...")
		if err = netListen.Close(); err != nil {
			log.Panic(err)
		}
	}()

	for {
		connection, err := netListen.Accept()
		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				break
			}
			log.Panic("accept error\n")
			continue
		}
		log.Println(connection.RemoteAddr().String(), " tcp connection success")
		waitGroup.Add(1)
		go handleConnection(connection)
	}

	waitGroup.Wait()
}

func handleConnection(connection net.Conn) {
	defer waitGroup.Done()
	defer connection.Close()
	addr := connection.RemoteAddr()
	// TODO: buffer size
	buffer := make([]byte, 2048)

	for {
		n, err := connection.Read(buffer)
		if err != nil {
			log.Println(addr.String(), " connectionion error: ", err)
			break
		}
		log.Println(addr.String(), "receive data string:\n", string(buffer[:n]))
	}
	log.Printf("%v disconnect:\n", addr.String())
}
