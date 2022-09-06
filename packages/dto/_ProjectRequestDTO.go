package dto

import (
	"entities"
	"fmt"
	"globalconstants"

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
		return fmt.Errorf("couldn't find firm id for firm name: %v", data.FirmName, " db error: %v", result.Error)
	}
	data.Firm = firm.Model.ID

	return nil
}
