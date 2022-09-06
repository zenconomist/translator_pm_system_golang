package entities

import "fmt"

/*
	The default process is that a task goes from open state
		to offered with the offering action, and then it goes
		either back to open, or to in progress, when a supplier
		accepts the task.
*/

type TaskStateOffered struct {
	Task *Task
	Tsc  TaskStateConst
	Ftsc TaskStateConst // Former TaskStateConst
}

func (ts *TaskStateOffered) GiveName() string {
	return ts.Tsc.String()
}
func (ts *TaskStateOffered) GiveTaskStateUint() uint {
	return uint(ts.Tsc)
}

func (ts *TaskStateOffered) CheckIfStatesAreValid(toBeStates []int) error {
	for _, i := range toBeStates {
		if !ts.Tsc.Contains(i) {
			return fmt.Errorf("provided states are not valid")
		}
	}
	return nil
}

func (ts *TaskStateOffered) SetStateParams(task *Task, taskStateConst TaskStateConst) {
	ts.Task = task
	ts.Ftsc = ts.Tsc // historicizing the state
	ts.Tsc = taskStateConst
}

// ====================SetStates================ \\

// if the supplier declines the offering, and there are no
// more suppliers who has been offered, and no one already
// accepted the task, it will be reset to open
func (ts *TaskStateOffered) SetTaskStateToOpen() error {
	// validations
	ts.Tsc = TaskStOpen
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateOpen{Task: ts.Task}

	return nil
}

func (ts *TaskStateOffered) SetTaskStateToOffered() error {
	return fmt.Errorf("cannot reset the state to itself")
}

// this is the standard process -> when a supplier accepts
// the task, it's state is set to in progress
func (ts *TaskStateOffered) SetTaskStateToInProgress() error {
	// validations
	ts.Tsc = TaskStInProgress
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateInProgress{Task: ts.Task}
	return nil
}

func (ts *TaskStateOffered) SetTaskStateToReady() error {
	return fmt.Errorf("cannot set the offered task directly to ready")
}

func (ts *TaskStateOffered) SetTaskStateToDelivered() error {
	return fmt.Errorf("cannot set the offered task directly to delivered")
}

func (ts *TaskStateOffered) SetTaskStateToBilled() error {
	return fmt.Errorf("cannot set the offered task directly to billed")
}

// if and when there is a problem with the job, the task can be
// set to pending, and no other action can be performed until
// the questions are settled. The Supplier can pull the task
// back from pending if he/she started the pending state,
// otherwise only the pm can set it to pending and put it back
// to it's original state (attributes to store: who put the
// task to this state, and what was it's previous state?)
func (ts *TaskStateOffered) SetTaskStateToPending() error {
	ts.Tsc = TaskStPending
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStatePending{Task: ts.Task}
	return nil
}

func (ts *TaskStateOffered) SetTaskStateToClaimed() error {
	return nil
}

func (ts *TaskStateOffered) SetTaskStateToArchived() error {
	ts.Tsc = TaskStArchived
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateArchived{Task: ts.Task}
	return nil
}
