package api

import (
	"encoding/json"
	"net/http"

	"github.com/easy-cloud-Knet/KWS_Control/service"
	"github.com/easy-cloud-Knet/KWS_Control/structure"
	"github.com/easy-cloud-Knet/KWS_Control/util"
)

type ApiVmInfoRequest struct {
	UUID structure.UUID `json:"uuid"`
}

type ApiVmInfoResponse struct {
	UUID   structure.UUID `json:"uuid"`
	CPU    uint32         `json:"cpu"`
	Memory uint32         `json:"memory"` // MiB
	Disk   uint32         `json:"disk"`   // MiB
	IP     string         `json:"ip"`
}

func (c *handlerContext) vmInfo(w http.ResponseWriter, r *http.Request) {
	log := util.GetLogger()

	var req ApiVmInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		log.Error("Invalid request body: %v", err, true)
		return
	}
	defer r.Body.Close()

	vmInfo, err := service.GetVMInfoFromRedis(r.Context(), c.rdb, req.UUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Error("failed to get vm info from redis: %v", err, true)
		return
	}

	response := ApiVmInfoResponse{
		UUID:   vmInfo.UUID,
		CPU:    vmInfo.CPU,
		Memory: vmInfo.Memory,
		Disk:   vmInfo.Disk,
		IP:     vmInfo.IP,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Error("failed to encode vm info response: %v", err, true)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	log.Info("retrieved vm info from redis: UUID=%s", string(req.UUID), true)
}
