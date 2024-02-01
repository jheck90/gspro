package gspro

import (
	"encoding/json"
	"fmt"
	"net"
)

// SimulatedGSProServer represents a mock GSPro server.
type SimulatedGSProServer struct {
	listener net.Listener
}

// NewSimulatedGSProServer creates a new instance of SimulatedGSProServer.
func NewSimulatedGSProServer() (*SimulatedGSProServer, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0921")
	if err != nil {
		return nil, err
	}
	return &SimulatedGSProServer{listener: listener}, nil
}

// Start starts the simulated GSPro server.
func (s *SimulatedGSProServer) Start() {
	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				fmt.Println("Error accepting connection:", err)
				continue
			}
			go s.handleConnection(conn)
		}
	}()
}

func (s *SimulatedGSProServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	// Simulate sending sample JSON data
	sampleData := map[string]interface{}{
		"Code":    200,
		"Message": "Success",
	}

	jsonData, err := json.Marshal(sampleData)
	if err != nil {
		fmt.Println("Error encoding sample data:", err)
		return
	}

	_, err = conn.Write(jsonData)
	if err != nil {
		fmt.Println("Error sending sample data:", err)
	}
}

// Close closes the simulated GSPro server.
func (s *SimulatedGSProServer) Close() {
	s.listener.Close()
}
