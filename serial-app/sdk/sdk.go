package sdk

import (
	"go.bug.st/serial"
)

type SDK struct {
	portName string

	port serial.Port
}

func New(portName string) (*SDK, error) {
	sdk := &SDK{
		portName: portName,
	}
	return sdk, sdk.connect()
}

func (s *SDK) connect() error {
	// TODO: implement connecting to device
	return nil
}

func (s *SDK) ReceiveEvent() (string, error) {
	// TODO: receive feedback from device and convert to event name

	return "", nil
}

func (s *SDK) SendCommand(cmd string) error {
	// TODO: implement

	return nil
}

func (s *SDK) Close() error {
	return s.port.Close()
}
