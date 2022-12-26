package http

import (
	"net/http"
	"si-test/pkg/utils"
)

func (d HttpSourceProductHandler) GetAllData(w http.ResponseWriter, r *http.Request) {
	data, err := d.usecase.GetAllProduct()
	if err != nil {
		utils.ResponseMapErrorBasedOnError(w, err)
		return
	}
	utils.RespondSuccessCallbackWithJSON(w, http.StatusOK, data)
	return
}
