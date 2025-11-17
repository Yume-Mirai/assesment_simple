#!/bin/bash

# Build virtual-device
cd virtual-device && go build -o sim-device . && cd ..

# Create virtual port first
socat pty,raw,echo=0,link=/dev/ttyV99 pty,raw,echo=0,link=/dev/ttyV98 &

# Wait for ports to be created
sleep 2

# Start virtual-device with the second port
./virtual-device/sim-device /dev/ttyV99 &

# Wait for virtual-device to initialize
sleep 2

# Start serial-app with hot reload
cd serial-app && air -c .serial-app.toml