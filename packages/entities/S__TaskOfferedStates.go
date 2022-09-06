package entities

import (
	"fmt"
)

type ITaskOfferedState interface {
	SetStateParams(*TaskOffered)
	SetToOffered() error
	SetToAccepted() error
	SetToDeclined() error
}

type TaskOfferedState struct{}

type TaskOfferedStatesConst int

const (
	TaskOfferedStateInvalid TaskOfferedStatesConst = iota
	TaskOfferedStateOffered
	TaskOfferedStateAccepted
	TaskOfferedStateDeclined
)

func (tosc TaskOfferedStatesConst) String() string {
	switch tosc {
	case TaskOfferedStateOffered:
		return "Offered"
	case TaskOfferedStateAccepted:
		return "Accepted"
	case TaskOfferedStateDeclined:
		return "Declined"
	default:
		return ""
	}
}

func (tosc TaskOfferedStatesConst) UintToState(state uint) ITaskOfferedState {
	switch state {
	case 1:
		return &TaskOfferStateOffered{}
	case 2:
		return &TaskOfferStateAccepted{}
	case 3:
		return &TaskOfferStateDeclined{}
	default:
		return &TaskOfferStateInvalid{}
	}

}

func (tosc TaskOfferedStatesConst) StringToState(state string) (TaskOfferedStatesConst, error) {
	switch state {
	case "Offered":
		return TaskOfferedStateOffered, nil
	case "Accepted":
		return TaskOfferedStateAccepted, nil
	case "Declined":
		return TaskOfferedStateDeclined, nil
	default:
		return TaskOfferedStateInvalid, fmt.Errorf("task offering state provided is invalid")
	}
}

var allTaskOfferingStates = []TaskOfferedStatesConst{TaskOfferedStateInvalid, TaskOfferedStateAccepted, TaskOfferedStateDeclined, TaskOfferedStateOffered}

func (ts *TaskOfferedStatesConst) Contains(toBeState int) bool {
	for _, s := range allTaskOfferingStates {
		if TaskOfferedStatesConst(toBeState) == s {
			return true
		}
	}
	return false
}

// ---------- State Objects with implementations ------------ \\

// --------INVALID------- \\

type TaskOfferStateInvalid struct {
	To *TaskOffered
}

func (tos *TaskOfferStateInvalid) SetStateParams(to *TaskOffered) {
	tos.To = to
}

func (tos *TaskOfferStateInvalid) SetToOffered() error {
	return fmt.Errorf("state is invalid, can't do any action from here")
}

func (tos *TaskOfferStateInvalid) SetToAccepted() error {
	return fmt.Errorf("state is invalid, can't do any action from here")
}

func (tos *TaskOfferStateInvalid) SetToDeclined() error {
	return fmt.Errorf("state is invalid, can't do any action from here")
}

// --------OFFERED-------- \\

type TaskOfferStateOffered struct {
	To *TaskOffered
}

func (tos *TaskOfferStateOffered) SetStateParams(to *TaskOffered) {
	tos.To = to
}

func (tos *TaskOfferStateOffered) SetToOffered() error {
	return fmt.Errorf("a taskoffering can not be offered again")
}

func (tos *TaskOfferStateOffered) SetToAccepted() error {
	tos.To.TaskOfferedStateID = uint(TaskOfferedStateAccepted)
	return nil
}

func (tos *TaskOfferStateOffered) SetToDeclined() error {
	tos.To.TaskOfferedStateID = uint(TaskOfferedStateDeclined)
	return nil
}

// --------ACCEPTED-------- \\

type TaskOfferStateAccepted struct {
	To *TaskOffered
}

func (tos *TaskOfferStateAccepted) SetStateParams(to *TaskOffered) {
	tos.To = to
}

func (tos *TaskOfferStateAccepted) SetToOffered() error {
	return fmt.Errorf("cannot set back to offered from accepted")
}

func (tos *TaskOfferStateAccepted) SetToAccepted() error {
	return fmt.Errorf("cannot set to itself from accepted")
}

func (tos *TaskOfferStateAccepted) SetToDeclined() error {
	return fmt.Errorf("cannot set to declined from accepted")
}

// --------DECLINED-------- \\

type TaskOfferStateDeclined struct {
	To *TaskOffered
}

func (tos *TaskOfferStateDeclined) SetStateParams(to *TaskOffered) {
	tos.To = to
}

func (tos *TaskOfferStateDeclined) SetToOffered() error {
	return fmt.Errorf("cannot set back to offered from declined")
}

func (tos *TaskOfferStateDeclined) SetToAccepted() error {
	return fmt.Errorf("cannot set back to accepted from declined")
}

func (tos *TaskOfferStateDeclined) SetToDeclined() error {
	return fmt.Errorf("cannot set state to itself")
}
