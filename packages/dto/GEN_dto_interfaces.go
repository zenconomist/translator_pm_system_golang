package dto

// ---------- GENERATED ---------- \\

import (
	"entities"
	"logging"

	"gorm.io/gorm"
)

type FirmDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.Firm, error)
	MapFromEntity(*gorm.DB, entities.Firm) error
}

type PermissionDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.Permission, error)
	MapFromEntity(*gorm.DB, entities.Permission) error
}

type ContactDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.Contact, error)
	MapFromEntity(*gorm.DB, entities.Contact) error
}

type SharePointFolderDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.SharePointFolder, error)
	MapFromEntity(*gorm.DB, entities.SharePointFolder) error
}

type AddressDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.Address, error)
	MapFromEntity(*gorm.DB, entities.Address) error
}

type TaskCustomerPropsDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.TaskCustomerProps, error)
	MapFromEntity(*gorm.DB, entities.TaskCustomerProps) error
}

type UPMLoggerDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.UPMLogger, error)
	MapFromEntity(*gorm.DB, entities.UPMLogger) error
}

type ProjectDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.Project, error)
	MapFromEntity(*gorm.DB, entities.Project) error
}

type BatchDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.Batch, error)
	MapFromEntity(*gorm.DB, entities.Batch) error
}

type TaskConfigDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.TaskConfig, error)
	MapFromEntity(*gorm.DB, entities.TaskConfig) error
}

type ClientOfferDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.ClientOffer, error)
	MapFromEntity(*gorm.DB, entities.ClientOffer) error
}

type DefaultSupplierPriceDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.DefaultSupplierPrice, error)
	MapFromEntity(*gorm.DB, entities.DefaultSupplierPrice) error
}

type TaskDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.Task, error)
	MapFromEntity(*gorm.DB, entities.Task) error
}

type SupplierPriceDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.SupplierPrice, error)
	MapFromEntity(*gorm.DB, entities.SupplierPrice) error
}

type ClientOfferTaskCustomerPropsDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.ClientOfferTaskCustomerProps, error)
	MapFromEntity(*gorm.DB, entities.ClientOfferTaskCustomerProps) error
}

type DefaultCustomerPriceDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.DefaultCustomerPrice, error)
	MapFromEntity(*gorm.DB, entities.DefaultCustomerPrice) error
}

type SharePointFolderConfigDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.SharePointFolderConfig, error)
	MapFromEntity(*gorm.DB, entities.SharePointFolderConfig) error
}

type BillingLogDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.BillingLog, error)
	MapFromEntity(*gorm.DB, entities.BillingLog) error
}

type EmailSendingLogDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.EmailSendingLog, error)
	MapFromEntity(*gorm.DB, entities.EmailSendingLog) error
}

type CustomerDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.Customer, error)
	MapFromEntity(*gorm.DB, entities.Customer) error
}

type UserDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.User, error)
	MapFromEntity(*gorm.DB, entities.User) error
}

type TaskOfferedDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.TaskOffered, error)
	MapFromEntity(*gorm.DB, entities.TaskOffered) error
}

type ClientOfferTaskDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.ClientOfferTask, error)
	MapFromEntity(*gorm.DB, entities.ClientOfferTask) error
}

type RoleDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.Role, error)
	MapFromEntity(*gorm.DB, entities.Role) error
}

type TaskSupplierPropsDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.TaskSupplierProps, error)
	MapFromEntity(*gorm.DB, entities.TaskSupplierProps) error
}

type EmailDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.Email, error)
	MapFromEntity(*gorm.DB, entities.Email) error
}

type CustomerPriceDTO interface {
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	GiveID() (uint, error)
	MapToEntity(*gorm.DB) (entities.CustomerPrice, error)
	MapFromEntity(*gorm.DB, entities.CustomerPrice) error
}
