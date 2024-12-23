package WorkerCont

import (
	"encoding/json"
	"net/http"
	"sync"

	vms "github.com/easy-cloud-Knet/KWS_Control/vm"
)
type GetStatus_Param struct{
	UUID string `json:"UUID"`
}
type CreateVM_Param struct{
	UUID string `json:"UUID"`
	RAM int `json:"RAM"`
	CPU int `json:"CPU"`
}
//parameters for each functions


type Task struct{
	FunctionName functionName 
	TaskSpecific TaskJustifier
}
type TaskWorker struct{
	taskLenMu sync.Mutex
	tasksLength int
	workLoads  chan *Task
	workerNum int
}

type TaskHandler struct{
	TaskHandlersList []*TaskWorker
	workingIndex int
}//테스크 임베딩, 필드 추가 필요


type TaskControl_CreateVM struct{
	ResultChann chan TaskExecutionResult
	UUID vms.UUID
	RAM int
	CPU int
	//추가 필요
}
type TaskControl_GetStatus struct{
	ResultChann chan TaskExecutionResult
	UUID vms.UUID
}



type TaskJustifier interface{
	TaskUnparsor(r *http.Request) error
}

//functions for each structures, needed for interface
func (t *TaskControl_CreateVM) TaskUnparsor(r *http.Request) error{
	var param CreateVM_Param
	if err:= json.NewDecoder(r.Body).Decode(&param); err!=nil{
		return err
	}
	t.RAM= param.RAM
	t.CPU= param.CPU
	return nil
}

func (t *TaskControl_GetStatus) TaskUnparsor(r *http.Request) error{
	var param GetStatus_Param
	if err := json.NewDecoder(r.Body).Decode(&param); 
	err != nil {
		//에러 정의 필요
		return err
	}
	return nil
}




type TaskExecutionResult struct{
	IsSuccess bool
	InVMContext vms.VMInfo 
}






