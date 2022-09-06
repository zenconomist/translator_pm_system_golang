package entities

import "fmt"

type TaskStateArchived struct {
	Task *Task
	Tsc  TaskStateConst
	Ftsc TaskStateConst // Former TaskStateConst
}

func (ts *TaskStateArchived) GiveName() string {
	return ts.Tsc.String()
}
func (ts *TaskStateArchived) GiveTaskStateUint() uint {
	return uint(ts.Tsc)
}
func (ts *TaskStateArchived) CheckIfStatesAreValid(toBeStates []int) error {
	for _, i := range toBeStates {
		if !ts.Tsc.Contains(i) {
			return fmt.Errorf("provided states are not valid")
		}
	}
	return nil
}

func (ts *TaskStateArchived) SetStateParams(task *Task, taskStateConst TaskStateConst) {
	ts.Task = task
	ts.Ftsc = ts.Tsc // historicizing the state
	ts.Tsc = taskStateConst
}

// ====================SetStates================ \\

func (ts *TaskStateArchived) SetTaskStateToOpen() error {
	ts.Tsc = TaskStOpen
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateOpen{Task: ts.Task}
	return nil
}

func (ts *TaskStateArchived) SetTaskStateToOffered() error {
	ts.Tsc = TaskStOffered
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateOffered{Task: ts.Task}
	return nil
}

func (ts *TaskStateArchived) SetTaskStateToInProgress() error {
	return nil
}

func (ts *TaskStateArchived) SetTaskStateToReady() error {
	return nil
}

func (ts *TaskStateArchived) SetTaskStateToDelivered() error {
	return nil
}

func (ts *TaskStateArchived) SetTaskStateToBilled() error {
	return nil
}

func (ts *TaskStateArchived) SetTaskStateToPending() error {
	return nil
}

func (ts *TaskStateArchived) SetTaskStateToClaimed() error {
	return nil
}

func (ts *TaskStateArchived) SetTaskStateToArchived() error {
	return nil
}
