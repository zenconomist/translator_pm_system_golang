package dto

import (
	"entities"
	"fmt"
	"logging"

	"gorm.io/gorm"
)

type TaskOfferedStateRequestDTO struct {
	ID         uint   `json:"id"`
	TaskID     uint   `json:"task_id"`
	SupplierID uint   `json:"supplier_id"`
	ToState    string `json:"to_state"`
	StateConst entities.TaskOfferedStatesConst
}

func (data *TaskOfferedStateRequestDTO) CustomDTOTransformations(db *gorm.DB) error {
	//implementation
	ts := entities.TaskOfferedStateInvalid
	validState, errValidState := ts.StringToState(data.ToState)
	if errValidState != nil {
		return errValidState
	}
	data.StateConst = validState
	return nil
}

func (data *TaskOfferedStateRequestDTO) CustomValidations(db *gorm.DB, logger logging.Logger) error {
	// check if ToState is valid
	return nil
}

func (data *TaskOfferedStateRequestDTO) GiveID() (uint, error) {
	if data.TaskID == 0 {
		return data.TaskID, fmt.Errorf("the dto doesn't have any valid ID")
	}
	return data.TaskID, nil
}

func (data *TaskOfferedStateRequestDTO) MapToEntity(db *gorm.DB, entity *entities.Task) (*entities.Task, error) {
	return entity, nil
}

func (data *TaskOfferedStateRequestDTO) MapFromEntity(db *gorm.DB, entity entities.Task) error {
	return nil
}

func (data *TaskOfferedStateRequestDTO) customMappingFromEntity(db *gorm.DB) error {
	return nil
}

func (data *TaskOfferedStateRequestDTO) customMappingToEntity(db *gorm.DB) error {
	return nil
}
