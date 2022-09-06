package entities

import "fmt"

type TaskStateClaimed struct {
	Task *Task
	Tsc  TaskStateConst
	Ftsc TaskStateConst // Former TaskStateConst
}

func (ts *TaskStateClaimed) GiveName() string {
	return ts.Tsc.String()
}
func (ts *TaskStateClaimed) GiveTaskStateUint() uint {
	return uint(ts.Tsc)
}
func (ts *TaskStateClaimed) CheckIfStatesAreValid(toBeStates []int) error {
	for _, i := range toBeStates {
		if !ts.Tsc.Contains(i) {
			return fmt.Errorf("provided states are not valid")
		}
	}

	return nil
}

func (ts *TaskStateClaimed) SetStateParams(task *Task, taskStateConst TaskStateConst) {
	ts.Task = task
	ts.Ftsc = ts.Tsc // historicizing the state
	ts.Tsc = taskStateConst
}

// ====================SetStates================ \\

func (ts *TaskStateClaimed) SetTaskStateToOpen() error {
	return nil
}

func (ts *TaskStateClaimed) SetTaskStateToOffered() error {
	return nil
}

func (ts *TaskStateClaimed) SetTaskStateToInProgress() error {
	return nil
}

func (ts *TaskStateClaimed) SetTaskStateToReady() error {
	return nil
}

func (ts *TaskStateClaimed) SetTaskStateToDelivered() error {
	return nil
}

func (ts *TaskStateClaimed) SetTaskStateToBilled() error {
	return nil
}

func (ts *TaskStateClaimed) SetTaskStateToPending() error {
	return nil
}

func (ts *TaskStateClaimed) SetTaskStateToClaimed() error {
	return nil
}

func (ts *TaskStateClaimed) SetTaskStateToArchived() error {
	return nil
}
