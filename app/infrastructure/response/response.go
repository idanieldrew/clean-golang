package response

import "net/http"

type (
	Response interface {
		Res(w http.ResponseWriter, code int, res any)
	}

	ResponseS struct {
	}
)
