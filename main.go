package main

import (
	"fmt"

	pkg "github.com/Headliner38/go-project4.git/Desktop/ioBound/pkg/server"
)

func main() {
	err := pkg.RunServer()
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}
