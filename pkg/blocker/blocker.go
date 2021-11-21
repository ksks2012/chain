package blocker

import (

	// "github.com/block-chain/global"

	"github.com/block-chain/pkg/rsakey"
)

var (
	privateKey []byte
	publicKey  []byte
)

func GenAddress() {
	privateKey, publicKey = rsakey.GenRsaKey()
}
