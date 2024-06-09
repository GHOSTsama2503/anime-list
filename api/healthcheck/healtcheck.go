package healthcheck

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func Use(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "healtcheck",
		Method:      http.MethodGet,
		Path:        "/healthcheck",
		Summary:     "Check api health",
	}, HealtCheckController)
}
