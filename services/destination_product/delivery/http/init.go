package http

import (
	"github.com/gorilla/mux"
	"net/http"
	"si-test/services/destination_product/usecase"
)

type HttpDestinationProductHandler struct {
	usecase usecase.IDestinationProductUseCase
}

func NewDeliveryHttpArea(r *mux.Router, usecase usecase.IDestinationProductUseCase) HttpDestinationProductHandler {
	handler := HttpDestinationProductHandler{usecase: usecase}
	r.HandleFunc("/destination-product/update-all", handler.UpdateDestinationProduct).Methods(http.MethodGet)
	return handler
}
