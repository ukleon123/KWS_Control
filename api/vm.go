package server

import (
	"net/http"
	"strconv"

	"github.com/easy-cloud-Knet/KWS_Control/util"
	vms "github.com/easy-cloud-Knet/KWS_Control/vm"
)

func Server(portNum int, contextStruct *vms.ControlInfra) error {
	http.HandleFunc("/vm", func(w http.ResponseWriter, r *http.Request) {
		if !util.CheckMethod(w, r, http.MethodPost) {
			return
		}
})

	err := http.ListenAndServe(":"+strconv.Itoa(portNum), nil)
	if err != nil {
		return err
	}

	return nil
}
