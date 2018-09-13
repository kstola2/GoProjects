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

// func main() {
// 	conn, err := net.Dial("udp", "0.0.0.0:00")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	localAddr := conn.LocalAddr()
// 	conn.Close()

// 	fmt.Println(localAddr.String())
// }

func getIP() string {
	s, err := net.Dial("tcp", "www.google.com:80")
	if err != nil {
		os.Stderr.WriteString("Error connecting to web")
	}
	locAddr := s.LocalAddr().String()
	return locAddr[:strings.IndexByte(locAddr, ':')]
}

// func main() {
// 	conn, err := net.Dial("udp", "0.0.0.0:00")
// 	defer conn.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	localAddr := conn.LocalAddr()

// 	host, port,err:=net.SplitHostPort(localAddr.String())
// 	if err!=nil{
// 		fmt.Println("Could not split host and port given the local address string",err)

// 	}
// 	fmt.Println(host, port, err)
// }

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
