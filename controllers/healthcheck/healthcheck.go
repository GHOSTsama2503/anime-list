package healthcheck

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type HealthCheckRequest struct {
}

type HealtCheckResponse struct {
	Body string
}

func HealtCheckController(context context.Context, input *HealthCheckRequest) (*HealtCheckResponse, error) {
	return &HealtCheckResponse{Body: http.StatusText(http.StatusOK)}, nil
}

func Use(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "healtcheck",
		Method:      http.MethodGet,
		Path:        "/healthcheck",
		Summary:     "check api health",
	}, HealtCheckController)
}
