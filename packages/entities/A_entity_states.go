package entities

import "fmt"

// ==================STATES===================\\

type IState interface {
	GiveName() string
	CheckIfStatesAreValid([]int) error
}

type State interface {
	EmptyState // |
	// TaskStateOpen | TaskStateOffered |
	// 	TaskStatePending | TaskStateBilled | TaskStateArchived |
	// 	TaskStateDelivered | TaskStateInProgress | TaskStateReady |

	BatchState | BatchTimeState | ClientOfferState |
		ClientOfferTimeState |
		ProjectState | ProjectTimeState | SupplierTimeState |
		TaskOfferedState | TaskTimeState
	GiveName() string
	CheckIfStatesAreValid([]int) error
}

type EmptyState struct{}

func (e *EmptyState) GiveName() string {
	return ""
}

func (e *EmptyState) CheckIfStatesAreValid(stateNums []int) error {
	return fmt.Errorf("empty states don't have any valid states")
}
