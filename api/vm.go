package api

import (
	"github.com/easy-cloud-Knet/KWS_Control/service"
	"net/http"
)

func (c *handlerContext) createVm(w http.ResponseWriter, r *http.Request) {
	err := service.CreateVM(w, r, c.context)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("VM created successfully"))
}
