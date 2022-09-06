package services

import (
	"dto"
	"entities"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

// ----------------CheckProjectStateAction----------------\\

func InitCheckProjectStatesAction[E entities.Entity, D dto.DTO[E]](entity E, dto D, db *gorm.DB) *CheckProjectStatesAction {
	var projectID uint
	switch reflect.TypeOf(entity).String() {
	case "entities.Project":
		projectID = entity.GiveID()
	case "entities.Task":
		taskID := entity.GiveID()
		var task entities.Task
		repo := entities.NewRepository(db, task)
		foundTask, errRepo := repo.FindByID(taskID)
		if errRepo != nil {
			return &CheckProjectStatesAction{}
		}
		projectID = foundTask.ProjectID
	}

	return &CheckProjectStatesAction{ProjectID: projectID}
}

type CheckProjectStatesAction struct {
	ProjectID uint `json:"project_id"`
}

// CheckProjectStatesAction Execute shall check a project's state
// depending on it's task
func (a *CheckProjectStatesAction) Execute() error {
	// get current project state
	fmt.Println("checkProjectStatesAction executed")
	// if above requirements are met, project state should be updated accordingly
	return nil
}
