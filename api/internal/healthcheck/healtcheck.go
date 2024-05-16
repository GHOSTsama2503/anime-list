package healthcheck

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type HealtCheckResponse struct {
	Body string
}

type HealtCheckRequest struct{}

func HealtCheckController(context context.Context, input *HealtCheckRequest) (*HealtCheckResponse, error) {
	return &HealtCheckResponse{Body: "OK"}, nil
}

func Use(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "healtcheck",
		Method:      http.MethodGet,
		Path:        "/healthcheck",
		Summary:     "Check api health",
	}, HealtCheckController)
}
