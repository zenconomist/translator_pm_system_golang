package services

import (
	"dto"
	"entities"
)

// aim is to be able to query with associations => so projects with tasks.

type CustomService struct {
	CustomRepo *entities.CustomRepo
}

func NewCustomService(customRepo *entities.CustomRepo) *CustomService {
	return &CustomService{
		CustomRepo: customRepo,
	}
}

func (cs *CustomService) GetAllProjectsAndTasks() ([]dto.ProjectResponseDTO, error) {

	var dtos []dto.ProjectResponseDTO
	var projectDto dto.ProjectResponseDTO
	db := cs.CustomRepo.Db
	var projects []entities.Project
	var tasks []entities.Task
	var taskDto dto.TaskResponseDTO
	if result := db.Find(&projects); result.Error != nil {
		return dtos, result.Error
	}
	for _, p := range projects {
		db.Model(&p).Association("Tasks").Find(&tasks)
		projectDto.MapFromEntity(cs.CustomRepo.Db, p)
		for _, t := range tasks {
			taskDto.MapFromEntity(cs.CustomRepo.Db, t)
			projectDto.Tasks = append(projectDto.Tasks, taskDto)
		}
		dtos = append(dtos, projectDto)
	}

	return dtos, nil
}

func (cs *CustomService) GetAllActiveProjectsAndTasks() ([]dto.ProjectResponseDTO, error) {

	var dtos []dto.ProjectResponseDTO
	var projectDto dto.ProjectResponseDTO
	var project entities.Project

	var taskDto dto.TaskResponseDTO
	repo := entities.NewRepository(cs.CustomRepo.Db, project)
	projects, errFind := repo.FindAll()
	if errFind != nil {
		return dtos, errFind
	}
	for _, p := range projects {
		tasks, errTasks := cs.CustomRepo.FindAllAssociatedTask(p)
		if errTasks != nil {
			return dtos, errTasks
		}
		projectDto.MapFromEntity(cs.CustomRepo.Db, p)
		for _, t := range tasks {
			taskDto.MapFromEntity(cs.CustomRepo.Db, t)
			projectDto.Tasks = append(projectDto.Tasks, taskDto)
		}
		dtos = append(dtos, projectDto)
	}

	return dtos, nil
}
