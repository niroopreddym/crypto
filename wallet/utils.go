package wallet

import (
	"github.com/mr-tron/base58"
	"github.com/niroopreddym/blockchain/common"
)

func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)
	return []byte(encode)
}

func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input))
	common.PanicErr(err)
	return decode
}
