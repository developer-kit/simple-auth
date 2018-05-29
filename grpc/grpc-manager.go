package grpc

import (
	"errors"
	"github.com/seaofstars-coder/simple-auth/config"
	"github.com/seaofstars-coder/simple-auth/proto"
	"google.golang.org/grpc"
	"net"
	"sync"
	"github.com/seaofstars-coder/simple-auth/log"
)

type GRPCManager struct {
	wg *sync.WaitGroup
	listener net.Listener
	server *grpc.Server
	addr string
}

var m *GRPCManager
var once sync.Once

func GetInstance() *GRPCManager {
	once.Do(func() {
		m = &GRPCManager{}
	})
	return m
}

func (self *GRPCManager) InitManager(wg *sync.WaitGroup) error {
	self.wg = wg
	var err error = nil
	self.addr, err = config.GetInstance().GetConfig("GRPC_ADDR")
	if nil != err {
		return err
	}
	self.listener, err = net.Listen("tcp", self.addr)
	if nil != err {
		return err
	}
	self.server = grpc.NewServer()
	if nil == self.server {
		return errors.New("grpc.NewServer return nil!!!")
	}
	proto.RegisterAuthServer(self.server, &AuthGrpcServer{})
	proto.RegisterVerificationServer(self.server, &VerificationGrpcServer{})
	self.wg.Add(1)
	go self.startServe()
	return nil
}

func (self *GRPCManager) startServe() {
	log.Info("Start GRPC Server on addr ", self.addr)
	defer self.wg.Done()
	self.server.Serve(self.listener)
}
