package transaction

import (
	"crypto"
)

type transaction struct {
	transactionId string
	value         float64
	info          string
	blockNumber   int
	sender        crypto.PrivateKey
	recipient     crypto.PublicKey
	signature     []byte
	sequence      int
}
