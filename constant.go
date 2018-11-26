package bloom

type UpdateFlag uint8

const (
	None         UpdateFlag = 0
	All          UpdateFlag = 1
	P2PubKeyOnly UpdateFlag = 2
)
