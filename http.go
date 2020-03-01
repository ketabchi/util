package util

import (
	"net"
	"net/http"
)

func NewSocketTransPort(socketPath string) *http.Transport {
	dial := func(network, addr string) (net.Conn, error) {
		return net.Dial("unix", socketPath)
	}

	return &http.Transport{
		Dial:    dial,
		DialTLS: dial,
	}
}
