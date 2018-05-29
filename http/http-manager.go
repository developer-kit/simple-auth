package http

import (
	"github.com/seaofstars-coder/simple-auth/config"
	"github.com/seaofstars-coder/simple-auth/log"
	"net/http"
	"os"
	"sync"
)

type HttpManager struct {
	addr string
	wg *sync.WaitGroup
}

var m *HttpManager
var once sync.Once

func GetInstance() *HttpManager {
	once.Do(func() {
		m = &HttpManager{}
	})
	return m
}

func (self *HttpManager) InitManager(wg *sync.WaitGroup) error {
	self.wg = wg
	var err error = nil
	self.addr, err = config.GetInstance().GetConfig("AUTH_HTTP_ADDR")
	if nil != err {
		return err
	}
	http.HandleFunc("/auth", handleAuthRequest)
	http.HandleFunc("/verification", handleVerificationRequest)
	http.HandleFunc("/registry", handleRegistryRequest)
	self.wg.Add(1)
	go self.startHttpServer()
	return nil
}

func (self *HttpManager) startHttpServer() {
	log.Info("Start Http Server on addr ", self.addr)
	defer self.wg.Done()
	http.ListenAndServe(self.addr, nil)
}

func (self *HttpManager) outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}
