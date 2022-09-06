package services

import (
	"entities"
	"environment"
	"fmt"
)

type TaskOfferingStateService struct {
	To          *entities.TaskOffered
	Env         environment.Environment
	ToscDefault entities.TaskOfferedStatesConst
}

func NewTaskOfferingStateService(to *entities.TaskOffered, env environment.Environment) *TaskOfferingStateService {
	return &TaskOfferingStateService{
		To:          to,
		Env:         env,
		ToscDefault: entities.TaskOfferedStateInvalid,
	}
}

func (toss *TaskOfferingStateService) SetState(state string) error {
	toRepo := entities.NewRepository(toss.Env.GiveDbHandler().PassConnection(), entities.TaskOffered{})

	// converting the given new state into a TaskOfferedStatesConst
	validState, errValid := toss.ToscDefault.StringToState(state)
	if errValid != nil {
		return errValid
	}
	// converting the persisted TaskOfferedStateID into state object
	newToState := toss.ToscDefault.UintToState(toss.To.TaskOfferedStateID)
	toss.To.TaskOfferedState = newToState

	// set inner params of the state object (taskOffering and state const)
	toss.To.TaskOfferedState.SetStateParams(toss.To)

	// in order to be able to return errors in the switch statement
	var errToState error

	switch validState {
	case entities.TaskOfferedStateOffered:
		if errToState = toss.To.TaskOfferedState.SetToOffered(); errToState != nil {
			return errToState
		}
	case entities.TaskOfferedStateAccepted:
		// validation: if the task can be found with an "offered" state
		cr := entities.NewCustomRepo(toss.Env.GiveDbHandler().PassConnection(), entities.Task{})
		task, errFind := cr.FindOfferedTask(toss.To.TaskID)
		if errFind != nil {
			return fmt.Errorf("couldn't retreive corresponding offered Task: %v", errFind)
		}
		tss := NewTaskStateDbBasedService(task, toss.Env)

		// setting the effective state of state offering
		if errToState = toss.To.TaskOfferedState.SetToAccepted(); errToState != nil {
			return errToState
		}

		// run task state service -> if the task is accepted, it's state goes into In Progress
		tss.SetState("ip")

	case entities.TaskOfferedStateDeclined:
		if errToState = toss.To.TaskOfferedState.SetToDeclined(); errToState != nil {
			return errToState
		}
		// if there is no more offering for the current task, the task's state should be set to Open again
		cr := entities.NewCustomRepo(toss.Env.GiveDbHandler().PassConnection(), entities.TaskOffered{})
		otherOfferings := cr.FindAnyOtherTaskOfferings(toss.To.TaskID)
		if !otherOfferings {
			taskRepo := entities.NewRepository(toss.Env.GiveDbHandler().PassConnection(), entities.Task{})
			task, errFindTask := taskRepo.FindByID(toss.To.TaskID)
			if errFindTask != nil {
				return errFindTask
			}
			tss := NewTaskStateService(&task, toss.Env)
			tss.SetState("op")
		}

	}

	// persist TaskOffering changes
	if errUpdate := toRepo.Update(*toss.To); errUpdate != nil {
		return errUpdate
	}
	return nil
}
