package dto

import (
	"encoding/json"
	"entities"
	"fmt"
	"globalconstants"
	"logging"

	"gorm.io/gorm"
)

func (data *CustomerRequestDTO) CustomDTOTransformations(db *gorm.DB) error {
	// currency transform
	cur, errCurrency := globalconstants.StringToCurrency(data.CurrencyName)
	if errCurrency != nil {
		return fmt.Errorf("couldn't convert currency into valid currency, error: %v", errCurrency)
	}
	data.Currency = cur

	// invoicelang transform
	il, errIL := globalconstants.StringToInvoiceLang(data.InvoiceLangName)
	if errIL != nil {
		return fmt.Errorf("couldn't convert invoice language name into valid invoice language, error: %v", errIL)
	}
	data.InvoiceLang = il

	// firm
	var firm entities.Firm
	result := db.Where("name = ?", data.DefaultFirmName).First(&firm)
	if result.Error != nil {
		return fmt.Errorf("couldn't find firm id for firm name: %v", data.DefaultFirmName, " db error: %v", result.Error)
	}
	data.DefaultFirm = firm.Model.ID
	return nil

}

func (data *CustomerRequestDTO) CustomValidations(db *gorm.DB, logger logging.Logger) error {
	// implement validations
	return nil
}

func (data *CustomerRequestDTO) GiveID() (uint, error) {
	if data.ID == 0 {
		return data.ID, fmt.Errorf("the dto doesn't have any valid ID")
	}
	return data.ID, nil
}

func (data *CustomerRequestDTO) MapToEntity(db *gorm.DB, entity *entities.Customer) (*entities.Customer, error) {
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

func (data *CustomerRequestDTO) MapFromEntity(db *gorm.DB, entity entities.Customer) error {
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

func (data *CustomerRequestDTO) customMappingFromEntity(entity entities.Customer, db *gorm.DB) error {
	return nil
}

func (data *CustomerRequestDTO) customMappingToEntity(db *gorm.DB) error {
	return nil
}
