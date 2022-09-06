package entities

import "fmt"

/*
	The default process is that a task goes from ready state
		to delivered with the delivery action.
*/

type TaskStateReady struct {
	Task *Task
	Tsc  TaskStateConst
	Ftsc TaskStateConst // Former TaskStateConst
}

func (ts *TaskStateReady) GiveName() string {
	return ts.Tsc.String()
}

func (ts *TaskStateReady) GiveTaskStateUint() uint {
	return uint(ts.Tsc)
}

func (ts *TaskStateReady) CheckIfStatesAreValid(toBeStates []int) error {
	for _, i := range toBeStates {
		if !ts.Tsc.Contains(i) {
			return fmt.Errorf("provided states are not valid")
		}
	}

	return nil
}

func (ts *TaskStateReady) SetStateParams(task *Task, taskStateConst TaskStateConst) {
	ts.Task = task
	ts.Ftsc = ts.Tsc // historicizing the state
	ts.Tsc = taskStateConst
}

// ====================SetStates================ \\

func (ts *TaskStateReady) SetTaskStateToOpen() error {
	setStateSteps(ts, TaskStOpen)
	ts.Task.TaskState = &TaskStateOpen{Task: ts.Task}
	return nil
}

func (ts *TaskStateReady) SetTaskStateToOffered() error {
	return fmt.Errorf("cannot reset the task from ready state to offered directly")
}

func (ts *TaskStateReady) SetTaskStateToInProgress() error {
	setStateSteps(ts, TaskStInProgress)
	ts.Task.TaskState = &TaskStateInProgress{Task: ts.Task}
	return nil
}

func (ts *TaskStateReady) SetTaskStateToReady() error {
	return fmt.Errorf("cannot reset the state to itself")
}

func (ts *TaskStateReady) SetTaskStateToDelivered() error {
	setStateSteps(ts, TaskStDelivered)
	ts.Task.TaskState = &TaskStateDelivered{Task: ts.Task}
	return nil
}

func (ts *TaskStateReady) SetTaskStateToBilled() error {
	return fmt.Errorf("cannot set a ready task directly to billed")
}

// the pm can set from open to pending if there is a question
// towards the client which needs to be settled before setting
// the task to delivered state.
func (ts *TaskStateReady) SetTaskStateToPending() error {
	setStateSteps(ts, TaskStPending)
	ts.Task.TaskState = &TaskStatePending{Task: ts.Task}
	return nil
}

func (ts *TaskStateReady) SetTaskStateToClaimed() error {
	return nil
}

func (ts *TaskStateReady) SetTaskStateToArchived() error {
	setStateSteps(ts, TaskStArchived)
	ts.Task.TaskState = &TaskStateArchived{Task: ts.Task}
	return nil
}

func setStateSteps(ts *TaskStateReady, tsc TaskStateConst) {
	ts.Ftsc = ts.Tsc
	ts.Tsc = tsc
	ts.Task.TaskStateID = uint(ts.Tsc)
}
