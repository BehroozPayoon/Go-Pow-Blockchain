package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	senderAddress    string
	recipientAddress string
	value            float32
}

func NewTransaction(senderAddress string, recipientAddress string, value float32) *Transaction {
	return &Transaction{
		senderAddress:    senderAddress,
		recipientAddress: recipientAddress,
		value:            value,
	}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_address       %s\n", t.senderAddress)
	fmt.Printf(" recipient_address    %s\n", t.recipientAddress)
	fmt.Printf(" value                %.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderAddress    string  `json:"sender_address"`
		RecipientAddress string  `json:"recipient_address"`
		value            float32 `json:"value"`
	}{
		SenderAddress:    t.senderAddress,
		RecipientAddress: t.recipientAddress,
		value:            t.value,
	})
}
