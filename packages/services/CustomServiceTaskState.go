package services

import (
	"entities"
	"environment"
	"fmt"
)

type IStateService interface {
	SetState(string) error
}

type TaskStateDbBasedService struct {
	Env  environment.Environment
	Task *entities.Task
}

func NewTaskStateDbBasedService(task *entities.Task, env environment.Environment) *TaskStateDbBasedService {
	return &TaskStateDbBasedService{
		Env:  env,
		Task: task,
	}
}

func (tss *TaskStateDbBasedService) SetState(state string) error {
	// we have to query the TaskStateChanges db with params:
	// 1. current task state from the task and
	// 2. given state of the state parameter
	// given state parameter has to be the 2 character task_state_code
	// representation of a code.

	var tsc entities.TaskStateChanges
	db := tss.Env.GiveDbHandler().PassConnection()
	result := db.Where("from_task_state_id = ? AND to_task_state_code = ?", tss.Task.TaskStateID, state).First(&tsc)
	if result.Error != nil {
		return result.Error
	}
	// the result of task_state_changes will indicate if the state change is ok or not
	// if state change is ok -> task's former state will be set to it's current state
	if tsc.IsAllowed {
		// if state change is ok -> task's former state will be set to it's current state
		db.Model(tss.Task).Update("former_task_state", tss.Task.TaskStateID)
		// and task's current state will be updated to the current state
		db.Model(tss.Task).Update("task_state_id", tsc.ToTaskStateCode)
	} else {
		return fmt.Errorf(tsc.StateChangeInfo)
	}

	// and actions will be performed -> kinds of actions below:
	// 1. State checks on project
	// 2. State checks on batches
	// 3. execution of notifications
	return nil
}

type TaskStateService struct {
	Task *entities.Task
	Env  environment.Environment
	// TscDefault entities.TaskStateConst
}

func NewTaskStateService(task *entities.Task, env environment.Environment) *TaskStateService {
	return &TaskStateService{
		Task: task,
		Env:  env,
		// TscDefault: entities.TaskStInvalid,
	}
}

func (tss *TaskStateService) SetState(state string) error {
	/*
		taskRepo := entities.NewRepository(tss.Env.GiveDbHandler().PassConnection(), entities.Task{})

		// converting the given new state into a TaskStateConst int type
		validState, errValid := tss.TscDefault.StringToState(state)
		if errValid != nil {
			return errValid
		}

		// converting the persisted TaskStateID into state object
		tss.Task.TaskState = tss.TscDefault.UintToState(tss.Task.TaskStateID)

		// set inner params of the state object (task and state const)
		tss.Task.TaskState.SetStateParams(tss.Task, validState)

		// in order to be able to return errors in the switch statement
		var errToState error

		switch validState {
		case entities.TaskStOpen:
			if errToState = tss.Task.TaskState.SetTaskStateToOpen(); errToState != nil {
				return errToState
			}
		case entities.TaskStOffered:
			ts := tss.Task.TaskState
			if errToState = ts.SetTaskStateToOffered(); errToState != nil {
				return errToState
			}
		case entities.TaskStInProgress:
			if errToState = tss.Task.TaskState.SetTaskStateToInProgress(); errToState != nil {
				return errToState
			}
		case entities.TaskStReady:
			if errToState = tss.Task.TaskState.SetTaskStateToReady(); errToState != nil {
				return errToState
			}
		case entities.TaskStDelivered:
			if errToState = tss.Task.TaskState.SetTaskStateToDelivered(); errToState != nil {
				return errToState
			}
		case entities.TaskStBilled:
			if errToState = tss.Task.TaskState.SetTaskStateToBilled(); errToState != nil {
				return errToState
			}
		case entities.TaskStPending:
			if errToState = tss.Task.TaskState.SetTaskStateToPending(); errToState != nil {
				return errToState
			}
		case entities.TaskStClaimed:
			if errToState = tss.Task.TaskState.SetTaskStateToClaimed(); errToState != nil {
				return errToState
			}
		case entities.TaskStArchived:
			if errToState = tss.Task.TaskState.SetTaskStateToArchived(); errToState != nil {
				return errToState
			}

		}
		// save to the db
		if errUpdate := taskRepo.Update(*tss.Task); errUpdate != nil {
			return errUpdate
		}
		th := InitNewTaskHistoricization(tss.Env)
		if errHist := th.Execute(); errHist != nil {
			return errHist
		}
	*/
	return nil
}
