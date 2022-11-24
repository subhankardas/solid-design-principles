package code

type CustomerProfile struct{} // ---> THIS IS A CONCRETE CUSTOMER PROFILE STRUCT
type CustomerHistory struct{} // ---> THIS IS A CONCRETE CUSTOMER HISTORY STRUCT

// USER MODULE HAS TWO DEPENDENCIES I.E. PROFILE AND HISTORY
type User struct {
	profile CustomerProfile // HERE WE HAVE BOUNDED TO CONCRETE STRUCT INSTANCES
	history CustomerHistory // FOR BOTH PROFILE AND HISTORY
}

// CONSTRUCTOR ONLY ACCEPTS CONCRETE STRUCT/CLASSES
func NewUser(profile CustomerProfile, history CustomerHistory) *User {
	return &User{
		profile: profile, // !!! HERE WE ARE VIOLATING THE DEPENDENCY INVERSION PRINCIPLE
		history: history, // !!! HERE AGAIN
	}
}

// ----- REWRITING THE ABOVE CODE FOLLOWING DEPENDENCY INVERSION PRINCIPLE ----- //

type Profile interface{ GetEmail() string }
type History interface{ GetHistory() []string }

type UserV2 struct {
	Profile Profile // DEPENDENCIES POINTS TO INTERFACES IN STRUCT
	History History // AND NOT CONCRETE STRUCT/CLASSES
}

// CONSTRUCTOR ALSO ACCEPTS INTERFACES AS ARGUMENTS
func NewUserV2(profile Profile, history History) *UserV2 {
	return &UserV2{
		Profile: profile, // SO WE ARE GOOD TO PASS ANY IMPLEMENTATION
		History: history, // OF PROFILE(CUSTOMER/SELLER) AND HISTORY(PURCHASES/SALES)
	}
}

// IMPLEMENTATIONS OF PROFILE INTERFACE
type CustomerProfileV2 struct{}

func (profile *CustomerProfileV2) GetEmail() string { return "john@doe.com" }

type SellerProfile struct{}

func (profile *SellerProfile) GetEmail() string { return "mary@jane.com" }

// IMPLEMENTATIONS OF HISTORY INTERFACE
type Purchase struct{}

func (purchase *Purchase) GetHistory() []string { return []string{"iphone", "books", "laptop"} }

type Sales struct{}

func (sales *Sales) GetHistory() []string { return []string{"cakes", "pens"} }
