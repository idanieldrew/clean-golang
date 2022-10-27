package response

import (
	"clean-golang/app/infrastructure/logger"
	"encoding/json"
	"net/http"
)

func (r *ResponseS) Res(w http.ResponseWriter, status int, res any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		logger.Error("problem in response user")
		return
	}
	return
}
