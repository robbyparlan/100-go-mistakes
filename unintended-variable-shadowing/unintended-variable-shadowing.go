package UnintendedVariableShadowing

import "log"

var isActive bool = true
var UserRoles string = "admin"
var Message string

func main() {

	// these condition unintended variable shadowing
	if isActive {
		UserRoles := "user"
		log.Println(UserRoles)
	} else {
		UserRoles := "guest"
		log.Println(UserRoles)
	}

	// this is not unintended variable shadowing
	if isActive {
		UserRoles = "user"
		Message = "the role is " + UserRoles
		log.Println(Message)
	} else {
		UserRoles = "guest"
		Message = "the role is " + UserRoles
		log.Println(Message)
	}
}
