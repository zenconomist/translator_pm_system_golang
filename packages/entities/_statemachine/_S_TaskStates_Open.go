package entities

import "fmt"

/*
	The default process is that a task goes from open state
		to offered with the offering action.
*/

type TaskStateOpen struct {
	Task *Task
	Tsc  TaskStateConst // Tsc = TaskStateConst
	Ftsc TaskStateConst // Ftsc = Former TaskStateConst
}

func (ts *TaskStateOpen) GiveName() string {
	return ts.Tsc.String()
}

func (ts *TaskStateOpen) GiveTaskStateUint() uint {
	return uint(ts.Tsc)
}

func (ts *TaskStateOpen) CheckIfStatesAreValid(toBeStates []int) error {
	for _, i := range toBeStates {
		if !ts.Tsc.Contains(i) {
			return fmt.Errorf("provided states are not valid")
		}
	}
	return nil
}

func (ts *TaskStateOpen) SetStateParams(task *Task, taskStateConst TaskStateConst) {
	ts.Task = task
	ts.Tsc = taskStateConst
}

// ====================SetStates================ \\

func (ts *TaskStateOpen) SetTaskStateToOpen() error {
	return fmt.Errorf("cannot reset the state to itself")
}

func (ts *TaskStateOpen) SetTaskStateToOffered() error {
	ts.Tsc = TaskStOffered
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateOffered{}
	ts.Task.TaskState.SetStateParams(ts.Task, TaskStOffered)
	return nil
}

func (ts *TaskStateOpen) SetTaskStateToInProgress() error {
	return fmt.Errorf("cannot reset the state to in progress directly. First it has to be offered")
}

func (ts *TaskStateOpen) SetTaskStateToReady() error {
	return fmt.Errorf("cannot set the task as ready directly from open, first it has to be in progress")
}

func (ts *TaskStateOpen) SetTaskStateToDelivered() error {
	return fmt.Errorf("cannot set the task as delivered directly from open, first it has to be in progress, then set to be ready")
}

func (ts *TaskStateOpen) SetTaskStateToBilled() error {
	return fmt.Errorf("cannot set the task as billed directly from open, first it has to be in progress, then set to be ready, then set as delivered")
}

// the pm can set from open to pending if there is a question
// towards the client which needs to be settled first.
func (ts *TaskStateOpen) SetTaskStateToPending() error {
	ts.Tsc = TaskStPending
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStatePending{Task: ts.Task}
	return nil
}

func (ts *TaskStateOpen) SetTaskStateToClaimed() error {
	return nil
}

func (ts *TaskStateOpen) SetTaskStateToArchived() error {
	ts.Tsc = TaskStArchived
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateArchived{Task: ts.Task}
	return nil
}
