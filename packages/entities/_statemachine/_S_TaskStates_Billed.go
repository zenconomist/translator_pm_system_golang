package entities

import "fmt"

type TaskStateBilled struct {
	Task *Task
	Tsc  TaskStateConst
	Ftsc TaskStateConst // Former TaskStateConst
}

func (ts *TaskStateBilled) GiveName() string {
	return ts.Tsc.String()
}
func (ts *TaskStateBilled) GiveTaskStateUint() uint {
	return uint(ts.Tsc)
}
func (ts *TaskStateBilled) CheckIfStatesAreValid(toBeStates []int) error {
	for _, i := range toBeStates {
		if !ts.Tsc.Contains(i) {
			return fmt.Errorf("provided states are not valid")
		}
	}

	return nil
}

func (ts *TaskStateBilled) SetStateParams(task *Task, taskStateConst TaskStateConst) {
	ts.Task = task
	ts.Ftsc = ts.Tsc // historicizing the state
	ts.Tsc = taskStateConst
}

// ====================SetStates================ \\

func (ts *TaskStateBilled) SetTaskStateToOpen() error {
	return nil
}

func (ts *TaskStateBilled) SetTaskStateToOffered() error {
	return nil
}

func (ts *TaskStateBilled) SetTaskStateToInProgress() error {
	return nil
}

func (ts *TaskStateBilled) SetTaskStateToReady() error {
	return nil
}

func (ts *TaskStateBilled) SetTaskStateToDelivered() error {
	return nil
}

func (ts *TaskStateBilled) SetTaskStateToBilled() error {
	return nil
}

func (ts *TaskStateBilled) SetTaskStateToPending() error {
	return nil
}

func (ts *TaskStateBilled) SetTaskStateToClaimed() error {
	return nil
}

func (ts *TaskStateBilled) SetTaskStateToArchived() error {
	return nil
}
