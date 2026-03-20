package api

import (
	"encoding/json"
	"net/http"

	"github.com/easy-cloud-Knet/KWS_Control/service"
	"github.com/easy-cloud-Knet/KWS_Control/structure"
)

type ApiShutdownVmRequest struct {
	UUID structure.UUID `json:"uuid"`
}

type ApiShutdownVmResponse struct {
	Message string `json:"message"`
}

type ApiForceShutdownVmRequest struct {
	UUID structure.UUID `json:"uuid"`
}

type ApiForceShutdownVmResponse struct {
	Message string `json:"message"`
}

func (c *handlerContext) shutdownVm(w http.ResponseWriter, r *http.Request) {
	var req ApiShutdownVmRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err := service.ShutdownVM(req.UUID, c.context, c.rdb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
