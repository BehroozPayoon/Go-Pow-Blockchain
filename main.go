package main

import (
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchain := NewBlockchain()
	blockchain.Print()

	blockchain.AddTransaction("A", "B", 1.0)
	previousHash := blockchain.LastBlock().Hash()
	nonce := blockchain.ProofOfWork()
	blockchain.CreateBlock(nonce, previousHash)
	blockchain.Print()

	blockchain.AddTransaction("C", "D", 4.0)
	blockchain.AddTransaction("X", "Y", 2.0)
	previousHash = blockchain.LastBlock().Hash()
	nonce = blockchain.ProofOfWork()
	blockchain.CreateBlock(nonce, previousHash)
	blockchain.Print()
}
