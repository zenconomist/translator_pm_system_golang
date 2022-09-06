package entities

import "fmt"

type TaskStateInvalid struct {
	Task *Task
	Tsc  TaskStateConst
	Ftsc TaskStateConst // Former TaskStateConst
}

func (ts *TaskStateInvalid) GiveName() string {
	return ts.Tsc.String()
}
func (ts *TaskStateInvalid) GiveTaskStateUint() uint {
	return uint(ts.Tsc)
}
func (ts *TaskStateInvalid) CheckIfStatesAreValid(toBeStates []int) error {
	for _, i := range toBeStates {
		if !ts.Tsc.Contains(i) {
			return fmt.Errorf("provided states are not valid")
		}
	}

	return nil
}

func (ts *TaskStateInvalid) SetStateParams(task *Task, taskStateConst TaskStateConst) {
	ts.Task = task
	ts.Ftsc = ts.Tsc // historicizing the state
	ts.Tsc = taskStateConst
}

// ====================SetStates================ \\

func (ts *TaskStateInvalid) SetTaskStateToOpen() error {
	return nil
}

func (ts *TaskStateInvalid) SetTaskStateToOffered() error {
	return nil
}

func (ts *TaskStateInvalid) SetTaskStateToInProgress() error {
	return nil
}

func (ts *TaskStateInvalid) SetTaskStateToReady() error {
	return nil
}

func (ts *TaskStateInvalid) SetTaskStateToDelivered() error {
	return nil
}

func (ts *TaskStateInvalid) SetTaskStateToBilled() error {
	return nil
}

func (ts *TaskStateInvalid) SetTaskStateToPending() error {
	return nil
}

func (ts *TaskStateInvalid) SetTaskStateToClaimed() error {
	return nil
}

func (ts *TaskStateInvalid) SetTaskStateToArchived() error {
	return nil
}
