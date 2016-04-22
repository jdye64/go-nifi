package remote

import (
	"fmt"
)

type PeerDescription struct {
	hostname string
	port int
	secure bool
}

func GetHostname(pd PeerDescription) string {
	return pd.hostname
}

func GetPort(pd PeerDescription) int {
	return pd.port
}

func String(pd PeerDescription) {
	fmt.Printf("PeerDescription toString()")
}

//TODO: Actually implement this
func HashCode() int {
	return 31
}