// test/gspro_test.go

package main_test

import (
   "testing"
   "github.com/jheck90/gspro/mock"
)

func TestGSProClient(t *testing.T) {
	// Create and start the simulated GSPro server
	simulatedServer, err := gspro.NewSimulatedGSProServer()
	if err != nil {
        t.Fatalf("Error creating simulated server: %v", err)
	}
	defer simulatedServer.Close()
	t.Log("Simulated GSPro server created and started successfully.")

	simulatedServer.Start()
    t.Log("Running main function...")

}