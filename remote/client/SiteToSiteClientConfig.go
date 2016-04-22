package client

type SiteToSiteClientConfig interface {
	GetUrl() string
	GetTimeout() int64
	GetIdleConnectionExpiration() int64
	GetPenalizedPeriod() int64
	//GetSslContext() SSLContext
	IsUseCompression() bool
	GetPortName() string
	GetPortIdentifier() string
	GetPreferredBatchDuration() int64
	GetPreferredBatchSize() int64
	GetPreferredBatchCount() int
	//GetEventReport() EventReporter
}