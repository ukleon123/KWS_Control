package model

type HardwareInfo struct {
	CPU    int `json:"cpu"`
	Memory int `json:"memory"`
}

type UserInfoVM struct {
	Name              string   `json:"name"`
	Groups            string   `json:"groups"`
	Password          string   `json:"passWord"`
	SSHAuthorizedKeys []string `json:"ssh"`
}

type CreateVMRequest struct {
	DomType      string       `json:"domType"`
	DomName      string       `json:"domName"`
	UUID         string       `json:"uuid"`
	OS           string       `json:"os"`
	HardwareInfo HardwareInfo `json:"HWInfo"`
	NetConf      NetDefine    `json:"network"`
	Users        []UserInfoVM `json:"users"`
}

type NetDefine struct {
	Ips     []string `json:"ips"`
	NetType NetType  `json:"NetType"`
}

type NetType uint

type CreateVMResponse struct {
	State     string `json:"state"`
	MaxMem    uint64 `json:"maxmem"`
	Memory    uint64 `json:"memory"`
	NrVirtCpu uint   `json:"nrVirtCpu"`
	CpuTime   uint64 `json:"cpuTime"`
}
