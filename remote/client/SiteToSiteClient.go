package client

import (
	//"github.com/jdye64/nifi/remote"
)

type SiteToSiteClient interface {
	//CreateTransaction() remote.Transaction
	IsSecure() bool
	GetConfig() SiteToSiteClientConfig
}