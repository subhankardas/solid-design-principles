package main

import (
	"fmt"

	c "github.com/subhankardas/solid-design-principles/code"
)

func main() {
	SingleResponsibilityPrinciple()
}

func SingleResponsibilityPrinciple() {
	invoice := &c.ProductInvoice{}
	invoice = invoice.CreateInvoice(123.45, 6)
	invoice.DeleteInvoice()

	fmt.Printf("invoice: %+v", invoice)
}
