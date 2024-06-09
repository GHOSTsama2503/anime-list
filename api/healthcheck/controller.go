package healthcheck

import (
	"context"
	"net/http"
)

func HealtCheckController(context context.Context, input *HealtCheckResponse) (*HealtCheckResponse, error) {
	return &HealtCheckResponse{Body: http.StatusText(http.StatusOK)}, nil
}
