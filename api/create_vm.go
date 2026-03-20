package api

import (
	"net/http"

	"github.com/easy-cloud-Knet/KWS_Control/service"
	"github.com/easy-cloud-Knet/KWS_Control/util"
)

func (c *handlerContext) createVm(w http.ResponseWriter, r *http.Request) {
	log := util.GetLogger()

	err := service.CreateVM(w, r, c.context, c.rdb)
	if err != nil {
		h := w.Header()
		h.Del("Content-Length")
		h.Set("Content-Type", "text/plain; charset=utf-8")
		h.Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusMethodNotAllowed)

		log.Error("Failed to create VM: %v", err, true)
		return
	}
	defer r.Body.Close()

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("VM created successfully"))
}
