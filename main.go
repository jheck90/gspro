package main

import (
	"encoding/json"
	"fmt"
	"bytes"
	"net"
)

func main() {
    // Connect to GSPro on 127.0.0.1:0921
    conn, err := net.Dial("tcp", "127.0.0.1:0921")
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    defer conn.Close()

    // Read incoming data from GSPro
    buffer := make([]byte, 1024)
    bytesRead, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error reading data:", err)
        return
    }

    // Parse JSON data using Decoder
    var data map[string]interface{}
    decoder := json.NewDecoder(bytes.NewReader(buffer[:bytesRead]))
    if err := decoder.Decode(&data); err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

    // Check the GSPro response code
    responseCode := int(data["Code"].(float64))
    if responseCode == 200 {
        fmt.Println("Mission success!")
        // Log relevant information, e.g., successful shot details
    } else if responseCode == 201 {
        fmt.Println("Player info received.")
        // Log player information
    } else if responseCode >= 501 && responseCode < 600 {
        fmt.Println("Mission failure. Reassess and adapt.")
        // Log the failure details
    }
}