package main

import (
	"net"
	"reflect"
	"strconv"
	"testing"
)

// Should return a string type when calling get IP function
func TestGetIP(t *testing.T) {
	testIP := getIP()
	subjType := reflect.TypeOf(testIP)
	if subjType.Name() != "string" {
		t.Errorf("Expected type 'IP' but got %v", subjType)
	}
}

// Should create a server with the correct port passed to it
func TestInitServer(t *testing.T) {
	testPort := 9000
	testServer := initServer(testPort)
	expected := "[::]:" + strconv.Itoa(testPort)
	actual := testServer.Addr().String()
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

// Should accept an incoming connection
func TestAcceptConnections(t *testing.T) {
	s := Server{
		listener: initServer(9000),
	}
	testIP := getIP()
	testAddr := testIP + ":9000"
	_, err := net.Dial("tcp", testAddr)
	(&s).acceptConnections()
	if err != nil {
		t.Errorf("Could not connect to server")
	}
	// timeElapsed := 0
	// for len(s.conns) < 1 {
	// 	time.Sleep(time.Second)
	// 	timeElapsed += 1
	// 	if timeElapsed > 5 {
	// 		t.Errorf("Connection too slow or not possible.")
	// 	}
	// 	continue
	// }
	// resultAddr := conn.RemoteAddr().String()
	// if resultAddr != testAddr {
	// 	t.Errorf("Error: test address does not match")
	// }
}
