package entities

import "fmt"

type IReviewState interface {
	SetToOpen() error
	SetToUnderReview() error
	SetToOk() error
}

type ReviewStateConst int

const (
	ReviewStInvalid ReviewStateConst = iota
	ReviewStOpen
	ReviewStUnderReview
	ReviewStOk
)

func (rsc ReviewStateConst) String() string {
	switch rsc {
	case ReviewStInvalid:
		return "Invalid"
	case ReviewStOpen:
		return "Open"
	case ReviewStUnderReview:
		return "UnderReview"
	case ReviewStOk:
		return "Ok"
	default:
		return ""
	}
}

func (rsc ReviewStateConst) UintToState(state uint) IReviewState {
	switch state {
	case 0:
		return &ReviewStateInvalid{}
	case 1:
		return &ReviewStateOpen{}
	case 2:
		return &ReviewStateUnderReview{}
	case 3:
		return &ReviewStateOk{}
	default:
		return &ReviewStateInvalid{}
	}
}

func (rsc ReviewStateConst) StringToState(state string) (ReviewStateConst, error) {
	switch state {
	case "Open":
		return ReviewStOpen, nil
	case "UnderReview":
		return ReviewStUnderReview, nil
	case "Ok":
		return ReviewStOk, nil
	default:
		return ReviewStInvalid, fmt.Errorf("prepare state provided is invalid")
	}
}

var allReviewStates = []ReviewStateConst{ReviewStInvalid, ReviewStOpen, ReviewStUnderReview, ReviewStOk}

func (ps *ReviewStateConst) Contains(toBeState int) bool {
	for _, s := range allReviewStates {
		if ReviewStateConst(toBeState) == s {
			return true
		}
	}
	return false
}

// ---------- State Objects with implementations ------------ \\

// --------INVALID------- \\

type ReviewStateInvalid struct{}

func (ps *ReviewStateInvalid) SetToOpen() error {
	return fmt.Errorf("state is invalid, can't do any action from here")
}

func (ps *ReviewStateInvalid) SetToUnderReview() error {
	return fmt.Errorf("state is invalid, can't do any action from here")
}

func (ps *ReviewStateInvalid) SetToOk() error {
	return fmt.Errorf("state is invalid, can't do any action from here")
}

// --------OPEN-------- \\

type ReviewStateOpen struct {
	Task              *Task
	ReviewState       ReviewStateConst
	FormerReviewState ReviewStateConst
}

func (ps *ReviewStateOpen) SetToOpen() error {
	return fmt.Errorf("cannot reset the state to itself")
}

func (rs *ReviewStateOpen) SetToUnderReview() error {
	rs.FormerReviewState = ReviewStOpen
	rs.ReviewState = ReviewStUnderReview
	rs.Task.ReviewState = rs
	return nil
}

func (ps *ReviewStateOpen) SetToOk() error {
	return fmt.Errorf("cannot set task's state to Ok directly from Open state")
}

// --------UNDERPREPARE-------- \\

type ReviewStateUnderReview struct {
	Task              *Task
	ReviewState       ReviewStateConst
	FormerReviewState ReviewStateConst
}

func (rs *ReviewStateUnderReview) SetToOpen() error {
	return fmt.Errorf("can not reset to open prepare state")
}

func (rs *ReviewStateUnderReview) SetToUnderReview() error {
	return fmt.Errorf("cannot reset the state to itself")
}

func (rs *ReviewStateUnderReview) SetToOk() error {
	return nil
}

// --------OK-------- \\

type ReviewStateOk struct {
	Task              *Task
	ReviewState       ReviewStateConst
	FormerReviewState ReviewStateConst
}

func (rs *ReviewStateOk) SetToOpen() error {
	return nil
}

func (rs *ReviewStateOk) SetToUnderReview() error {
	return nil
}

func (rs *ReviewStateOk) SetToOk() error {
	return fmt.Errorf("cannot reset the state to itself")
}
