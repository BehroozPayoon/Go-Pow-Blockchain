package main

import (
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchain := NewBlockchain("My Address")
	blockchain.Print()

	blockchain.AddTransaction("A", "B", 1.0)
	blockchain.Mining()
	blockchain.Print()

	blockchain.AddTransaction("C", "D", 4.0)
	blockchain.AddTransaction("X", "Y", 2.0)
	blockchain.Mining()
	blockchain.Print()
}
