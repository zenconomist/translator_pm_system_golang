package dto

import (
	"encoding/json"
	"entities"
	"fmt"
	"logging"

	"gorm.io/gorm"
)

func (data *TaskSupplierPropsRequestDTO) CustomDTOTransformations(db *gorm.DB) error {
	//implementation
	return nil
}

func (data *TaskSupplierPropsRequestDTO) CustomValidations(db *gorm.DB, logger logging.Logger) error {
	// implement validations
	return nil
}

func (data *TaskSupplierPropsRequestDTO) GiveID() (uint, error) {
	if data.ID == 0 {
		return data.ID, fmt.Errorf("the dto doesn't have any valid ID")
	}
	return data.ID, nil
}

func (data *TaskSupplierPropsRequestDTO) MapToEntity(db *gorm.DB, entity *entities.TaskSupplierProps) (*entities.TaskSupplierProps, error) {
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

func (data *TaskSupplierPropsRequestDTO) MapFromEntity(db *gorm.DB, entity entities.TaskSupplierProps) error {
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

func (data *TaskSupplierPropsRequestDTO) customMappingFromEntity(taskSupplProps entities.TaskSupplierProps, db *gorm.DB) error {
	return nil
}

func (data *TaskSupplierPropsRequestDTO) customMappingToEntity(db *gorm.DB) error {
	return nil
}
