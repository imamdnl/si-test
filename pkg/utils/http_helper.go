package utils

import (
	"net/http"
	"si-test/pkg/exception"
)

func ResponseMapErrorBasedOnError(w http.ResponseWriter, err error) {
	if domainErr, ok := err.(*exception.DomainError); ok {
		RespondWithSingleError(w, http.StatusBadRequest, domainErr)
	} else if businessErr, ok := err.(*exception.BusinessError); ok {
		RespondWithSingleError(w, http.StatusBadRequest, businessErr)
	} else if infrastructureErr, ok := err.(*exception.InfrastructureError); ok {
		RespondWithSingleError(w, http.StatusInternalServerError, infrastructureErr)
	} else {
		RespondWithSingleError(w, http.StatusInternalServerError, exception.NewInfrastructureError("Internal server error"))
	}
}
