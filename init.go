package freya

import (
	"github.com/kataras/iris"
	"github.com/taplion-ggao/freya/backend"
	"github.com/taplion-ggao/freya/decoder"
	"github.com/taplion-ggao/freya/global"
	"github.com/taplion-ggao/freya/handler"
	"github.com/taplion-ggao/freya/log"
	"github.com/taplion-ggao/freya/net"
	"github.com/taplion-ggao/freya/utils"
)

func NewApp() {
	// 初始化配置
	global.InitGlobal()
	// 其他初始化
	utils.InitIdWorker()
	utils.InitIdWorkerLow()
}

func NewTcpServer(handlers ...handler.TcpHandler) {
	// 初始化tcp服务
	net.InitTcpServer(handlers...)
	// 初始化decoder
	decoder.InitDecoder(handlers...)
}

func NewHttpServer(handlers ...handler.HttpHandler) {
	net.RegHttpHandlers(handlers...)
}

func NewHttpServerParty(party iris.Party, handlers ...handler.HttpHandler) {
	net.RegHttpHandlersParty(party, handlers...)
}

func NewHttpParty(h handler.HttpHandler) iris.Party {
	return net.NewHttpParty(h)
}

func RunApp() {
	// 开启tcp服务
	net.StartTcpServer()
	log.GetLogger().Debugf("freya listenTcp port:%d", global.G.TcpPort)
	// 开启http服务
	net.StartHttpServer()
	log.GetLogger().Debugf("freya listenHttp port:%d", global.G.HttpPort)
	// 启动backend服务
	backend.Start()

	log.GetLogger().Infof("Server:%d Start OK...", global.G.ServerId)
	<-net.Done
}

func SetBackend(be backend.Backend) {
	backend.SetBackend(be)
}