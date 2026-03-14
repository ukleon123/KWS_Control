package api

import (
	"encoding/json"
	"net/http"

	"github.com/easy-cloud-Knet/KWS_Control/service"
	"github.com/easy-cloud-Knet/KWS_Control/structure"
	"github.com/easy-cloud-Knet/KWS_Control/util"
)

type ApiVmConnectRequest struct {
	UUID structure.UUID `json:"uuid"`
}

type ApiVmConnectResponse struct {
	AuthToken string `json:"authToken"`
}

func (c *handlerContext) vmConnect(w http.ResponseWriter, r *http.Request) {
	log := util.GetLogger()

	uuidStr := r.URL.Query().Get("uuid")
	if uuidStr == "" {
		http.Error(w, "Missing 'uuid' query parameter", http.StatusBadRequest)
		log.Error("Missing 'uuid' query parameter", nil, true)
		return
	}

	uuid := structure.UUID(uuidStr)
	authToken, err := service.GetGuacamoleToken(uuid, c.context)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error("Failed to get Guacamole token: %v", err, true)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ApiVmConnectResponse{AuthToken: authToken}); err != nil {
		log.Error("Failed to encode response: %v", err, true)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
