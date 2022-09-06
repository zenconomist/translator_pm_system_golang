package dto

import (
	"encoding/json"
	"entities"
	"fmt"
	"logging"

	"gorm.io/gorm"
)

func (data *TaskOfferedRequestDTO) CustomDTOTransformations(db *gorm.DB) error {
	invalidTosc := entities.TaskOfferedStateInvalid
	state, errToState := invalidTosc.StringToState(data.TaskOfferedState)
	if errToState != nil {
		return errToState
	}
	data.TaskOfferedStateID = uint(state)
	return nil
}

func (data *TaskOfferedRequestDTO) CustomValidations(db *gorm.DB, logger logging.Logger) error {
	// check if there is already a record for the given
	// task id and supplier id, and if there is, respond with error
	cr := entities.NewCustomRepo(db, entities.TaskOffered{})
	taskOfferExists := cr.GetTaskOfferToValidate(data.TaskID, data.SupplierID)
	if taskOfferExists {
		return fmt.Errorf("a task offer already exists for the same task and supplier")
	}

	// check if TaskID is valid, active, and in proper state
	return nil
}

func (data *TaskOfferedRequestDTO) GiveID() (uint, error) {
	if data.ID == 0 {
		return data.ID, fmt.Errorf("the dto doesn't have any valid ID")
	}
	return data.ID, nil
}

func (data *TaskOfferedRequestDTO) MapToEntity(db *gorm.DB, entity *entities.TaskOffered) (*entities.TaskOffered, error) {
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

func (data *TaskOfferedRequestDTO) MapFromEntity(db *gorm.DB, entity entities.TaskOffered) error {
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

func (data *TaskOfferedRequestDTO) customMappingFromEntity(taskOffered entities.TaskOffered, db *gorm.DB) error {
	return nil
}

func (data *TaskOfferedRequestDTO) customMappingToEntity(db *gorm.DB) error {
	return nil
}
