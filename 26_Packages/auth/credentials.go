package auth

import "fmt"

// Conventions :- agar mare aaj package ma function use karvu che toh hu function nu name small letter ma aapis
// 			 	  and baar use karvu hoi toh first letter capital ma aapis
func LoginWithCredentials(username string, password string) {
	fmt.Println("login user using", username, password)
}
