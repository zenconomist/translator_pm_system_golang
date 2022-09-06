package dto

import (
	"entities"
	"fmt"
	"logging"

	"gorm.io/gorm"
)

type TaskStateRequestDTO struct {
	TaskID  uint   `json:"taskid"`
	ToState string `json:"to_state"`
}

func (data *TaskStateRequestDTO) CustomDTOTransformations(db *gorm.DB) error {
	//implementation
	return nil
}

func (data *TaskStateRequestDTO) CustomValidations(db *gorm.DB, logger logging.Logger) error {
	// check if ToState is valid
	// ts := entities.TaskStInvalid
	// if _, errValidState := ts.StringToState(data.ToState); errValidState != nil {
	// 	return errValidState
	// }
	return nil
}

func (data *TaskStateRequestDTO) GiveID() (uint, error) {
	if data.TaskID == 0 {
		return data.TaskID, fmt.Errorf("the dto doesn't have any valid ID")
	}
	return data.TaskID, nil
}

func (data *TaskStateRequestDTO) MapToEntity(db *gorm.DB, entity *entities.Task) (*entities.Task, error) {
	return entity, nil
}

func (data *TaskStateRequestDTO) MapFromEntity(db *gorm.DB, entity entities.Task) error {
	return nil
}

func (data *TaskStateRequestDTO) customMappingFromEntity(db *gorm.DB) error {
	return nil
}

func (data *TaskStateRequestDTO) customMappingToEntity(db *gorm.DB) error {
	return nil
}
