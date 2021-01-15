package main 

// Visit pkg.go.dev if acceptable for the project

// Chatty-GoGo requires GinGonic for the http framework

import (
	"fmt"
	"mygithub.com/Routes"
)

func main() {
	// This message will only be funny to old IBMers
	fmt.Println("Chatty-GoGo open for e-business!")  
	Routes.HandleCRUD()
}