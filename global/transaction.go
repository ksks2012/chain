package global

import "fmt"

type Transaction struct {
	Sender   string
	Receiver string
	Amounts  int64 //TODO:
	Fee      int64 //TODO:
	Message  string
}

func (t Transaction) transactionToString() (transactionstring string) {
	transactionstring = fmt.Sprintf("%v%v%v%v%v", t.Sender, t.Receiver, t.Amounts, t.Fee, t.Message)
	return
}
