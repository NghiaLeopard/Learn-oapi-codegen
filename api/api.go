package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type API struct{}

func (a *API) GetApiTest(ctx context.Context, request GetApiTestRequestObject) (GetApiTestResponseObject, error) {

	return &GetApiTest200JSONResponse{
		Email: "nghiabeo1605@gmail.com",
		Id:    1,
	}, nil
}

func NewApi() *fiber.App {
	api := &API{}

	server := fiber.New()

	handler := NewStrictHandler(api, nil)
	RegisterHandlers(server, handler)

	return server
}