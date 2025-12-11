package main

// Important

import (
	"fmt"

	"github.com/202412013/Golang/auth"
	"github.com/202412013/Golang/user"
	"github.com/fatih/color"
)

// Packages : Conflict avoid kari sakiye. Compilation time speedup kari sakiye che packages na help thi.
// -> Go Aej packages ne recompile kare jema changes thaya hoi.

// -> Project etle modules and aapde init karvu pdse module ne
// -> go mod init github.com/codersgyan/projectname(reponame)
// -> Aa go.mod name ni file create kari de che jema module and version batave

// Login Feature and make a package (Folder -> Multiple files)
func main() {
	auth.LoginWithCredentials("codersgyan", "secret")

	session := auth.GetSession()
	fmt.Println("Session ->", session)

	user := user.User{
		Email: "user@gmail.com",
		Name:  "John Doe",
	}

	fmt.Println(user.Email, user.Name)

	// Package install like express in JS -> go get github.com/fatih/color

	color.Green(user.Email)

	// go.sum -> similar to package-lock.json

	// go mod tidy -> indirect dependency and direct dependency will be different now from this command

	// go mod tidy -> Aa work kare che npm install jm je javascript ma aapde kariye che em

	// Important Point to remember in Scope :-
	//	1) Capital and small Letter
	//	2) 3rd party packages (go get github.com/...)
	//	3) go mod tidy

}
