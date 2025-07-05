package pkg

import (
	"fmt"
	"net/http"

	"github.com/Headliner38/go-project4.git/Desktop/ioBound/pkg/api"
)

const port = ":8080"

func RunServer() error {
	api.Init()
	fmt.Println("Server is running on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return err
	}
	return nil
}
