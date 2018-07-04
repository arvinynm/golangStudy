package ipc

import (
	"testing"
	"fmt"
)

type EchoServer struct {

}

func (server *EchoServer) Handle(method, params string) *Response{
	return &Response{"ok","this is the server"}
}

func (Server *EchoServer) Name() string{
	return "EchoServer"
}

func TestIpc(t *testing.T){
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1,_ := client1.Call("From Client1","params1")
	resp2,_ := client1.Call("From Client2","params2")

	fmt.Println("客户端1请求内容：",resp1.Body)
	fmt.Println("客户端2请求内容：",resp2.Body)
	client1.Close()
	client2.Close()
}