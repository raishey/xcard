package main

import (
    "log"
    "time"
)

func main() {
    log.Println("MiSTer FPGA service starting...")
    for {
        log.Println("Service is running on MiSTer FPGA")
        time.Sleep(10 * time.Second) // Log every 10 seconds
    }
}