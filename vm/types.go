package vms

import (
	libvirt "libvirt.org/go/libvirt"
)

type UUID string

type InfraContext struct {
	Computers         []Computer
	VMPoolUnallocated []*VM
	VMPoolAllocated   []*VM
	VMPool            map[UUID]*VM
}

//모든 컴퓨터, vm에 대한 정보를 담고 있음.
// json이나 yaml에 지속적으로 업데이트해서 재부팅시에도 유지하는 것 필수

type InfraManage interface {
	UpdateList()
}

type VM struct {
	VMInfo      VMInfo   `json:"vmInfo"`
	IsAlive     bool     `json:"isAlive"`
	IsAllocated bool     `json:"isAllocated"`
	IsLocatedAt Computer `json:"isLocatedAt"`
}

//VM 내부의 상세 정보 표시
type VMInfo struct {
	State     libvirt.DomainState `json:"state"`
	MaxMem    uint64              `json:"maxmem"`
	Memory    uint64              `json:"memory"`
	NrVirtCpu uint                `json:"nrVirtCpu"`
	CpuTime   uint64              `json:"cpuTime"`
	UUID      UUID                `json:"uuid"`
	IP        string              `json:"ip"`
}

type Computer struct {
	Name      string `json:"name"`
	Allocated []VM   `json:"allocated"`
	IP        string `json:"ip"`
	MAC       string `json:"mac"`
	IsAlive   bool   `json:"isAlive"`
}
