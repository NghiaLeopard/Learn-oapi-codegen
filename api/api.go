package api

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

type API struct{}

func (a *API) GetApiTest(ctx context.Context, request GetApiTestRequestObject) (GetApiTestResponseObject, error) {

	return &GetApiTest200JSONResponse{
		Email: "nghiabeo1605@gmail.com",
		Id:    1,
	}, nil
}

func (a *API) PostApiTest(ctx context.Context, request PostApiTestRequestObject) (PostApiTestResponseObject, error) {

	// example id is not exist

	if request.Body.Id == 0 {
		return &PostApiTest401JSONResponse{
			Message: "authorization",
		}, nil
	}
	return &PostApiTest200JSONResponse{
		Email: request.Body.Email,
		Id:    request.Body.Id,
	}, nil
}

func NewApi() *fiber.App {
	api := &API{}

	server := fiber.New()

	_, err := GetSwagger()

	if err != nil {
		log.Println("Swagger error")
	}

	server.Get("/swagger/*", swagger.HandlerDefault)

	server.Get("/swagger.json", func(c *fiber.Ctx) error {
		spec, err := GetSwagger()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error: %s", err))
		}

		return c.JSON(spec)
	})

	handler := NewStrictHandler(api, nil)
	RegisterHandlers(server, handler)

	return server
}
