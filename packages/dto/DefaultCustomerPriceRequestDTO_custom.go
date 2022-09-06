package dto

import (
	"encoding/json"
	"entities"
	"fmt"
	"logging"

	"gorm.io/gorm"
)

func (data *DefaultCustomerPriceRequestDTO) CustomDTOTransformations(db *gorm.DB) error {
	//implementation
	return nil
}

func (data *DefaultCustomerPriceRequestDTO) CustomValidations(db *gorm.DB, logger logging.Logger) error {
	// implement validations
	return nil
}

func (data *DefaultCustomerPriceRequestDTO) GiveID() (uint, error) {
	if data.ID == 0 {
		return data.ID, fmt.Errorf("the dto doesn't have any valid ID")
	}
	return data.ID, nil
}

func (data *DefaultCustomerPriceRequestDTO) MapToEntity(db *gorm.DB, entity *entities.DefaultCustomerPrice) (*entities.DefaultCustomerPrice, error) {
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

func (data *DefaultCustomerPriceRequestDTO) MapFromEntity(db *gorm.DB, entity entities.DefaultCustomerPrice) error {
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

func (data *DefaultCustomerPriceRequestDTO) customMappingFromEntity(entity entities.DefaultCustomerPrice, db *gorm.DB) error {
	return nil
}

func (data *DefaultCustomerPriceRequestDTO) customMappingToEntity(db *gorm.DB) error {
	return nil
}
