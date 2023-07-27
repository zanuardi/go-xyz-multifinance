package middleware

import (
	"net/http"

	"github.com/zanuardi/go-xyz-multifinance/helper"
	"github.com/zanuardi/go-xyz-multifinance/model/response"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "X-Secret-Key" == r.Header.Get("X-API-Key") {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := response.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized.",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
