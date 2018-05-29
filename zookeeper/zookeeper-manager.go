package zookeeper

import (
	"encoding/json"
	"github.com/seaofstars-coder/simple-auth/config"
	"github.com/seaofstars-coder/simple-auth/log"
	"github.com/samuel/go-zookeeper/zk"
	"sync"
)

type ZookeeperManager struct {
	zkClient      ZookeeperClient
	clientDataMap ZKClientDataMap
}

var m *ZookeeperManager
var once sync.Once

func GetInstance() *ZookeeperManager {
	once.Do(func() {
		m = &ZookeeperManager{}
	})
	return m
}

func (self *ZookeeperManager) InitManager() error {
	self.clientDataMap = make(ZKClientDataMap)
	addrs, err := config.GetInstance().GetConfigArray("ZOOKEEPER_ADDRS")
	if nil != err {
		return err
	}
	timeout, err := config.GetInstance().GetInt("ZOOKEEPER_TIMEOUT")
	if nil != err {
		return err
	}
	err = self.zkClient.Init(addrs, timeout)
	if nil != err {
		return err
	}
	self.loadAllClientData()
	return nil
}

func (self *ZookeeperManager) GetClientSecret(clientID string) (string, error) {
	zkClientData := self.clientDataMap[clientID]
	if nil != zkClientData {
		return zkClientData.ClientSecret, nil
	}
	clientData, err := self.zkClient.GetValue(clientID)
	if nil != err {
		return "", err
	}
	zkClientData = new(ZKClientData)
	err = json.Unmarshal(clientData, zkClientData)
	if nil != err {
		self.clientDataMap[clientID] = zkClientData
		return zkClientData.ClientSecret, err
	}
	return "", err
}

func (self *ZookeeperManager) loadAllClientData() {
	path, err := config.GetInstance().GetConfig("ZOOKEEPER_CLIENTS_PATH")
	if nil != err {
		log.Error(err.Error())
		return
	}
	children, err := self.zkClient.GetChildren(path)
	if nil != err {
		log.Error(err.Error())
		return
	}
	for _, key := range children {
		data, err := self.zkClient.GetValue(path + "/" + key)
		if nil != err {
			continue
		}
		zkClientData := new(ZKClientData)
		json.Unmarshal(data, zkClientData)
		self.clientDataMap[zkClientData.ClientId] = zkClientData
		log.Info(zkClientData)
	}
	self.WatchChildren(path)
}

func (self *ZookeeperManager) WatchChildren(path string) {
	go func() {
		for {
			ch, err := self.zkClient.Children(path)
			if nil != err {
				continue
			}
			self.Watcher(ch)
		}
	}()
}

func (self *ZookeeperManager) Watcher(childCh <-chan zk.Event) {
	ev := <-childCh
	if ev.Err != nil {
		log.Error("Child watcher error:", ev.Err.Error())
		return
	}
	self.ZookeeperCallback(ev)
}

func (self *ZookeeperManager) ZookeeperCallback(event zk.Event) {
	switch event.Type {
	case zk.EventNodeChildrenChanged:
		{
			self.ProcessChildrenChanged()
		}
		break
	case zk.EventNodeCreated:
	case zk.EventNodeDeleted:
	case zk.EventNodeDataChanged:
	case zk.EventSession:
	case zk.EventNotWatching:
	default:
		{

		}
	}
}

func (self *ZookeeperManager) ProcessChildrenChanged() {
	path, err := config.GetInstance().GetConfig("ZOOKEEPER_CLIENTS_PATH")
	if nil != err {
		return
	}
	// Don not process delete client op
	children, err := self.zkClient.GetChildren(path)
	if nil != err {
		return
	}
	for _, key := range children {
		if nil == self.clientDataMap[key] {
			data, err := self.zkClient.GetValue(path + "/" + key)
			if nil != err {
				continue
			}
			zkClientData := new(ZKClientData)
			json.Unmarshal(data, zkClientData)
			self.clientDataMap[zkClientData.ClientId] = zkClientData
			log.Info("AddClient:", key)
		}
	}
}
