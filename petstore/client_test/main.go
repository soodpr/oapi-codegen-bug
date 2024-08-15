package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/soodpr/oapi-codegen-bug/petstore/api"
)

func main() {
	// custom HTTP client
	hc := http.Client{}

	// or to get a struct with the parsed response body
	{
		c, err := api.NewClientWithResponses("http://localhost:8080", api.WithHTTPClient(&hc))
		if err != nil {
			log.Fatal(err)
		}
		petNames := api.PetNames{Names: []string{"dog", "cat"}}

		resp, err := c.ValidatePetsWithResponse(context.TODO(), petNames)
		if err != nil {
			log.Fatal(err)
		}
		if resp.HTTPResponse.StatusCode != http.StatusOK {
			log.Fatalf("Expected HTTP 200 but received %d", resp.HTTPResponse.StatusCode)
		}

		fmt.Printf("resp.JSON200: %v\n", resp.JSON200)

		petNames1 := api.PetNames{Names: []string{"pigeon", "Parrot"}}
		resp1, err1 := c.GeneratePetsWithResponse(context.TODO(), petNames1)
		if err1 != nil {
			log.Fatal(err1)
		}
		if resp1.HTTPResponse.StatusCode != http.StatusOK {
			log.Fatalf("Expected HTTP 200 but received %d", resp1.HTTPResponse.StatusCode)
		}

		fmt.Printf("resp.JSON200: %v\n", resp1.JSON200)
	}

}
