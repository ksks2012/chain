package blocker

import (
	"github.com/block-chain/pkg/rsakey"
)

var (
	PrivateKey []byte
	PublicKey  []byte
)

func GenAddress() {
	PrivateKey, PublicKey = rsakey.GenRsaKey()
}
