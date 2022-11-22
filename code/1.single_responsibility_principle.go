package code

// This interface defines multiple methods with different responsibilities. As the single responsibility principle
// says, your class or method should have only one responsibility and only one reason to change this does not satisfy
// the single responsibility principle.
type Invoice interface {
	CreateInvoice(price float32, quantity int) *Invoice // ----> This method should be a part of invoice interface.
	DeleteInvoice() bool                                // ----> This method should also be a part of invoice since both of them are responsible for invoice creation and deletion.
	GenerateReport()                                    // ----> This method should move to a separate interface/struct.
	EmailReport(email string)                           // ----> This method should move to a separate interface/struct.
}

// This struct should not implement the Invoice interface, rather we will redefine the interface.
type ProductInvoice struct {
	price    float32
	quantity int
}

// ----- REWRITING THE ABOVE CODE FOLLOWING SINGLE RESPONSIBILITY PRINCIPLE ----- //
type SRPInvoice interface { // Here invoice includes methods only responsible for invoice related tasks.
	CreateInvoice(price float32, quantity int) *ProductInvoice
	DeleteInvoice() bool
}

type Report interface { // We created a separate interface to handle reporting tasks.
	GenerateReport(invoice *SRPInvoice)
}

type Email interface { // And another interface to handle emailing functionality.
	EmailReport(email string, report *Report)
}

// Now our ProductInvoice struct can implement the SRPInvoice interface.
func (invoice *ProductInvoice) CreateInvoice(price float32, quantity int) *ProductInvoice {
	return &ProductInvoice{
		price:    price,
		quantity: quantity,
	}
}

func (invoice *ProductInvoice) DeleteInvoice() bool {
	return true
}
