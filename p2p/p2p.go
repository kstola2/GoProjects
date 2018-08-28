package main

import (
	"net"
	"os"
	"strconv"
	"strings"
)

type Server struct {
	listener net.Listener
	conns    []net.Conn
	buf      chan []byte
}

func main() {
}

func getIP() string {
	s, err := net.Dial("tcp", "www.google.com:80")
	if err != nil {
		os.Stderr.WriteString("Error connecting to web")
	}
	locAddr := s.LocalAddr().String()
	return locAddr[:strings.IndexByte(locAddr, ':')]
}

func initServer(port int) net.Listener {
	portStr := ":" + strconv.Itoa(port)
	server, err := net.Listen("tcp", portStr)
	if err != nil {
		os.Stderr.WriteString("Error initializing server")
	}
	return server
}

func (s *Server) acceptConnections() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			os.Stderr.WriteString("Client attempted to connect")
		}
		s.conns = append(s.conns, conn)
	}
}

// func (s *Server) handleConnection(conn *net.Conn) {
// 	msg := make([]byte, 1024)
// 	for {
// 		numBytes, err := (*conn).Read(msg)
// 		if err != nil {
// 			break
// 		} else {
// 			data := make([]byte, numBytes)
// 			copy(data, msg[:numBytes])
// 			s.buf <- data
// 		}
// 	}
// }
