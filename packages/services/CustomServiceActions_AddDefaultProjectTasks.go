package services

import (
	"dto"
	"entities"
	"environment"
	"fmt"
	"reflect"
	"time"
)

// ----------------DefaultProjectTasks----------------\\

func InitNewDefaultProjectTasks[E entities.Entity, D dto.DTO[E]](entity E, dto D, env environment.Environment, newID uint) *InitDefaultProjectTasks {
	var idp InitDefaultProjectTasks
	if reflect.TypeOf(entity).String() != "entities.Project" {
		defer env.GiveLogger().LogError(fmt.Errorf("by InitDefaultProject the given entity is not a Project entity"))
		return &idp
	}
	return &InitDefaultProjectTasks{
		Env:       env,
		ProjectID: newID,
	}
}

type InitDefaultProjectTasks struct {
	Env       environment.Environment
	ProjectID uint
}

func (ids *InitDefaultProjectTasks) Execute() error {
	// get default taskconfigs
	var tcs []entities.TaskConfig
	db := ids.Env.GiveDbHandler().PassConnection()
	result := db.Where("add_to_default = ?", true).Find(&tcs)
	if result.Error != nil {
		return result.Error
	}
	var tsch entities.TaskStateConfigHead
	resTsch := db.Where("? between active_from and active_to", time.Now()).Last(&tsch)
	if resTsch.Error != nil {
		return resTsch.Error
	}
	var ts entities.TaskState
	resTs := db.Where("task_state_config_head_id = ? AND task_state_code = ?", tsch.ID, "op").First(&ts)
	if resTs.Error != nil {
		return resTs.Error
	}

	taskRepo := entities.NewRepository(db, entities.Task{})
	service := NewService(taskRepo, ids.Env, &dto.TaskRequestDTO{}, entities.Task{}, &entities.TaskHistory{})
	// initiate Tasks from TaskConfigs and persist them for the given ProjectID
	for _, tc := range tcs {
		task := entities.Task{
			ProjectID:             ids.ProjectID,
			BatchID:               0,
			OrderWithinBatch:      0,
			TaskType:              tc.TaskType,
			ProjectManager:        1, // TODO: replace with user!
			TaskStateID:           ts.ID,
			TaskStateConfigHeadID: tsch.ID,
			TaskTimeStateID:       0,
			SourceLang:            tc.SourceLang,
			TargetLang:            tc.TargetLang,
			PrepDisabled:          tc.PrepDisabled,
			PrepareStateID:        uint(entities.PrepStOpen),
			PrepBillable:          tc.PrepBillable,
			PreparerID:            tc.PreparerID,
			ReviewDisabled:        tc.ReviewDisabled,
			ReviewStateID:         uint(entities.ReviewStOpen),
			ReviewBillable:        tc.ReviewBillable,
			ReviewerID:            tc.ReviewerID,
			TaskCustomerProps:     entities.TaskCustomerProps{
				// CustomerDueDate: , // TODO: create function to convert from TDDTime into Due Date
			},
		}

		ID, errCreate := taskRepo.Create(task)
		if errCreate != nil {
			return errCreate
		}

		// historicization
		if errHistory := service.createHistory(task, ID); errHistory != nil {
			return errHistory
		}

	}

	fmt.Println("InitDefaultProjectTasks executed")
	return nil
}
