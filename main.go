package main

import (
	"fmt"

	c "github.com/subhankardas/solid-design-principles/code"
)

func main() {
	SingleResponsibilityPrinciple()
	OpenClosedPrinciple()
}

func SingleResponsibilityPrinciple() {
	invoice := &c.ProductInvoice{} // Here invoice handles single responsibility
	invoice = invoice.CreateInvoice(123.45, 6)
	invoice.DeleteInvoice()

	fmt.Printf("invoice: %+v \n", invoice)
}

func OpenClosedPrinciple() {
	// !!! APPROACH VIOLATING OPEN-CLOSED PRINCIPLE !!!
	cart := &c.ShoppingCart{}
	cart.Prices = []float32{1.2, 3.4, 5.6}

	cart.Checkout() // Add new payment methods to this method will violate OCP

	// APPROACH COMPLYING WITH THE OPEN-CLOSED PRINCIPLE
	// HERE WE CAN ADD ANY NO. OF PAYMENT METHODS TO THE CHECKOUT FEATURE
	// WITHOUT ADDING OR MODIFYING EXISTING CHECKOUT LOGIC.
	cart = &c.ShoppingCart{}
	cart.Prices = []float32{7.8, 9.0, 1.2, 3.4}
	cart.CheckoutV2(&c.CreditCard{})

	cart = &c.ShoppingCart{}
	cart.Prices = []float32{5.6, 7.8, 9.0}
	cart.CheckoutV2(&c.Cash{})
}
