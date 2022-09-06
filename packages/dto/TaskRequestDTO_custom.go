package dto

import (
	"encoding/json"
	"entities"
	"fmt"
	"logging"

	"gorm.io/gorm"
)

func (data *TaskRequestDTO) CustomDTOTransformations(db *gorm.DB) error {
	// task state conversion
	// tsc := entities.TaskStInvalid
	// ts, errTs := tsc.StringToState(data.TaskStateName)
	// if errTs != nil {
	// 	return errTs
	// }
	// data.TaskStateID = uint(ts)

	return nil
}

func (data *TaskRequestDTO) CustomValidations(db *gorm.DB, logger logging.Logger) error {
	// implement validations
	return nil
}

func (data *TaskRequestDTO) GiveID() (uint, error) {
	if data.ID == 0 {
		return data.ID, fmt.Errorf("the dto doesn't have any valid ID")
	}
	return data.ID, nil
}

func (data *TaskRequestDTO) MapToEntity(db *gorm.DB, entity *entities.Task) (*entities.Task, error) {
	errorCustomMapping := data.customMappingToEntity(db)
	if errorCustomMapping != nil {
		return entity, errorCustomMapping
	}
	jsonDtoBytes, errJsonMarshal := json.Marshal(data)
	if errJsonMarshal != nil {
		return entity, errJsonMarshal
	}
	errJsonUnmarshal := json.Unmarshal(jsonDtoBytes, entity)
	if errJsonUnmarshal != nil {
		return entity, errJsonUnmarshal
	}
	return entity, nil
}

func (data *TaskRequestDTO) MapFromEntity(db *gorm.DB, entity entities.Task) error {
	jsonDataBytes, errJsonMarshal := json.Marshal(entity)
	if errJsonMarshal != nil {
		return errJsonMarshal
	}
	errJsonUnmarshal := json.Unmarshal(jsonDataBytes, data)
	if errJsonUnmarshal != nil {
		return errJsonUnmarshal
	}
	errorCustomMapping := data.customMappingFromEntity(entity, db)
	if errorCustomMapping != nil {
		return errorCustomMapping
	}
	return nil
}

func (data *TaskRequestDTO) customMappingFromEntity(task entities.Task, db *gorm.DB) error {
	return nil
}

func (data *TaskRequestDTO) customMappingToEntity(db *gorm.DB) error {
	return nil
}
