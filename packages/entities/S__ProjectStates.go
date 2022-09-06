package entities

type IProjectState interface {
	GiveProjectStateName() string
}

type ProjectState struct {
	Name  string
	Order uint
}

func (ps *ProjectState) GiveProjectStateName() string {
	return ps.Name
}
