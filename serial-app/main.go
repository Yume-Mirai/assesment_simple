package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"th-backend-sji/serial-app/sdk"
	"time"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	client, err := sdk.New("/dev/ttyV98")
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return
	}
	defer func() {
		fmt.Println("Closing serial port...")
		if err := client.Close(); err != nil {
			fmt.Printf("Error closing port: %v\n", err)
		}
		// Give the port a moment to close properly
		time.Sleep(100 * time.Millisecond)
	}()

	for {
		select {
		case <-sigChan:
			fmt.Println("Received shutdown signal...")
			return
		default:
			fb, err := client.ReceiveEvent()
			if err != nil {
				fmt.Printf("Err: %v\n", err)
				continue
			}

			if fb != "" {
				fmt.Printf("Event: % X\n", []byte(fb))
			}

			time.Sleep(1 * time.Second)
		}
	}
}
