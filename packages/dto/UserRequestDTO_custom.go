package dto

import (
	"encoding/json"
	"entities"
	"fmt"
	"globalconstants"
	"logging"

	"gorm.io/gorm"
)

func (data *UserRequestDTO) CustomDTOTransformations(db *gorm.DB) error {
	// currency transform
	cur, errCurrency := globalconstants.StringToCurrency(data.BillingCurrencyName)
	if errCurrency != nil {
		return fmt.Errorf("couldn't convert currency into valid currency, error: %v", errCurrency)
	}
	data.BillingCurrency = cur
	return nil
}

func (data *UserRequestDTO) CustomValidations(db *gorm.DB, logger logging.Logger) error {
	// implement validations
	return nil
}

func (data *UserRequestDTO) GiveID() (uint, error) {
	if data.ID == 0 {
		return data.ID, fmt.Errorf("the dto doesn't have any valid ID")
	}
	return data.ID, nil
}

func (data *UserRequestDTO) MapToEntity(db *gorm.DB, entity *entities.User) (*entities.User, error) {
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

func (data *UserRequestDTO) MapFromEntity(db *gorm.DB, entity entities.User) error {
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

func (data *UserRequestDTO) customMappingFromEntity(user entities.User, db *gorm.DB) error {
	return nil
}

func (data *UserRequestDTO) customMappingToEntity(db *gorm.DB) error {
	return nil
}
