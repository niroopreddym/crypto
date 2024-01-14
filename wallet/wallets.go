package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/niroopreddym/blockchain/common"
)

const (
	walletFile = "./tmp/wallets.data"
)

type Wallets struct {
	Wallets map[string]*Wallet
}

// ctor
func CreateWallets() *Wallets {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)
	wallets.LoadFile()
	return &wallets
}

func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

func (ws Wallets) AddWallet() string {
	wallet := MakeWallet()
	add := fmt.Sprintf("%s", wallet.Adress())
	ws.Wallets[add] = wallet
	return add
}

func (ws Wallets) LoadFile() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	var wallets Wallets
	fileContent, err := ioutil.ReadFile(walletFile)
	common.PanicErr(err)

	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	common.PanicErr(err)
	ws.Wallets = wallets.Wallets

	return nil
}

func (ws Wallets) GetAllAddress() []string {
	addresses := []string{}
	for add := range ws.Wallets {
		addresses = append(addresses, add)
	}

	return addresses
}

func (ws Wallets) SaveFile() {
	var content bytes.Buffer
	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(ws)
	common.PanicErr(err)

	err = os.WriteFile(walletFile, content.Bytes(), 0644)
	common.PanicErr(err)
}
