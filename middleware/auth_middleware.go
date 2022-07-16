package middleware

import (
	"learn-golang-restful-api/helper"
	"learn-golang-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Hander http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Hander: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	apiKey := "RAHASIA"
	if apiKey == request.Header.Get("X-API-Key") {
		middleware.Hander.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
