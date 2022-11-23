package code

import "fmt"

// Since Go does not inherently support classes and inheritance.
// Hence we use composition to construct child struct using base struct as an anonymous field.
// Child struct can directly access base struct variables and methods.
type Order struct{}
type SpecialOrder struct{ Order }  // SPECIAL ORDER (CHILD)   <--- ORDER
type GiftCardOrder struct{ Order } // GIFT CARD ORDER (CHILD) <--- ORDER

func (order *Order) GetDeliveryAddress() { // BASE IMPLEMENTATION
	fmt.Println("v1: getting simple order address...")
}

func (order *SpecialOrder) GetDeliveryAddress() { // CHILD 1 OVERRIDE BASE IMPL.
	fmt.Println("v1: getting special order address...")
}

func (order *GiftCardOrder) GetDeliveryAddress() { // CHILD 2 OVERRIDE BASE IMPL.
	// !!! CANNOT IMPLEMENT SINCE GIFT CARDS DO NOT HAVE A DELIVERY ADDRESS !!!

	// Now we can change the working of the base GetDeliveryAddress() to support gift card orders
	// but that will violate the Liskov Substitution Principle since we couldn't replace the Order
	// base struct with GiftCardOrder child struct without breaking the application.
}

// ----- REWRITING THE ABOVE CODE FOLLOWING LISKOV SUBSTITUTION PRINCIPLE ----- //

type OrderV2 struct{}               // BASE ORDER
type ProductOrder struct{ OrderV2 } // PRODUCT (CHILD) <--- ORDER
type ServiceOrder struct{ OrderV2 } // SERVICE (CHILD) <--- ORDER

func (order *ProductOrder) GetDeliveryAddress() { fmt.Println("v2: getting product order address...") }
func (order *ServiceOrder) ActivateService()    { fmt.Println("v2: activating service order...") }

type SpecialOrderV2 struct{ ProductOrder }  // SPECIAL ORDER    <--- PRODUCT
type GiftCardOrderV2 struct{ ServiceOrder } // GIFT CARD ORDER  <--- SERVICE
