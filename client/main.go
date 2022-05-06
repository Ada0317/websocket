package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

func main() { //创建一个dialer  拨号器
	dl := websocket.Dialer{ //根据业务拉进行学习
		NetDial:           nil,
		NetDialContext:    nil,
		NetDialTLSContext: nil,
		Proxy:             nil,
		TLSClientConfig:   nil,
		HandshakeTimeout:  0,
		ReadBufferSize:    0,
		WriteBufferSize:   0,
		WriteBufferPool:   nil,
		Subprotocols:      nil,
		EnableCompression: false,
		Jar:               nil,
	}
	conn, _, err := dl.Dial("ws://127.0.0.1:8888", nil) //使用websocket协议请求url 建立连接
	if err != nil {
		log.Println(err)
		return
	}

	//拿到了连接之后 就可以通过链接对服务器发生消息
	err = conn.WriteMessage(websocket.TextMessage, []byte("我来了")) //messageType int, data []byte
	go send(conn)
	//接下来模拟一个用命令行输入进行长连接交流
	for {
		message, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(message, string(p))
	}
}

func send(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		line, _, _ := reader.ReadLine()
		conn.WriteMessage(websocket.TextMessage, line) //messageType int, data []byte
	}
}
