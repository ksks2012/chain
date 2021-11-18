package global

type Transaction struct {
	Sender   string
	Receiver string
	Amounts  int64 //TODO:
	Fee      int64 //TODO:
	Message  string
}
