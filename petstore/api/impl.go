package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// (POST /pets:validate)
func (s Server) ValidatePets(ctx echo.Context) error {
	var petnames PetNames
	var pets []Pet
	fmt.Printf("Calling ValidatePets....")
	err := ctx.Bind(&petnames)
	if err != nil {
		return sendPetStoreError(ctx, http.StatusBadRequest, "Invalid format for PetNames")
	}
	for i := range petnames.Names {
		pet := Pet{
			Name: petnames.Names[i],
		}
		pets = append(pets, pet)
	}

	fmt.Printf("%+v", pets[0])
	//fmt.Printf("%+v", pets[1])
	return ctx.JSON(http.StatusOK, &pets)
}

// (POST /pets:generate)
func (s Server) GeneratePets(ctx echo.Context) error {
	var petnames PetNames
	var pets []Pet
	fmt.Printf("Calling GeneratePets....")
	err := ctx.Bind(&petnames)
	if err != nil {
		return sendPetStoreError(ctx, http.StatusBadRequest, "Invalid format for PetNames")
	}
	for i := range petnames.Names {
		pet := Pet{
			Name: petnames.Names[i],
		}
		pets = append(pets, pet)
	}

	fmt.Printf("%+v", pets[1])
	return ctx.JSON(http.StatusOK, &pets)
}

func sendPetStoreError(ctx echo.Context, code int, message string) error {
	petErr := Error{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, petErr)
	return err
}
