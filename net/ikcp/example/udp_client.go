package main

import (
	"fmt"
	"github.com/prestonTao/mandela/net/ikcp"
	"net"
	"time"
)

func main() {

	server := new(Server)
	server.Init()

}

type Server struct {
	kcp       *ikcp.Ikcpcb
	conn      net.Conn
	remotAddr *net.UDPAddr
}

func (this *Server) Init() {
	// 创建两个端点的 kcp对象，第一个参数 conv是会话编号，同一个会话需要相同
	// 最后一个是 user参数，用来传递标识
	a := []byte{0}
	// b := []byte{1}
	this.kcp = ikcp.Ikcp_create(0x11223344, a)
	// kcp2 := Ikcp_create(0x11223344, b)

	this.kcp.Output = this.send
	// kcp2.Output = send

	// 配置窗口大小：平均延迟200ms，每20ms发送一个包，
	// 而考虑到丢包重发，设置最大收发窗口为128
	ikcp.Ikcp_wndsize(this.kcp, 128, 128)
	// Ikcp_wndsize(kcp2, 128, 128)

	// 默认模式
	ikcp.Ikcp_nodelay(this.kcp, 0, 10, 0, 0)
	// Ikcp_nodelay(kcp2, 0, 10, 0, 0)

	remotAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9981")
	if err != nil {
		fmt.Println(err.Error())
	}
	locaAddr, err := net.ResolveUDPAddr("udp", ":9980")
	if err != nil {
		fmt.Println(err.Error())
	}
	this.conn, err = net.ListenUDP("udp", locaAddr)
	if err != nil {
		fmt.Println(err.Error())
	}

	this.Send("hello word!")

	time.Sleep(100 * time.Second)

	// this.recv(conn)

}

// func (this *Server) recv(conn net.Conn) {
// 	buffer := make([]byte, 1024)
// 	for {
// 		// conn.Read(b)
// 		_, err := conn.Read(buffer)
// 		if err != nil {
// 			fmt.Println("ReadFromUDP error ", err)
// 			return
// 		}
// 		// 如果 p1收到udp，则作为下层协议输入到kcp1
// 		ikcp.Ikcp_input(this.kcp, buffer, int(1024))
// 	}
// }

func (this *Server) print() {

}

// 发送一个 udp包
func (this *Server) send(buf []byte, _len int32, kcp *ikcp.Ikcpcb, user interface{}) int32 {
	// arr := (user).([]byte)
	// var id uint32 = uint32(arr[0])
	// //println("send!!!!", id, _len)
	// if vnet.send(int(id), buf, int(_len)) != 1 {
	// 	//println("wocao !!!", id, _len)
	// }
	return 0
}

func (this *Server) Send(msg string) {
	err := this.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println(err.Error())
	}
}
