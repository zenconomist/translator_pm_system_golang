package entities

import "fmt"

type IPrepareState interface {
	SetToOpen() error
	SetToUnderPrepare() error
	SetToOk() error
}

type PrepareStateConst int

const (
	PrepStInvalid PrepareStateConst = iota
	PrepStOpen
	PrepStUnderPrepare
	PrepStOk
)

func (psc PrepareStateConst) String() string {
	switch psc {
	case PrepStInvalid:
		return "Invalid"
	case PrepStOpen:
		return "Open"
	case PrepStUnderPrepare:
		return "UnderPrepare"
	case PrepStOk:
		return "Ok"
	default:
		return ""
	}
}

func (psc PrepareStateConst) UintToState(state uint) IPrepareState {
	switch state {
	case 0:
		return &PrepareStateInvalid{}
	case 1:
		return &PrepareStateOpen{}
	case 2:
		return &PrepareStateUnderPrepare{}
	case 3:
		return &PrepareStateOk{}
	default:
		return &PrepareStateInvalid{}
	}
}

func (psc PrepareStateConst) StringToState(state string) (PrepareStateConst, error) {
	switch state {
	case "Open":
		return PrepStOpen, nil
	case "UnderPrepare":
		return PrepStUnderPrepare, nil
	case "Ok":
		return PrepStOk, nil
	default:
		return PrepStInvalid, fmt.Errorf("prepare state provided is invalid")
	}
}

var allPrepStates = []PrepareStateConst{PrepStInvalid, PrepStOpen, PrepStUnderPrepare, PrepStOk}

func (ps *PrepareStateConst) Contains(toBeState int) bool {
	for _, s := range allPrepStates {
		if PrepareStateConst(toBeState) == s {
			return true
		}
	}
	return false
}

// ---------- State Objects with implementations ------------ \\

// --------INVALID------- \\

type PrepareStateInvalid struct{}

func (ps *PrepareStateInvalid) SetToOpen() error {
	return fmt.Errorf("state is invalid, can't do any action from here")
}

func (ps *PrepareStateInvalid) SetToUnderPrepare() error {
	return fmt.Errorf("state is invalid, can't do any action from here")
}

func (ps *PrepareStateInvalid) SetToOk() error {
	return fmt.Errorf("state is invalid, can't do any action from here")
}

// --------OPEN-------- \\

type PrepareStateOpen struct {
	Task            *Task
	PrepState       PrepareStateConst
	FormerPrepState PrepareStateConst
}

func (ps *PrepareStateOpen) SetToOpen() error {
	return fmt.Errorf("cannot reset the state to itself")
}

func (ps *PrepareStateOpen) SetToUnderPrepare() error {
	ps.FormerPrepState = PrepStOpen
	ps.PrepState = PrepStUnderPrepare
	ps.Task.PrepareState = ps
	return nil
}

func (ps *PrepareStateOpen) SetToOk() error {
	return fmt.Errorf("cannot set task's state to Ok directly from Open state")
}

// --------UNDERPREPARE-------- \\

type PrepareStateUnderPrepare struct {
	Task            *Task
	PrepState       PrepareStateConst
	FormerPrepState PrepareStateConst
}

func (ps *PrepareStateUnderPrepare) SetToOpen() error {
	return fmt.Errorf("can not reset to open prepare state")
}

func (ps *PrepareStateUnderPrepare) SetToUnderPrepare() error {
	return fmt.Errorf("cannot reset the state to itself")
}

func (ps *PrepareStateUnderPrepare) SetToOk() error {
	return nil
}

// --------OK-------- \\

type PrepareStateOk struct {
	Task            *Task
	PrepState       PrepareStateConst
	FormerPrepState PrepareStateConst
}

func (ps *PrepareStateOk) SetToOpen() error {
	return nil
}

func (ps *PrepareStateOk) SetToUnderPrepare() error {
	return nil
}

func (ps *PrepareStateOk) SetToOk() error {
	return fmt.Errorf("cannot reset the state to itself")
}
