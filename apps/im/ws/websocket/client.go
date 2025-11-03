package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"net/url"
)

type Client interface {
	Close() error

	Send(v any) error
	Read(v any) error
}

type client struct {
	*websocket.Conn
	host string

	opt dailOption

	sendChan chan []byte
	closeCh  chan struct{}
}

func NewClient(host string, opts ...DailOptions) *client {
	opt := newDailOptions(opts...)

	c := client{
		Conn:     nil,
		host:     host,
		opt:      opt,
		sendChan: make(chan []byte, 1000),
		closeCh:  make(chan struct{}),
	}

	conn, err := c.dail()
	if err != nil {
		panic(err)
	}

	go c.writePump()
	c.Conn = conn
	return &c
}

func (c *client) writePump() {
	for {
		select {
		case msg := <-c.sendChan:
			if err := c.safeWrite(msg); err != nil {
				//写失败，尝试自动重连
				conn, err := c.dail()
				logx.Errorf("write pump err:%v", err)
				c.Conn = conn
				_ = c.safeWrite(msg)
			}
		case <-c.closeCh:
			return
		}
	}
}

// 实际执行写操作
func (c *client) safeWrite(msg []byte) error {
	return c.Conn.WriteMessage(websocket.TextMessage, msg)
}

func (c *client) dail() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: c.host, Path: c.opt.pattern}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), c.opt.header)
	return conn, err
}

func (c *client) Send(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	select {
	case c.sendChan <- data:
		return nil
	default:
		return errors.New("send message is full")
	}
}

//func (c *client) Send(v any) error {
//	data, err := json.Marshal(v)
//	if err != nil {
//		return err
//	}
//	err = c.WriteMessage(websocket.TextMessage, data)
//	if err == nil {
//		return nil
//	}
//	// todo: 再增加一个重连发送
//	conn, err := c.dail()
//	if err != nil {
//		return err
//	}
//	c.Conn = conn
//	return c.WriteMessage(websocket.TextMessage, data)
//}

func (c *client) Read(v any) error {
	_, msg, err := c.Conn.ReadMessage()
	if err != nil {
		return err
	}

	return json.Unmarshal(msg, v)
}
