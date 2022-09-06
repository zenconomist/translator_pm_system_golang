package services

import (
	"entities"
	"environment"
)

type TaskHistoricization struct {
	Env environment.Environment
}

func InitNewTaskHistoricization(env environment.Environment) *TaskHistoricization {
	return &TaskHistoricization{
		Env: env,
	}
}

func (th *TaskHistoricization) Execute() error {
	cr := entities.NewCustomRepo(th.Env.GiveDbHandler().PassConnection(), entities.Task{})
	if err := cr.TaskHistoricization(); err != nil {
		return err
	}
	return nil
}
