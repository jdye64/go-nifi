package socket

import (
	
)

type EndpointConnectionPool struct {
	peerRefreshPeriod uint64
	category string
}

func NewEndpointConnectionPool() *EndpointConnectionPool {
	ep := new(EndpointConnectionPool)

	ep.category = "Site-to-Site"

	return ep
}