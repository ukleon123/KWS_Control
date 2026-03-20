package api

import (
	"encoding/json"
	"net/http"

	"github.com/easy-cloud-Knet/KWS_Control/service"
	"github.com/easy-cloud-Knet/KWS_Control/structure"
)

type ApiDeleteVmRequest struct {
	UUID structure.UUID `json:"uuid"`
}

func (c *handlerContext) deleteVm(w http.ResponseWriter, r *http.Request) {
	var req ApiDeleteVmRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err := service.DeleteVM(req.UUID, c.context, c.rdb)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
