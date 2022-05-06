package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var UP = websocket.Upgrader{
	HandshakeTimeout:  0,     //握手时间0为不限制
	ReadBufferSize:    1024,  //以字节为单位的IO缓冲区,如果缓冲区大小为零，则使用HTTP服务器分配的缓冲区
	WriteBufferSize:   1024,  //以字节为单位的IO缓冲区,如果缓冲区大小为零，则使用HTTP服务器分配的缓冲区
	WriteBufferPool:   nil,   // WriteBufferPool是用于写操作的缓冲池。
	Subprotocols:      nil,   //按顺序指定服务器支持的协议
	Error:             nil,   //指定用于生成HTTP错误响应的函数
	CheckOrigin:       nil,   //对过来的请求做校验用的
	EnableCompression: false, //指定服务器是否应尝试根据进行协商消息压缩
}

func handler(rw http.ResponseWriter, r *http.Request) {
	conn, err := UP.Upgrade(rw, r, nil) //返回长连接和报错
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	if err != nil {
		log.Println(err)
		return
	}
	for { //持续从连接取消息
		message, p, err := conn.ReadMessage() //返回msgType  内容 和 err
		if err != nil {
			break
		}
		fmt.Println(message, string(p))
		conn.WriteMessage(websocket.TextMessage, []byte("吃了吗  你说的是:"+string(p)+"吗"))
	}
	log.Println("服务shut down service")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
