package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/niroopreddym/blockchain/common"
	"golang.org/x/crypto/ripemd160"
)

const (
	checksumLength = 4
	version        = byte(0x00)
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

// ctor
func MakeWallet() *Wallet {
	private, public := NewKeyPair()
	wallet := Wallet{
		PrivateKey: private,
		PublicKey:  public,
	}

	return &wallet
}

// generate a random private key and public key based on ecdsa algorithm
// ecdsa radnom eliptical curve distribution set
// returns the private key and public key
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256() //curve would be 256 bytes in length
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	common.PanicErr(err)
	pub := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)
	return *privateKey, pub
}

// for creating address we need version and checksum and piblickeyhash
func (w *Wallet) Adress() []byte {
	pubHash := PublicKeyHash(w.PublicKey)
	versionedHash := append([]byte{version}, pubHash...)
	checkSum := checkSum(versionedHash)
	fullHash := append(versionedHash, checkSum...)
	address := Base58Encode(fullHash)

	return address
}

// PublicKeyHash takes pub key as input and generates a pub hash
func PublicKeyHash(publicKey []byte) []byte {
	//take the key pass through sha256 algo
	pubHash := sha256.Sum256(publicKey)
	//rehash with ripemd160
	hasher := ripemd160.New()
	_, err := hasher.Write(pubHash[:])
	common.PanicErr(err)
	//add a sum to the ripMDHash
	publicRipMD := hasher.Sum(nil)
	fmt.Printf("public Hash: %x\n", string(publicRipMD))
	return publicRipMD
}

// takes in the public key hash and generetes a checksum
// remember checksum is the first 4 bytes of the hash genereted after the calculations
func checkSum(publicKeyHash []byte) []byte {
	firstHash := sha256.Sum256(publicKeyHash)
	secondHash := sha256.Sum256(firstHash[:])
	return secondHash[:checksumLength]
}
