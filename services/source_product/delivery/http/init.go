package http

import (
	"github.com/gorilla/mux"
	"net/http"
	domain "si-test/services/source_product/usecase"
)

type HttpSourceProductHandler struct {
	usecase domain.ISourceProductUseCase
}

func NewDeliveryHttpArea(r *mux.Router, usecase domain.ISourceProductUseCase) HttpSourceProductHandler {
	handler := HttpSourceProductHandler{usecase: usecase}
	r.HandleFunc("/source-product/get-all", handler.GetAllData).Methods(http.MethodGet)
	return handler
}
