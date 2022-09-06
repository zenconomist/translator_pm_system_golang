package entities

import "fmt"

type ITaskState interface {
	// SetTaskStateName() // no entity level for name
	GiveName() string
	GiveTaskStateUint() uint
	SetStateParams(task *Task, taskStateConst TaskStateConst)
	CheckIfStatesAreValid([]int) error
	SetTaskStateToOpen() error
	SetTaskStateToOffered() error
	SetTaskStateToInProgress() error
	SetTaskStateToReady() error
	SetTaskStateToDelivered() error
	SetTaskStateToBilled() error
	SetTaskStateToPending() error
	SetTaskStateToClaimed() error
	SetTaskStateToArchived() error
}

// =============CONSTS=============== \\

type TaskStateConst int

const (
	TaskStInvalid TaskStateConst = iota
	TaskStOpen
	TaskStOffered
	TaskStInProgress
	TaskStReady
	TaskStDelivered
	TaskStBilled
	TaskStPending
	TaskStClaimed
	TaskStArchived
)

func (tsc TaskStateConst) String() string {
	switch tsc {
	case TaskStOpen:
		return "Open"
	case TaskStOffered:
		return "Offered"
	case TaskStInProgress:
		return "InProgress"
	case TaskStReady:
		return "Ready"
	case TaskStDelivered:
		return "Delivered"
	case TaskStBilled:
		return "Billed"
	case TaskStPending:
		return "Pending"
	case TaskStClaimed:
		return "Claimed"
	case TaskStArchived:
		return "Archived"
	case TaskStInvalid:
		return "Invalid"
	default:
		return ""
	}
}

func (tsc TaskStateConst) StringToState(state string) (TaskStateConst, error) {
	switch state {
	case "Open":
		return TaskStOpen, nil
	case "Offered":
		return TaskStOffered, nil
	case "InProgress":
		return TaskStInProgress, nil
	case "Ready":
		return TaskStReady, nil
	case "Delivered":
		return TaskStDelivered, nil
	case "Billed":
		return TaskStBilled, nil
	case "Pending":
		return TaskStPending, nil
	case "Claimed":
		return TaskStClaimed, nil
	case "Archived":
		return TaskStArchived, nil
	default:
		return TaskStInvalid, fmt.Errorf("task state provided is invalid")
	}
}

func (tsc TaskStateConst) UintToState(state uint) ITaskState {
	switch state {
	case 0:
		return &TaskStateInvalid{}
	case 1:
		return &TaskStateOpen{}
	case 2:
		return &TaskStateOffered{}
	case 3:
		return &TaskStateInProgress{}
	case 4:
		return &TaskStateReady{}
	case 5:
		return &TaskStateDelivered{}
	case 6:
		return &TaskStateBilled{}
	case 7:
		return &TaskStatePending{}
	case 8:
		return &TaskStateClaimed{}
	case 9:
		return &TaskStateArchived{}
	default:
		return &TaskStateInvalid{}
	}
}

func (tsc TaskStateConst) ToStruct() ITaskState {
	switch tsc {
	case TaskStOpen:
		return &TaskStateOpen{}
	case TaskStOffered:
		return &TaskStateOffered{}
	case TaskStInProgress:
		return &TaskStateInProgress{}
	case TaskStReady:
		return &TaskStateReady{}
	case TaskStDelivered:
		return &TaskStateDelivered{}
	case TaskStBilled:
		return &TaskStateBilled{}
	case TaskStPending:
		return &TaskStatePending{}
	case TaskStClaimed:
		return &TaskStateClaimed{}
	case TaskStArchived:
		return &TaskStateArchived{}
	case TaskStInvalid:
		return &TaskStateInvalid{}
	default:
		return &TaskStateInvalid{}
	}
}

var allTaskStates = []TaskStateConst{TaskStOpen, TaskStOffered, TaskStInProgress, TaskStPending, TaskStReady, TaskStDelivered, TaskStBilled, TaskStArchived}

func (ts *TaskStateConst) Contains(toBeState int) bool {
	for _, s := range allTaskStates {
		if TaskStateConst(toBeState) == s {
			return true
		}
	}
	return false
}
