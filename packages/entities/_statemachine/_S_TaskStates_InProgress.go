package entities

import "fmt"

/*
	The default process is that a task goes from open state
		to offered with the offering action. When a supplier
		accepts the offering, the state goes into in progress
		state. From in progress it can go back to open, it can
		go to ready, to pending, or archived.
*/

type TaskStateInProgress struct {
	Task *Task
	Tsc  TaskStateConst
	Ftsc TaskStateConst // Former TaskStateConst
}

func (ts *TaskStateInProgress) GiveName() string {
	return ts.Tsc.String()
}

func (ts *TaskStateInProgress) GiveTaskStateUint() uint {
	return uint(ts.Tsc)
}

func (ts *TaskStateInProgress) CheckIfStatesAreValid(toBeStates []int) error {
	for _, i := range toBeStates {
		if !ts.Tsc.Contains(i) {
			return fmt.Errorf("provided states are not valid")
		}
	}
	return nil
}

func (ts *TaskStateInProgress) SetStateParams(task *Task, taskStateConst TaskStateConst) {
	ts.Task = task
	ts.Ftsc = ts.Tsc // historicizing the state
	ts.Tsc = taskStateConst
}

// ====================SetStates================ \\

// if the supplier somehow gives the job back, or the task
// is retreived by the pm from the supplier, it can be set
// back to open state
func (ts *TaskStateInProgress) SetTaskStateToOpen() error {
	ts.Tsc = TaskStOpen
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateOpen{Task: ts.Task}
	return nil
}

func (ts *TaskStateInProgress) SetTaskStateToOffered() error {
	return fmt.Errorf("cannot set the task's state to offered, since all other supplier's have been informed that the task was accepted by someone else")
}

func (ts *TaskStateInProgress) SetTaskStateToInProgress() error {
	return fmt.Errorf("cannot reset the state to itself")
}

// this is the normal process -> from in progress it goes to ready
func (ts *TaskStateInProgress) SetTaskStateToReady() error {
	ts.Tsc = TaskStReady
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateReady{Task: ts.Task}
	return nil
}

func (ts *TaskStateInProgress) SetTaskStateToDelivered() error {
	return fmt.Errorf("cannot set the task's state to delivered from in progress directly")
}

func (ts *TaskStateInProgress) SetTaskStateToBilled() error {
	return fmt.Errorf("cannot set the task's state to billed from in progress directly")
}

// the pm or the supplier can set from open to pending
// if there is a question
// towards the client which needs to be settled first.
func (ts *TaskStateInProgress) SetTaskStateToPending() error {
	ts.Tsc = TaskStPending
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStatePending{Task: ts.Task}
	return nil
}

func (ts *TaskStateInProgress) SetTaskStateToClaimed() error {
	return nil
}

func (ts *TaskStateInProgress) SetTaskStateToArchived() error {
	ts.Tsc = TaskStArchived
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateArchived{Task: ts.Task}
	return nil
}
