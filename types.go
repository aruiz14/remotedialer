package remotedialer

import "time"

const (
	PingWaitDuration  = 60 * time.Second
	PingWriteInterval = 5 * time.Second
	MaxRead           = 8192
	HandshakeTimeOut  = 10 * time.Second
	// SendErrorTimeout sets the maximum duration for sending an error message to close a single connection
	SendErrorTimeout = 5 * time.Second
)
