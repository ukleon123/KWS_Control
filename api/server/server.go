package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	WorkerCont "github.com/easy-cloud-Knet/KWS_Control/api/workercont"
	vms "github.com/easy-cloud-Knet/KWS_Control/vm"
)


func Server(portNum int, taskPool *WorkerCont.TaskHandler, contextStruct *vms.InfraContext ){
	// main server와 통신하기 위한 http 서버
	// gin.DefaultWriter = io.Discard

	http.HandleFunc("Get /getStatus",func(w http.ResponseWriter, r *http.Request){
 
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		workerControl:= &WorkerCont.TaskControl_GetStatus{
			ResultChann: make(chan WorkerCont.TaskExecutionResult),
		}
		resultChannel:= workerControl.ResultChann
		defer close(resultChannel)
		workerControl.TaskUnparsor(r)
		
		newTask:=&WorkerCont.Task{
			FunctionName: WorkerCont.GetStatus,	
			TaskSpecific: workerControl,
		}
		
		taskPool.WorkerAllocate(newTask)
		result:= <-resultChannel
		encoder := json.NewEncoder(w)
		encoder.Encode(result)
	})

	http.HandleFunc("POST /CreateVM",func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		workerControl:= &WorkerCont.TaskControl_CreateVM{
			ResultChann: make(chan WorkerCont.TaskExecutionResult),
		}
		resultChannel:= workerControl.ResultChann
		defer close(resultChannel)
		workerControl.TaskUnparsor(r)
		
		newTask:=&WorkerCont.Task{
			FunctionName: WorkerCont.CreateV,	
			TaskSpecific: workerControl,
		}
		
		taskPool.WorkerAllocate(newTask)
		result:= <-resultChannel
		encoder := json.NewEncoder(w)
		encoder.Encode(result)
	})
	
	http.HandleFunc("GET /DeleteVM",func(w http.ResponseWriter, b *http.Request){
		taskPool.WorkerAllocate(&WorkerCont.Task{ 
			FunctionName: WorkerCont.DeleteV,

		})
	})
	http.HandleFunc("GET /ConnectVM",func(w http.ResponseWriter, b *http.Request){
		taskPool.WorkerAllocate(&WorkerCont.Task{ 
			FunctionName: WorkerCont.ConnectV,

		})
	})
	http.HandleFunc("GET /CheckVMHealth", func(w http.ResponseWriter, b *http.Request){
		taskPool.WorkerAllocate(&WorkerCont.Task{ 
			FunctionName: WorkerCont.UpdateStat,
		})
	})


	http.ListenAndServe(":"+strconv.Itoa(portNum), nil)



}