package iface

import (
	ma "github.com/multiformats/go-multiaddr/interface"
)

type DatagramHandler interface {
	ReadFrom([]byte) (int, ma.Multiaddr, error)
	WriteTo([]byte, ma.Multiaddr) (int, error)
}

type Multigram interface {
	AddDatagramHandler(string, DatagramHandler) error
}
