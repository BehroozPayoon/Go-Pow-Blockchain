package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	address    string
}

func NewWallet() *Wallet {
	w := &Wallet{}

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &privateKey.PublicKey
	w.generateAddress()

	return w
}

func (w *Wallet) generateAddress() {
	hash := sha256.New()
	hash.Write(w.publicKey.X.Bytes())
	hash.Write(w.publicKey.Y.Bytes())
	digest := hash.Sum(nil)

	ripHash := ripemd160.New()
	ripHash.Write(digest)
	ripDigest := ripHash.Sum(nil)

	vd := make([]byte, 21)
	vd[0] = 0x00
	copy(vd[1:], ripDigest[:])

	vdHash := sha256.New()
	vdHash.Write(vd)
	vdDigest := vdHash.Sum(nil)

	lastHash := sha256.New()
	lastHash.Write(vdDigest)
	lastDigest := lastHash.Sum(nil)

	checkSum := lastDigest[:4]

	dc := make([]byte, 25)
	copy(dc[:21], vd[:])
	copy(dc[21:], checkSum)

	w.address = base58.Encode(dc)
}

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *Wallet) PublickKetStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}

func (w *Wallet) Address() string {
	return w.address
}
