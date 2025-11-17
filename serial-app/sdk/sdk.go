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
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open(s.portName, mode)
	if err != nil {
		return err
	}
	s.port = port
	return nil
}

func (s *SDK) ReceiveEvent() (string, error) {
	buf := make([]byte, 128)
	n, err := s.port.Read(buf)
	if err != nil {
		return "", err
	}
	if n == 0 {
		return "", nil
	}
	return string(buf[:n]), nil
}

func (s *SDK) SendCommand(cmd string) error {
	_, err := s.port.Write([]byte(cmd))
	return err
}

func (s *SDK) Close() error {
	return s.port.Close()
}
