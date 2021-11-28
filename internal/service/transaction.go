package service

import (
	"fmt"
	"log"

	"github.com/block-chain/pkg/blocker"
)

type Transaction struct {
	Sender   []byte
	Receiver string // TODO: []byte
	Amounts  int64  //TODO:
	Fee      int64  //TODO:
	Message  string
}

func (t Transaction) transactionToString() (transactionstring string) {
	transactionstring = fmt.Sprintf("%v%v%v%v%v", t.Sender, t.Receiver, t.Amounts, t.Fee, t.Message)
	return
}

func (t Transaction) New(sender string, receiver string, amount int64, fee int64, message string) {
	t.Sender = blocker.PublicKey
	t.Receiver = receiver
	t.Amounts = amount
	t.Fee = fee
	t.Message = message
}

func InitialTransaction(sender []byte, receiver string, amount int64, fee int64, message string) Transaction {
	if MainChain.GetSurplus(sender) < amount+fee {
		log.Printf("Surplus not enough!")
		return Transaction{}
	}
	return Transaction{
		Sender:   sender,
		Receiver: receiver,
		Amounts:  amount,
		Fee:      fee,
		Message:  message,
	}
}
