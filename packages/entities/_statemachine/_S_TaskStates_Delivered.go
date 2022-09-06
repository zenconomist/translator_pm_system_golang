package entities

import "fmt"

/*
	The default process is that a task goes from open state
		to offered with the offering action, when accepted
		it goes to inprogress, then to ready, and from
		ready it goes to delivered, and from delivered
		it can go to billed or archived.
*/

type TaskStateDelivered struct {
	Task *Task
	Tsc  TaskStateConst
	Ftsc TaskStateConst // Former TaskStateConst
}

func (ts *TaskStateDelivered) GiveName() string {
	return ts.Tsc.String()
}

func (ts *TaskStateDelivered) GiveTaskStateUint() uint {
	return uint(ts.Tsc)
}

func (ts *TaskStateDelivered) CheckIfStatesAreValid(toBeStates []int) error {
	for _, i := range toBeStates {
		if !ts.Tsc.Contains(i) {
			return fmt.Errorf("provided states are not valid")
		}
	}
	return nil
}

func (ts *TaskStateDelivered) SetStateParams(task *Task, taskStateConst TaskStateConst) {
	ts.Task = task
	ts.Ftsc = ts.Tsc // historicizing the state
	ts.Tsc = taskStateConst
}

// ====================SetStates================ \\

func (ts *TaskStateDelivered) SetTaskStateToOpen() error {
	return fmt.Errorf("cannot set the delivered task to open")
}

func (ts *TaskStateDelivered) SetTaskStateToOffered() error {
	return fmt.Errorf("cannot set the delivered task to offered")
}

// if the delivered task was sent back (claimed) by the customer
// it can be set back to InProgress. The Reviewer can set
// it also back to InProgress
func (ts *TaskStateDelivered) SetTaskStateToInProgress() error {
	ts.Tsc = TaskStInProgress
	ts.Task.TaskStateID = uint(ts.Tsc)
	ts.Task.TaskState = &TaskStateInProgress{Task: ts.Task}
	return nil
}

func (ts *TaskStateDelivered) SetTaskStateToReady() error {
	return nil
}

func (ts *TaskStateDelivered) SetTaskStateToDelivered() error {
	return fmt.Errorf("cannot reset the state to itself")
}

func (ts *TaskStateDelivered) SetTaskStateToBilled() error {
	return nil
}

func (ts *TaskStateDelivered) SetTaskStateToPending() error {
	return nil
}

func (ts *TaskStateDelivered) SetTaskStateToClaimed() error {
	return nil
}

func (ts *TaskStateDelivered) SetTaskStateToArchived() error {
	return nil
}
