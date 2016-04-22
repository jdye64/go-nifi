package socket

import (
	"fmt"
	"github.com/jdye64/nifi/remote/client"
)

type SocketClient struct {
	siteToSiteConfig client.SiteToSiteClientConfig
	endpointConnectionPool EndpointConnectionPool
	compress bool
	portName string
	penalizationNanoes uint64
	portIdentifier string
	isClosed bool
}

func NewSocketClient() SocketClient {
	sc := new(SocketClient)

	//Creates a new EndpointConnectionPool


	sc.compress = true
	return *sc
}

// func CreateTransaction(sc SocketClient) remote.Transaction {
// 	return remote.Transaction{}
// }
	
func IsSecure(sc SocketClient) bool {
	return true
}

func GetConfig(sc SocketClient) client.SiteToSiteClientConfig {
	fmt.Printf("Getting SiteToSiteConfiguration for SiteToSiteClient")
	return sc.siteToSiteConfig
}