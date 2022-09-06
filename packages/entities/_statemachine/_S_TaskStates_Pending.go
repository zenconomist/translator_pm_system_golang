package entities

import "fmt"

type TaskStatePending struct {
	Task *Task
	Tsc  TaskStateConst
	Ftsc TaskStateConst // Former TaskStateConst
}

func (ts *TaskStatePending) GiveName() string {
	return ts.Tsc.String()
}
func (ts *TaskStatePending) GiveTaskStateUint() uint {
	return uint(ts.Tsc)
}
func (ts *TaskStatePending) CheckIfStatesAreValid(toBeStates []int) error {
	for _, i := range toBeStates {
		if !ts.Tsc.Contains(i) {
			return fmt.Errorf("provided states are not valid")
		}
	}
	return nil
}

func (ts *TaskStatePending) SetStateParams(task *Task, taskStateConst TaskStateConst) {
	ts.Task = task
	ts.Ftsc = ts.Tsc // historicizing the state
	ts.Tsc = taskStateConst
}

// ====================SetStates================ \\

// pm can set state to open from pending if it was the previous state
func (ts *TaskStatePending) SetTaskStateToOpen() error {
	ts.Tsc = TaskStOpen
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateOpen{Task: ts.Task}
	return nil
}

// pm can set state to offered from pending if it was the previous state
func (ts *TaskStatePending) SetTaskStateToOffered() error {
	ts.Tsc = TaskStOffered
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateOffered{Task: ts.Task}
	return nil
}

// pm can set state to inprogress from pending if it was the previous state
func (ts *TaskStatePending) SetTaskStateToInProgress() error {
	ts.Tsc = TaskStInProgress
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateInProgress{Task: ts.Task}
	return nil
}

// pm can set state to ready from pending if it was the previous state
func (ts *TaskStatePending) SetTaskStateToReady() error {
	ts.Tsc = TaskStReady
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateReady{Task: ts.Task}
	return nil
}

// pm can set state to delivered from pending if it was the previous state
func (ts *TaskStatePending) SetTaskStateToDelivered() error {
	ts.Tsc = TaskStDelivered
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateDelivered{Task: ts.Task}
	return nil
}

// pm can set state to billed from pending if it was the previous state
func (ts *TaskStatePending) SetTaskStateToBilled() error {
	ts.Tsc = TaskStBilled
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateBilled{Task: ts.Task}
	return nil
}

func (ts *TaskStatePending) SetTaskStateToPending() error {
	return fmt.Errorf("cannot reset the state to itself")
}

func (ts *TaskStatePending) SetTaskStateToClaimed() error {
	return nil
}

func (ts *TaskStatePending) SetTaskStateToArchived() error {
	ts.Tsc = TaskStArchived
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateArchived{Task: ts.Task}
	return nil
}
