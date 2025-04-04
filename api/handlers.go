package api

import (
	"net/http"
	"strconv"

	vms "github.com/easy-cloud-Knet/KWS_Control/structure"
)

type handlerContext struct {
	context *vms.ControlInfra
}

func Server(portNum int, contextStruct *vms.ControlInfra) error {
	h := handlerContext{
		context: contextStruct,
	}
	http.HandleFunc("POST /vm", h.createVm)

	err := http.ListenAndServe(":"+strconv.Itoa(portNum), nil)
	if err != nil {
		return err
	}

	return nil
}
