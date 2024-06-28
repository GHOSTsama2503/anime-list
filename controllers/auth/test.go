package auth

import (
	"context"
	"net/http"
)

type TestRequest struct{}

type TestResponse struct {
	Body string
}

func (controller *AuthController) Test(ctx context.Context, request *TestRequest) (*TestResponse, error) {
	response := &TestResponse{}
	response.Body = http.StatusText(http.StatusOK)
	return response, nil
}
