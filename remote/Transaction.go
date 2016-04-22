package remote

type Transaction interface {
	Send()
	Receive()
	Confirm()
	Complete()
	Cancel()
	GetState()
	GetCommunicant()
}