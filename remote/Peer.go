package remote

import (
	"fmt"
)

type Peer struct {
	PeerDescription
	url string
	clusterUrl string
	host string
	port int
	closed bool
}

func GetUrl(p Peer) string {
	fmt.Printf("Peer URL " + p.url)
	return p.url
}

func GetClusterUrl(p Peer) string {
	return p.clusterUrl
}