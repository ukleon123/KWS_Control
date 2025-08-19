package model

import "github.com/easy-cloud-Knet/KWS_Control/structure"

type ApiDeleteVmRequest struct {
	UUID structure.UUID `json:"uuid"`
}

type ApiShutdownVmRequest struct {
	UUID structure.UUID `json:"uuid"`
}

type ApiVmStatusRequest struct {
	UUID structure.UUID `json:"uuid"`
	Type string         `json:"type"` // "cpu", "memory", or "disk"
}

type ApiVmConnectRequest struct {
	UUID structure.UUID `json:"uuid"`
}

// 요건 core쪽이랑 이야기 맞춰야--
// request/model/vm.go 와 동일하게 유지
const (
	VMStatusPrepareBegin     = "prepare begin"
	VMStatusStartBegin       = "start begin"
	VMStatusStarted          = "started begin"
	VMStatusStopped          = "stopped end"
	VMStatusReleaseShutdown  = "release end shutdown"
	VMStatusReleaseDestroyed = "release end destroyed"
	VMStatusMigrate          = "migrate begin"
	VMStatusRestore          = "restort begin"
	VMStatusUnknown          = "unknown"
)

type Redis struct {
	UUID   structure.UUID `json:"UUID"`
	Status string         `json:"status"`
}

func ValidateAndNormalizeStatus(status string) string {
	if status == "" || status == "null" {
		return VMStatusUnknown
	}

	switch status {
	case VMStatusPrepareBegin, VMStatusStartBegin, VMStatusStarted, VMStatusStopped,
		VMStatusReleaseShutdown, VMStatusReleaseDestroyed, VMStatusMigrate, VMStatusRestore, VMStatusUnknown:
		return status
	default:
		return VMStatusUnknown
	}
}
