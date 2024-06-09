package healthcheck

import (
	"context"
	"net/http"
)

func HealtCheckController(context context.Context, input *HealthCheckRequest) (*HealtCheckResponse, error) {
	return &HealtCheckResponse{Body: http.StatusText(http.StatusOK)}, nil
}
