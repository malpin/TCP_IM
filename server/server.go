package server

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

//创建一个server的接口
func NewServer(ip string, port int) *Server {
	s := &Server{
		Ip:   ip,
		Port: port,
	}
	return s
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go s.HandleConn(conn)
	}
}

func (s *Server) HandleConn(conn net.Conn) {
	fmt.Println("链接")
}
