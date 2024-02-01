// test/gspro_test.go

package test

import (
   "testing"
   "github.com/jheck90/gspro"
)

func TestRun1(t *testing.T) {
    // Create and start the simulated GSPro server
    simulatedServer, err := gspro.NewSimulatedGSProServer()
    if err != nil {
        t.Fatal("Error creating simulated server:", err)
    }
    defer simulatedServer.Close()
    simulatedServer.Start()

    // Call the run function and assert the expected behavior
    if err := gspro.Run(); err != nil {
        t.Errorf("Expected no error, but got: %v", err)
    }
}

func TestRunWithError(t *testing.T) {
    // Create and start the simulated GSPro server
    simulatedServer, err := gspro.NewSimulatedGSProServer()
    if err != nil {
        t.Fatal("Error creating simulated server:", err)
    }
    defer simulatedServer.Close()
    simulatedServer.Start()

    // Call the run function with an error and assert the expected behavior
    if err := gspro.RunWithError(); err == nil {
        t.Error("Expected an error, but got none")
    } else {
        expectedErrorMessage := "Simulated connection failure"
        if err.Error() != expectedErrorMessage {
            t.Errorf("Expected error message '%s', but got: %v", expectedErrorMessage, err)
        }
    }
}

func TestRun(t *testing.T) {
    // Call the run function and assert the expected behavior

    // Example: Test the run function when everything is successful
    if err := gspro.Run(); err != nil {
        t.Errorf("Expected no error, but got: %v", err)
    }

    // Example: Test the run function when there's an error
    // You can modify your run function to return an error in certain scenarios
    // For example, if the connection fails
    // if err := gspro.RunWithError(); err == nil {
    //     t.Error("Expected an error, but got none")
    // }
}