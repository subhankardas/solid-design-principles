package code

// !!! THIS INTERFACE VIOLATES INTERFACE SEGREGATION PRINCIPLE !!!
type Processor interface {
	CheckStock(item string) bool
	UpdateStock(item string, count int)
	BalanceStock()
	InitiatePayment(method string)
	AcknowledgePayment(transactionID string)
}

//	HERE THIS CLIENT IS FORCED TO IMPLEMENT METHODS THAT ARE NOT RELATED TO STOCK.
type StockManager struct{}

func (stock *StockManager) CheckStock(item string) bool             { return true }
func (stock *StockManager) UpdateStock(item string, count int)      {}
func (stock *StockManager) BalanceStock()                           {}
func (stock *StockManager) InitiatePayment(method string)           {} // THESE METHODS ARE UNNECESSARILY IMPLEMENTED
func (stock *StockManager) AcknowledgePayment(transactionID string) {} // JUST TO USE THE OTHER METHODS OF THIS INTERFACE

// ----- REWRITING THE ABOVE CODE FOLLOWING INTERFACE SEGREGATION PRINCIPLE ----- //

// SEGREGATING INTERFACES BASED ON RELATED FUNCTIONALITIES
type Stock interface {
	CheckStock(item string) bool
	UpdateStock(item string, count int)
	BalanceStock()
}

type Payment interface {
	InitiatePayment(method string)
	AcknowledgePayment(transactionID string)
}
