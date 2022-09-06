package entities

// ---------- GENERATED ---------- \\

import (
    "time"
    dbc "dbconn"
    "globalconstants"
)

type BatchHistory struct {
    HistoryModel
    BatchStateID uint `json:"batch_state_id"` 
BatchTimeStateID uint `json:"batch_time_state_id"` 

}

func (item *BatchHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *BatchHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *BatchHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *BatchHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type CustomerPriceHistory struct {
    HistoryModel
    CustomerID uint `json:"customer_id"` 
TaskTypeID uint `json:"task_type"` 
SourceLang string `json:"source_lang"` 
TargetLang string `json:"target_lang"` 
Price float64 `json:"price"` 

}

func (item *CustomerPriceHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *CustomerPriceHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *CustomerPriceHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *CustomerPriceHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type DefaultSupplierPriceHistory struct {
    HistoryModel
    TaskTypeID uint `json:"task_type"` 
SourceLang string `json:"source_lang"` 
TargetLang string `json:"target_lang"` 
Currency globalconstants.Currency `json:"currency"` 
Price float64 `json:"price"` 

}

func (item *DefaultSupplierPriceHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *DefaultSupplierPriceHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *DefaultSupplierPriceHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *DefaultSupplierPriceHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type FirmHistory struct {
    HistoryModel
    DefaultCustomers []Customer `json:"customers" gorm:"-:all"` 
Projects []Project `json:"projects" gorm:"-:all"` 
Name string `json:"name" gorm:"not null"` 
FirmAddress Address `json:"address" gorm:"embedded"` 
MainEmail string `json:"main_email" gorm:"not null"` 

}

func (item *FirmHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *FirmHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *FirmHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *FirmHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type UPMLoggerHistory struct {
    HistoryModel
    UserID uint `json:"user_id"` 
Info string `json:"info"` 
FuncName string `json:"func_name"` 
EventStartTime time.Time `json:"log_begin"` 
EventEndTime time.Time `json:"log_end"` 
CriticalError bool `json:"critical_error"` 
ErrorMessage string `json:"error_message"` 
dbHandler dbc.DbHandler `gorm:"-:all"` 

}

func (item *UPMLoggerHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *UPMLoggerHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *UPMLoggerHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *UPMLoggerHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type BillingLogHistory struct {
    HistoryModel
    BillTStamp string `json:"billing_stamp"` 
RespBody string `json:"response_body"` 
ItemName string `json:"item_name"` 
UnitPrice float64 `json:"unit_price"` 
Currency string `json:"currency"` 
Quantity float64 `json:"item_quantity"` 
Unit string `json:"item_unit"` 
VAT string `json:"vat"` 
Entitlement string `json:"entitlement"` 
Comment string `json:"comment"` 
Language string `json:"billing_language"` 
FulfillmentDate string `json:"fulfillment_date"` 
DueDate string `json:"due_date"` 

}

func (item *BillingLogHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *BillingLogHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *BillingLogHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *BillingLogHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type QuestionHistory struct {
    HistoryModel
    TaskID uint `json:"task_id"` 
Question string `json:"question"` 
AskedBy uint `json:"asked_by"` 
QuestionAskedAt time.Time `json:"question_asked_at"` 
Answer string `json:"answer"` 
AnsweredBy uint `json:"answered_by"` 
AnsweredAt time.Time `json:"answered_at"` 
Closed bool `json:"closed"` 
ClosedBy uint `json:"closed_by"` 
ClosedAt time.Time `json:"closed_at"` 

}

func (item *QuestionHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *QuestionHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *QuestionHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *QuestionHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type ProjectHistory struct {
    HistoryModel
    CustomerID uint `json:"customer_id"` 
ClientOfferID uint `json:"client_offer_id"` 
ExternalProjectCode string `json:"external_project_code"` 
PurchaseOrder string `json:"purchase_order"` 
ProjectArrivalDate time.Time `json:"project_arrival_date"` 
ProjectDeadlineDate time.Time `json:"project_deadline_date"` 
FixedPrice float64 `json:"fixed_price"` 
SumToBill float64 `json:"sum_to_bill"` 
CurrencyName string `json:"currency_name"` 
Currency globalconstants.Currency `json:"currency"` 
ProjectStateID uint `json:"project_state_id"` 
ProjectStateName string `json:"project_state_name"` 
ProjectTimeStateID uint `json:"project_time_state_id"` 
ProjectTimeStateName string `json:"project_time_state_name"` 
FulfillmentDate time.Time `json:"fulfillment_date"` 
FirmID uint `json:"firm_id"` 
ProjectManager uint `json:"project_manager_id"` 
PartiallyBillable bool `json:"partially_billable"` 
Tasks []Task `json:"tasks" gorm:"-:all"` 

}

func (item *ProjectHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *ProjectHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *ProjectHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *ProjectHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type SharePointFolderConfigHistory struct {
    HistoryModel
    ConfigID int `json:"config_id"` 
FolderType string `json:"folder_type"` 
Hierarchy int `json:"hierarchy"` 
FolderName string `json:"folder_name"` 
ParentFolderConfigID int `json:"parent_folder_config_id"` 

}

func (item *SharePointFolderConfigHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *SharePointFolderConfigHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *SharePointFolderConfigHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *SharePointFolderConfigHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type CustomerHistory struct {
    HistoryModel
    Name string `json:"name"` 
Address Address `json:"address" gorm:"embedded"` 
Projects []Project `json:"projects" gorm:"-:all"` 
Prices []CustomerPrice `json:"prices" gorm:"-:all"` 
BillPID int64 `json:"bill_pid"` 
Email string `json:"email"` 
TaxCode string `json:"taxcode"` 
Iban string `json:"iban"` 
Swift string `json:"swift"` 
AccountNumber string `json:"account_number"` 
Phone string `json:"phone"` 
GeneralLedgerNumber string `json:"general_ledger_number"` 
TaxType string `json:"tax_type"` 
CustGeneralInfo string `json:"cust_general_info"` 
Currency globalconstants.Currency `json:"currency"` 
FirmID uint `json:"default_firm"` 
InvoiceLang globalconstants.InvoiceLang `json:"invoice_language"` 
PaymentDueDays int `json:"payment_due_days"` 
Contacts []Contact `json:"contacts" gorm:"-:all"` 

}

func (item *CustomerHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *CustomerHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *CustomerHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *CustomerHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type ContactHistory struct {
    HistoryModel
    CustomerID uint `json:"customer_id"` 
Default bool `json:"default"` 
Salutation string `json:"salutation"` 
FirstName string `json:"first_name"` 
LastName string `json:"last_name"` 
Position string `json:"position"` 
Email string `json:"email"` 
Email2 string `json:"email2"` 
Phone string `json:"phone"` 
Phone2 string `json:"phone2"` 
Phone3 string `json:"phone3"` 
Phone4 string `json:"phone4"` 
Fax string `json:"fax"` 
Comment string `json:"comment"` 

}

func (item *ContactHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *ContactHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *ContactHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *ContactHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type TaskOfferedHistory struct {
    HistoryModel
    TaskID uint `json:"task_id"` 
SupplierID uint `json:"supplier_id"` 
TaskOfferedStateID uint `json:"task_offered_state_id"` 

}

func (item *TaskOfferedHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *TaskOfferedHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *TaskOfferedHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *TaskOfferedHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type DefaultCustomerPriceHistory struct {
    HistoryModel
    TaskTypeID uint `json:"task_type"` 
SourceLang string `json:"source_lang"` 
TargetLang string `json:"target_lang"` 
Currency globalconstants.Currency `json:"currency"` 
Price float64 `json:"price"` 

}

func (item *DefaultCustomerPriceHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *DefaultCustomerPriceHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *DefaultCustomerPriceHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *DefaultCustomerPriceHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type ClientOfferTaskHistory struct {
    HistoryModel
    ProjectID uint `json:"project_id"` 
BatchID uint `json:"batch_id"` 
OrderWithinBatch uint `json:"order_within_batch"` 
TaskTypeID uint `json:"task_type"` 
ProjectManager uint `json:"project_manager_id"` 
SourceLang string `json:"source_lang"` 
TargetLang string `json:"target_lang"` 
TaskSpeciality string `json:"task_speciality"` 
PrepDisabled bool `json:"prepare_disabled"` 
ReviewDisabled bool `json:"review_disabled"` 
TaskCustomerProps TaskCustomerProps `json:"task_cust_props" gorm:"embedded"` 

}

func (item *ClientOfferTaskHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *ClientOfferTaskHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *ClientOfferTaskHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *ClientOfferTaskHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type ActualStateChangesHistory struct {
    HistoryModel
    From string `json:"from"` 
To string `json:"to"` 
IsAllowed bool `json:"is_allowed"` 
StateChangeInfo string `json:"state_change_info"` 
NeedsComment bool `json:"needs_comment"` 
PMPerform bool `json:"pm_perform"` 
PreparerPerfrom bool `json:"preparer_perfomr"` 
ReviewerPerform bool `json:"reviewer_perform"` 
SupplierPerform bool `json:"supplier_perform"` 
PMNotified bool `json:"pm_notified"` 
PreparerNotified bool `json:"preparer_notified"` 
ReviewerNotified bool `json:"reviewer_notified"` 
SupplierNotified bool `json:"supplier_notified"` 
ConditionOrExplanation string `json:"condition_or_explanation"` 

}

func (item *ActualStateChangesHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *ActualStateChangesHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *ActualStateChangesHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *ActualStateChangesHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type TaskStateChangeCommentHistory struct {
    HistoryModel
    TaskID uint `json:"task_id"` 
FromStateID uint `json:"from_state_id"` 
ToStateID uint `json:"to_state_id"` 
FromStateCode string `json:"from_state_code"` 
ToStateCode string `json:"to_state_code"` 
FromStateName string `json:"from_state_name"` 
ToStateName string `json:"to_state_name"` 
Comment string `json:"comment"` 
CommentedBy uint `json:"commented_by"` 
CommentedAt time.Time `json:"commented_at"` 

}

func (item *TaskStateChangeCommentHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *TaskStateChangeCommentHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *TaskStateChangeCommentHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *TaskStateChangeCommentHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type TaskConfigHistory struct {
    HistoryModel
    AddToDefault bool `json:"add_to_default"` 
TaskType string `json:"task_type"` 
SourceLang string `json:"source_lang"` 
TargetLang string `json:"target_lang"` 
DefTDDInterval int `json:"def_tdd_interval"` 
DefTDDIntervalType string `json:"def_tdd_interval_type"` 
TDDTime string `json:"tdd_time"` 
TaskSpeciality string `json:"task_speciality"` 
Billable bool `json:"billable"` 
CustomerUnitType string `json:"customer_unit_type"` 
SupplierUnitType string `json:"supplier_unit_type"` 
PrepDisabled bool `json:"prep_disabled"` 
PrepBillable bool `json:"prep_billable"` 
ReviewDisabled bool `json:"review_disabled"` 
ReviewBillable bool `json:"review_billable"` 
PreparerID uint `json:"preparer_id"` 
ReviewerID uint `json:"reviewer_id"` 
TranslBilling bool `json:"transl_billing"` 

}

func (item *TaskConfigHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *TaskConfigHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *TaskConfigHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *TaskConfigHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type ClientOfferHistory struct {
    HistoryModel
    CustomerID uint `json:"customer_id"` 
OfferRequestID string `json:"offerrequest_id"` 
FixDeadline bool `json:"fixdeadline_deadline"` 
DeadlineIntervalDays int `json:"deadlineintervaldays_days"` 
DeadlineIntervalHours int `json:"deadlineintervalhours_hours"` 
FixedPrice float64 `json:"fixedprice_price"` 
Currency globalconstants.Currency `gorm:"-:all"` 
CurrencyID uint `json:"currency_id"` 
PrSpecInfo string `json:"prspecinfo_info"` 
FirmID uint `json:"firm_id"` 
PartiallyBillable bool `json:"partiallybillable_billable"` 
ClientOfferTimeStateID uint `json:"clientoffertimestate_state_id"` 
ClientOfferStateID uint `json:"clientofferstate_state_id"` 
CopyOfferID uint `json:"copyoffer_id"` 
ClientOfferContactMail string `json:"clientoffercontactmail_mail"` 
ClientOfferContactName string `json:"clientoffercontactname_name"` 
ClientOfferInnerDetails string `json:"clientofferinnerdetails_details"` 
DiscountPercentage float64 `json:"discountpercentage_percentage"` 
DiscountSum float64 `json:"discountsum_sum"` 

}

func (item *ClientOfferHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *ClientOfferHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *ClientOfferHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *ClientOfferHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type TaskHistory struct {
    HistoryModel
    ProjectID uint `json:"project_id"` 
BatchID uint `json:"batch_id"` 
OrderWithinBatch uint `json:"order_within_batch"` 
TaskType string `json:"task_type"` 
ProjectManager uint `json:"project_manager_id"` 
TaskStateConfigHeadID uint `json:"task_state_config_head_id"` 
FormerTaskState uint `json:"former_task_state_id"` 
TaskStateID uint `json:"task_state_id"` 
TaskTimeStateID uint `json:"task_time_state_id"` 
SourceLang string `json:"source_lang"` 
TargetLang string `json:"target_lang"` 
TaskSpeciality string `json:"task_speciality"` 
PrepDisabled bool `json:"prepare_disabled"` 
PrepareState IPrepareState `gorm:"-:all"` 
PrepareStateID uint `json:"prepare_state_id"` 
ReviewDisabled bool `json:"review_disabled"` 
ReviewState IReviewState `gorm:"-:all"` 
ReviewStateID uint `json:"review_state_id"` 
PrepBillable bool `json:"prep_billable"` 
ReviewBillable bool `json:"review_billable"` 
PreparerID uint `json:"preparer_id"` 
ReviewerID uint `json:"reviewer_id"` 
TaskCustomerProps TaskCustomerProps `json:"task_cust_props" gorm:"embedded"` 
TaskSupplierProps TaskSupplierProps `json:"task_suppl_props" gorm:"embedded"` 
Questions []Question `json:"questions"` 
TaskStateChangeComments []TaskStateChangeComment `json:"task_state_change_comments"` 

}

func (item *TaskHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *TaskHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *TaskHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *TaskHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type SupplierPriceHistory struct {
    HistoryModel
    SupplierID uint `json:"supplier"` 
TaskTypeID uint `json:"task_type"` 
SourceLang string `json:"source_lang"` 
TargetLang string `json:"target_lang"` 
Price float64 `json:"price"` 

}

func (item *SupplierPriceHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *SupplierPriceHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *SupplierPriceHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *SupplierPriceHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type UserHistory struct {
    HistoryModel
    Name string `json:"name" gorm:"not null"` 
Password string `json:"pw" gorm:"not null"` 
UserName string `json:"user_name" gorm:"not null"` 
Email string `json:"email" gorm:"not null"` 
Active bool `json:"active"` 
Car bool `json:"car"` 
PrimaryPhoneNumber string `json:"primary_phone_number"` 
SecondaryPhoneNumber string `json:"secondary_phone_number"` 
Address Address `json:"address" gorm:"embedded"` 
VendorInvoiceFolderLink string `json:"vendor_invoice_folder_link"` 
BillingCurrency globalconstants.Currency `json:"billing_currency" gorm:"not null"` 
Prices []SupplierPrice `json:"prices" gorm:"-:all"` 
Specialities string `json:"specialities"` 
Languages string `json:"languages"` 
Projects []Project `json:"projects" gorm:"-:all"` 

}

func (item *UserHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *UserHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *UserHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *UserHistory) GiveID() uint {
	return item.HistoryModel.ID
}

type SharePointFolderHistory struct {
    HistoryModel
    FolderType string `json:"folder_type"` 
Hierarchy int `json:"hierarchy"` 
ProjectID uint `json:"project_id"` 
TaskID uint `json:"task_id"` 
SPCreatedTime string `json:"spcreated_time"` 
SPID string `json:"spid"` 
LastModified string `json:"last_modified"` 
FolderName string `json:"folder_name"` 
WebURL string `json:"web_url"` 
Size int `json:"size"` 
Context string `json:"context"` 
Etag string `json:"etag"` 
Ctag string `json:"ctag"` 
CreatedBy string `json:"created_by"` 
LastModifiedBy string `json:"last_modified_by"` 
ParentDriveID string `json:"parent_drive_id"` 
ParentDriveType string `json:"parent_drive_type"` 
ConfigID int `json:"config_id"` 
ParentFolderConfigID int `json:"parent_folder_config_id"` 
ParentSPID string `json:"parent_spid"` 
ParentPath string `json:"parent_path"` 
ChildCount int `json:"child_count"` 
ReadLink string `json:"read_link"` 
ReadWriteLink string `json:"read_write_link"` 

}

func (item *SharePointFolderHistory) SetModelToCreated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = true
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatFrom = time.Now()
	item.HistoryModel.DatTo = globalconstants.Infinity
}

func (item *SharePointFolderHistory) SetModelToUpdated(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = false
	item.HistoryModel.DatTo = time.Now()
}

func (item *SharePointFolderHistory) SetModelToDeleted(id uint) {
	item.HistoryModel.OriginalID = id
	item.HistoryModel.FlgIsCurrent = false
	item.HistoryModel.FlgIsDeleted = true
	item.HistoryModel.DatTo = time.Now()
}

func (item *SharePointFolderHistory) GiveID() uint {
	return item.HistoryModel.ID
}
