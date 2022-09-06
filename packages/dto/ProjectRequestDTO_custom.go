package dto

import (
	"encoding/json"
	"entities"
	"fmt"
	"globalconstants"
	"logging"

	"gorm.io/gorm"
)

func (data *ProjectRequestDTO) CustomDTOTransformations(db *gorm.DB) error {
	// currency
	cur, errCurrency := globalconstants.StringToCurrency(data.CurrencyName)
	if errCurrency != nil {
		return fmt.Errorf("couldn't convert currency into valid currency, error: %v", errCurrency)
	}
	data.Currency = cur

	// firm
	var firm entities.Firm
	result := db.Where("name = ?", data.FirmName).First(&firm)
	if result.Error != nil {
		return fmt.Errorf("couldn't find firm id for firm name: " + data.FirmName + " db error: " + result.Error.Error())
	}
	data.Firm = firm.Model.ID

	return nil
}

func (data *ProjectRequestDTO) CustomValidations(db *gorm.DB, logger logging.Logger) error {
	// implement validations
	return nil
}

func (data *ProjectRequestDTO) GiveID() (uint, error) {
	if data.ID == 0 {
		return data.ID, fmt.Errorf("the dto doesn't have any valid ID")
	}
	return data.ID, nil
}

func (data *ProjectRequestDTO) MapToEntity(db *gorm.DB, entity *entities.Project) (*entities.Project, error) {
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

func (data *ProjectRequestDTO) MapFromEntity(db *gorm.DB, entity entities.Project) error {
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

func (data *ProjectRequestDTO) customMappingFromEntity(entity entities.Project, db *gorm.DB) error {
	return nil
}

func (data *ProjectRequestDTO) customMappingToEntity(db *gorm.DB) error {
	return nil
}
