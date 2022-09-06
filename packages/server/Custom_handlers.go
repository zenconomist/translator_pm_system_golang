package server

import (
	"dto"
	"encoding/json"
	"entities"
	"fmt"
	"net/http"
	"services"
)

// to PUT subrouter
func (sh ServerHandler) SetTaskState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                             // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add project init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "setting task state process started",
	}
	dto := dto.TaskStateRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}
	repo := entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Task{})
	task, errRepo := repo.FindByID(dto.TaskID)
	if errRepo != nil {
		defer logger.LogError(errRepo)
		res.Message = fmt.Sprintf("repo error: %v", errRepo)
		json.NewEncoder(w).Encode(res)
		return
	}

	tss := services.NewTaskStateDbBasedService(&task, sh.env)
	if errSetState := tss.SetState(dto.ToState); errSetState != nil {
		defer logger.LogError(errSetState)
		res.Message = fmt.Sprintf("service error: %v", errSetState)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("task state set successfully to: %v", dto.ToState)
	json.NewEncoder(w).Encode(res)

}

// to PUT subrouter
func (sh ServerHandler) SetTaskOfferedState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                             // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add project init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "setting task state process started",
	}
	dto := dto.TaskOfferedStateRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}
	repo := entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.TaskOffered{})
	to, errRepo := repo.FindByID(dto.ID)
	if errRepo != nil {
		defer logger.LogError(errRepo)
		res.Message = fmt.Sprintf("repo error: %v", errRepo)
		json.NewEncoder(w).Encode(res)
		return
	}
	toss := services.NewTaskOfferingStateService(&to, sh.env)
	if errSetState := toss.SetState(dto.ToState); errSetState != nil {
		defer logger.LogError(errSetState)
		res.Message = fmt.Sprintf("set state service error: %v", errSetState)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("task offer " + dto.ToState)
	json.NewEncoder(w).Encode(res)

}
