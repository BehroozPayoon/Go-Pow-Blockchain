package main

import (
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	b := NewBlock(0, "Init Hash")
	b.Print()
}
