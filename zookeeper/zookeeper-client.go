package zookeeper

import (
	"errors"
	"github.com/samuel/go-zookeeper/zk"
	"sync"
	"time"
)

type ZookeeperClient struct {
	host    []string
	timeout time.Duration
	conn    *zk.Conn
	conSucc bool
	mutex   sync.Mutex
}

func (self *ZookeeperClient) Init(host []string, timeout int) error {
	self.host = host
	self.timeout = time.Duration(timeout)
	return self.connect()
}

func (self *ZookeeperClient) connect() error {
	var err error = nil
	self.conn, _, err = zk.Connect(self.host, time.Second*self.timeout)
	if nil != err {
		self.conSucc = false
		return err
	}
	self.conSucc = true
	return err
}

func (self *ZookeeperClient) GetValue(path string) ([]byte, error) {
	if !self.conSucc {
		return nil, errors.New("zookeeper client has not connected to server!!!")
	}
	data, _, err := self.conn.Get(path)
	if nil != err {
		return nil, err
	}
	return data, err
}

func (self *ZookeeperClient) GetChildren(path string) ([]string, error) {
	if !self.conSucc {
		return nil, errors.New("zookeeper client has not connected to server!!!")
	}
	children, _, err := self.conn.Children(path)
	if nil != err {
		return nil, err
	}
	return children, err
}

func (self *ZookeeperClient) SetValue(path, value string) error {
	if !self.conSucc {
		return errors.New("zookeeper client has not connected to server!!!")
	}
	if _, err := self.conn.Create(path, []byte(value), 0, zk.WorldACL(zk.PermAll)); err != nil {
		return err
	}
	return nil
}

func (self *ZookeeperClient) Delete(path string) error {
	if !self.conSucc {
		return errors.New("zookeeper client has not connected to server!!!")
	}
	return self.conn.Delete(path, 0)
}

func (self *ZookeeperClient) Children(path string) (<-chan zk.Event, error) {
	if !self.conSucc {
		return nil,errors.New("zookeeper client has not connected to server!!!")
	}
	_, _, ch, err := self.conn.ChildrenW(path)
	return ch, err
}
