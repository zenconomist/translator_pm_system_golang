package globalconstants

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

// ==================CONSTANTS===================\\

// ------------------CURRENCY---------------------\\

type Currency int

const (
	HUF Currency = iota
	EUR
	GBP
	USD
)

func (c Currency) String() string {
	switch c {
	case HUF:
		return "HUF"
	case EUR:
		return "EUR"
	case GBP:
		return "GBP"
	case USD:
		return "USD"
	}
	return ""
}

func StringToCurrency(cur string) (Currency, error) {
	switch cur {
	case "HUF":
		return HUF, nil
	case "EUR":
		return EUR, nil
	case "GBP":
		return GBP, nil
	case "USD":
		return USD, nil
	default:
		return 0, fmt.Errorf("not valid Currency")
	}
}

func (cur Currency) IsValid() error {
	switch cur {
	case HUF, EUR, GBP, USD:
		return nil
	default:
		return errors.New("invalid currency type")
	}
}

func (cur Currency) IsValidCurrencyString(c string) error {
	switch c {
	case "HUF", "EUR", "GBP", "USD":
		return nil
	default:
		return errors.New("invalid currency type")
	}
}

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

// ------------------PRICETYPE---------------------\\

type PriceType int

const (
	Agreed PriceType = iota
	Default
	Unique
)

func (pt PriceType) String() string {
	switch pt {
	case Agreed:
		return "Agreed"
	case Default:
		return "Default"
	case Unique:
		return "Unique"
	}
	return ""
}

// ------------------UNITTYPE---------------------\\

type UnitType int

const (
	Word UnitType = iota
	Character
	Row
	Page
	Document
	Hour
	Day
)

func (ut UnitType) String() string {
	switch ut {
	case Word:
		return "Word"
	case Character:
		return "Character"
	case Row:
		return "Row"
	case Page:
		return "Page"
	case Document:
		return "Document"
	case Hour:
		return "Hour"
	case Day:
		return "Day"
	}
	return ""
}

// ------------------INVOICELANGUAGE---------------------\\

type InvoiceLang int

const (
	HU InvoiceLang = iota
	DE
	EN
)

func (il InvoiceLang) String() string {
	switch il {
	case HU:
		return "hu"
	case DE:
		return "de"
	case EN:
		return "en"
	default:
		return ""
	}
}

func StringToInvoiceLang(il string) (InvoiceLang, error) {
	switch il {
	case "hu":
		return HU, nil
	case "de":
		return DE, nil
	case "en":
		return EN, nil
	default:
		return 0, fmt.Errorf("not valid Invoice language")
	}
}

// ------------------CRUD---------------------\\

type CRUD string

const (
	I CRUD = "Insert"
	U CRUD = "Update"
	D CRUD = "Delete"
)

var Infinity = time.Date(2099, 12, 31, 23, 59, 59, 0, time.UTC)
