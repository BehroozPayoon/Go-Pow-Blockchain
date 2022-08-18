package main

import (
	"fmt"
	"log"
	"pow-blockchain/core"
	"pow-blockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	// Wallet
	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.Address(), walletB.Address(), 1.0)

	// Blockchain
	blockchain := core.NewBlockchain(walletM.Address())
	isAdded := blockchain.AddTransaction(walletA.Address(), walletB.Address(), 1.0,
		walletA.PublicKey(), t.GenerateSignature())
	fmt.Println("Added? ", isAdded)

	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.Address()))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.Address()))
	fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.Address()))
}
