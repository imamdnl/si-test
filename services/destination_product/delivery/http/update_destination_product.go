package http

import (
	"net/http"
	"si-test/pkg/utils"
)

func (d HttpDestinationProductHandler) UpdateDestinationProduct(w http.ResponseWriter, r *http.Request) {
	err := d.usecase.Update()
	if err != nil {
		utils.ResponseMapErrorBasedOnError(w, err)
		return
	}
	utils.RespondSuccessCallbackWithJSON(w, http.StatusOK, true)
	return
}
