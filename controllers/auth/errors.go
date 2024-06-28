package auth

import (
	"github.com/danielgtaylor/huma/v2"
)

var EmptyUsernameOrPasswordError = huma.Error400BadRequest("empty username or password")
var InvalidUsernameOrPasswordError = huma.Error403Forbidden("invalid username or password")
