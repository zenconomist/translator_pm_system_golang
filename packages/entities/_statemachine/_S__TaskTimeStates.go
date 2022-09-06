package entities

type ITaskTimeState interface {
	GiveTaskTimeStateName() string
}

/*
type TaskTimeStateConst int

const (
	TaskTimeStOk  TaskTimeStateConst = iota
	TaskTimeStUpd                    //upcoming deadline
	TaskTimeStOverdue
)

type TaskTimeState struct {
	Name  string
	Order uint
}

func (tts *TaskTimeState) GiveTaskTimeStateName() string {
	return tts.Name
}

type TaskTimeStateOk struct {
	Name  string
	Order uint
}

func (tts *TaskTimeStateOk) GiveTaskTimeStateName() string {
	return tts.Name
}
*/
