package entities

type IProjectTimeState interface {
	GiveProjectTimeStateName() string
}

type ProjectTimeState struct {
	Name  string
	Order uint
}

func (pts *ProjectTimeState) GiveProjectTimeStateName() string {
	return pts.Name
}
