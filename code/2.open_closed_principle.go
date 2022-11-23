package code

import "fmt"

// This below sample code works fine for processing checkout of cart items for an e-commerce application.
// Here we declare checkout as method for cart interface as its feature. ShoppingCart implements this
// interface and adds logic to add up list of prices and process payment in acceptCash().
type Cart interface{ Checkout() }
type ShoppingCart struct{ Prices []float32 }

func (cart *ShoppingCart) Checkout() {
	var amount float32
	for _, price := range cart.Prices {
		amount = amount + price
	}
	acceptCash(amount) // !! NOW IF WE WANT TO ADD CREDIT CARD SUPPORT !!
	// WE COULD JUST ADD "if" STATEMENT BUT WILL VIOLATE OPEN-CLOSED PRINCIPLE.
}

func acceptCash(amount float32) {
	fmt.Printf("accepting $%v cash... \n", amount)
}

// ----- REWRITING THE ABOVE CODE FOLLOWING OPEN-CLOSED PRINCIPLE ----- //

type PaymentMethod interface{ acceptPayment(amount float32) } // Payment method declaration
type Cash struct{}                                            // Cash payment method
type CreditCard struct{}                                      // Credit card method

func (cart *ShoppingCart) CheckoutV2(payment PaymentMethod) { // This implementation of checkout implies OCP
	var amount float32
	for _, price := range cart.Prices {
		amount = amount + price
	}
	payment.acceptPayment(amount) // HERE WE ARE OPEN TO PAYMENT METHOD EXTENSION BUT CLOSED FOR LOGIC MODIFICATION
	// FEATURE THAT ARE SUBJECT TO CHANGE IN BUSINESS REQUIREMENT IN FUTURE SHOULD USE DEPENDENCIES AS INTERFACES TO FOLLOW OCP.
}

func (cash *Cash) acceptPayment(amount float32) {
	fmt.Printf("accepting $%v as cash payment... \n", amount)
}

func (credit *CreditCard) acceptPayment(amount float32) {
	fmt.Printf("accepting $%v on credit card... \n", amount)
}
