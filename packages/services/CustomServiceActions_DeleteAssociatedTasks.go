package services

import (
	"dto"
	"entities"
	"environment"
)

type DeleteAssociatedTasks struct {
	Env       environment.Environment
	ProjectID uint
}

func InitDeleteAssociatedTasks(projectID uint, env environment.Environment) *DeleteAssociatedTasks {
	return &DeleteAssociatedTasks{
		ProjectID: projectID,
		Env:       env,
	}
}

func (dat *DeleteAssociatedTasks) Execute() error {
	repo := entities.NewRepository(dat.Env.GiveDbHandler().PassConnection(), entities.Project{})
	project, errGetProject := repo.FindByID(dat.ProjectID)
	if errGetProject != nil {
		return errGetProject
	}
	cr := entities.NewCustomRepo(dat.Env.GiveDbHandler().PassConnection(), project)
	tasks, errGetTasks := cr.FindAllAssociatedTask(project)
	if errGetTasks != nil {
		return errGetTasks
	}
	service := NewService(entities.NewRepository(dat.Env.GiveDbHandler().PassConnection(), entities.Task{}), dat.Env, &dto.TaskRequestDTO{}, entities.Task{}, &entities.TaskHistory{})
	for _, t := range tasks {
		if errDelete := service.DeleteItem(t.Model.ID); errDelete != nil {
			return errDelete
		}

	}
	return nil
}
