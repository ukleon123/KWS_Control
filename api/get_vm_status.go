package api

import (
	"encoding/json"
	"net/http"

	"github.com/easy-cloud-Knet/KWS_Control/service"
	"github.com/easy-cloud-Knet/KWS_Control/structure"
	"github.com/easy-cloud-Knet/KWS_Control/util"
)

type ApiVmStatusRequest struct {
	UUID structure.UUID `json:"uuid"`
	Type string         `json:"type"` // "cpu", "memory", or "disk"
}

func (c *handlerContext) vmStatus(w http.ResponseWriter, r *http.Request) {
	log := util.GetLogger()

	var req ApiVmStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Error("Failed to decode request body: %v", err, true)
		return
	}
	defer r.Body.Close()

	statusType := req.Type
	if statusType != "cpu" && statusType != "memory" && statusType != "disk" {
		http.Error(w, "Invalid status type. Must be 'cpu', 'memory', or 'disk'", http.StatusBadRequest)
		return
	}

	var data any
	var err error

	switch statusType {
	case "cpu":
		data, err = service.GetVMCpuInfo(req.UUID, c.context)
	case "memory":
		data, err = service.GetVMMemoryInfo(req.UUID, c.context)
	case "disk":
		data, err = service.GetVMDiskInfo(req.UUID, c.context)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error("Failed to get VM status: %v", err, true)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Error("Failed to encode response: %v", err, true)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
