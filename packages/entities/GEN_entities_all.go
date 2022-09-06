package entities

// ---------- GENERATED ---------- \\

import (
    "time"
    dbc "dbconn"
    "globalconstants"
)

type Firm struct {
    Model
    DefaultCustomers []Customer `json:"customers"` 
Projects []Project `json:"projects"` 
Name string `json:"name" gorm:"unique;not null"` 
FirmAddress Address `json:"address" gorm:"embedded;embeddedPrefix:address_"` 
MainEmail string `json:"main_email" gorm:"not null"` 

}



func (item Firm) GiveID() uint {
	return item.Model.ID
}

type User struct {
    Model
    Name string `json:"name" gorm:"not null"` 
Password string `json:"pw" gorm:"not null"` 
UserName string `json:"user_name" gorm:"unique;not null"` 
Email string `json:"email" gorm:"unique;not null"` 
Active bool `json:"active" gorm:"default:true"` 
Car bool `json:"car"` 
PrimaryPhoneNumber string `json:"primary_phone_number"` 
SecondaryPhoneNumber string `json:"secondary_phone_number"` 
Address Address `json:"address" gorm:"embedded;embeddedPrefix:address_"` 
VendorInvoiceFolderLink string `json:"vendor_invoice_folder_link"` 
BillingCurrency globalconstants.Currency `json:"billing_currency" gorm:"not null"` 
Prices []SupplierPrice `json:"prices" gorm:"foreignKey:SupplierID"` 
Specialities string `json:"specialities"` 
Languages string `json:"languages"` 
Projects []Project `json:"projects" gorm:"foreignKey:ProjectManager"` 

}



func (item User) GiveID() uint {
	return item.Model.ID
}

type TaskState struct {
    Model
    TaskStateConfigHeadID uint `json:"task_state_config_head_id"` 
TaskStateCode string `json:"task_state_code"` 
TaskStateName string `json:"task_state_name"` 
TaskStateOrder int `json:"task_state_order"` 
Tasks []Task `json:"tasks"` 

}



func (item TaskState) GiveID() uint {
	return item.Model.ID
}

type Email struct {
    
    To string `json:"to"` 
Subject string `json:"subject"` 
Text string `json:"text"` 

}


type ClientOfferTaskCustomerProps struct {
    
    CustomerID uint `json:"customer_id"` 
CustQuantity float64 `json:"customer_quantity"` 
CustomerPrice float64 `json:"customer_price"` 
CustomerPriceType string `json:"customer_price_type"` 
CustomerPriceTypeID uint `json:"customer_price_type_id"` 
CustomerUnitType string `json:"customer_unit_type"` 
CustomerUnitTypeID uint `json:"customer_unit_type_id"` 
CustomerUniquePrice bool `json:"customer_unique_price"` 

}


type EmailSendingLog struct {
    Model
    SentByFunc string `json:"sent_by_func"` 
SentTo string `json:"sent_to"` 
Subject string `json:"subject"` 
Body string `json:"body"` 

}



func (item EmailSendingLog) GiveID() uint {
	return item.Model.ID
}

type TaskOffered struct {
    Model
    TaskID uint `json:"task_id"` 
SupplierID uint `json:"supplier_id" gorm:"foreignKey:UserID"` 
TaskOfferedState ITaskOfferedState `gorm:"-:all"` 
TaskOfferedStateID uint `json:"task_offered_state_id"` 

}



func (item TaskOffered) GiveID() uint {
	return item.Model.ID
}

type BillingLog struct {
    Model
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



func (item BillingLog) GiveID() uint {
	return item.Model.ID
}

type Question struct {
    Model
    TaskID uint `json:"task_id"` 
Question string `json:"question"` 
AskedBy uint `json:"asked_by" gorm:"foreignKey:UserID"` 
QuestionAskedAt time.Time `json:"question_asked_at"` 
Answer string `json:"answer"` 
AnsweredBy uint `json:"answered_by" gorm:"foreignKey:UserID"` 
AnsweredAt time.Time `json:"answered_at"` 
Closed bool `json:"closed"` 
ClosedBy uint `json:"closed_by" gorm:"foreignKey:UserID"` 
ClosedAt time.Time `json:"closed_at"` 

}



func (item Question) GiveID() uint {
	return item.Model.ID
}

type CustomerPrice struct {
    Model
    CustomerID uint `json:"customer_id"` 
TaskTypeID uint `json:"task_type"` 
SourceLang string `json:"source_lang"` 
TargetLang string `json:"target_lang"` 
Price float64 `json:"price"` 

}



func (item CustomerPrice) GiveID() uint {
	return item.Model.ID
}

type DefaultCustomerPrice struct {
    Model
    TaskTypeID uint `json:"task_type"` 
SourceLang string `json:"source_lang"` 
TargetLang string `json:"target_lang"` 
Currency globalconstants.Currency `json:"currency"` 
Price float64 `json:"price"` 

}



func (item DefaultCustomerPrice) GiveID() uint {
	return item.Model.ID
}

type DefaultSupplierPrice struct {
    Model
    TaskTypeID uint `json:"task_type"` 
SourceLang string `json:"source_lang"` 
TargetLang string `json:"target_lang"` 
Currency globalconstants.Currency `json:"currency"` 
Price float64 `json:"price"` 

}



func (item DefaultSupplierPrice) GiveID() uint {
	return item.Model.ID
}

type ClientOffer struct {
    Model
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
ClientOfferTimeState IClientOfferTimeState `gorm:"-:all"` 
ClientOfferTimeStateID uint `json:"clientoffertimestate_state_id"` 
ClientOfferState IClientOfferState `gorm:"-:all"` 
ClientOfferStateID uint `json:"clientofferstate_state_id"` 
CopyOfferID uint `json:"copyoffer_id"` 
ClientOfferContactMail string `json:"clientoffercontactmail_mail"` 
ClientOfferContactName string `json:"clientoffercontactname_name"` 
ClientOfferInnerDetails string `json:"clientofferinnerdetails_details"` 
DiscountPercentage float64 `json:"discountpercentage_percentage"` 
DiscountSum float64 `json:"discountsum_sum"` 

}



func (item ClientOffer) GiveID() uint {
	return item.Model.ID
}

type SharePointFolder struct {
    Model
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



func (item SharePointFolder) GiveID() uint {
	return item.Model.ID
}

type TaskTimeStateConfigHead struct {
    Model
    ActiveFrom time.Time `json:"active_from"` 
ActiveTo time.Time `json:"active_to"` 
TaskTimeStates []TaskTimeState `json:"task_time_states"` 

}



func (item TaskTimeStateConfigHead) GiveID() uint {
	return item.Model.ID
}

type TaskStateChangeComment struct {
    Model
    TaskID uint `json:"task_id"` 
FromStateID uint `json:"from_state_id" gorm:"foreignKey:TaskStateID"` 
ToStateID uint `json:"to_state_id" gorm:"foreignKey:TaskStateID"` 
FromStateCode string `json:"from_state_code"` 
ToStateCode string `json:"to_state_code"` 
FromStateName string `json:"from_state_name"` 
ToStateName string `json:"to_state_name"` 
Comment string `json:"comment"` 
CommentedBy uint `json:"commented_by" gorm:"foreignKey:UserID"` 
CommentedAt time.Time `json:"commented_at"` 

}



func (item TaskStateChangeComment) GiveID() uint {
	return item.Model.ID
}

type Customer struct {
    Model
    Name string `json:"name"` 
Address Address `json:"address" gorm:"embedded;embeddedPrefix:address_"` 
Projects []Project `json:"projects"` 
Prices []CustomerPrice `json:"prices"` 
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
Contacts []Contact `json:"contacts"` 

}



func (item Customer) GiveID() uint {
	return item.Model.ID
}

type SharePointFolderConfig struct {
    Model
    ConfigID int `json:"config_id"` 
FolderType string `json:"folder_type"` 
Hierarchy int `json:"hierarchy"` 
FolderName string `json:"folder_name"` 
ParentFolderConfigID int `json:"parent_folder_config_id"` 

}



func (item SharePointFolderConfig) GiveID() uint {
	return item.Model.ID
}

type TaskStateConfigHead struct {
    Model
    ActiveFrom time.Time `json:"active_from"` 
ActiveTo time.Time `json:"active_to"` 
TaskStates []TaskState `json:"task_states"` 

}



func (item TaskStateConfigHead) GiveID() uint {
	return item.Model.ID
}

type TaskStateChanges struct {
    Model
    FromTaskStateCode string `json:"from_task_state_code"` 
ToTaskStateCode string `json:"to_task_state_code"` 
FromTaskStateID uint `json:"from_task_state_id"` 
ToTaskStateID uint `json:"to_task_state_id"` 
PermissionID uint `json:"permission_id"` 
IsAllowed bool `json:"is_allowed"` 
StateChangeInfo string `json:"state_change_info"` 

}



func (item TaskStateChanges) GiveID() uint {
	return item.Model.ID
}

type TaskTimeState struct {
    Model
    TaskTimeStateConfigHeadID uint `json:"task_time_state_config_head_id"` 
TaskTimeStateCode string `json:"task_time_state_code"` 
TaskTimeStateName string `json:"task_time_state_name"` 
TaskTimeStateOrder int `json:"task_time_state_order_order"` 
Tasks []Task `json:"tasks"` 

}



func (item TaskTimeState) GiveID() uint {
	return item.Model.ID
}

type EmailNotifications struct {
    Model
    TaskStateChangesID uint `json:"task_state_changes_id"` 
ToList string `json:"to_list"` 
CCList string `json:"cc_list"` 
BCCList string `json:"bcc_list"` 
HTMLTemplate string `json:"html_template"` 
Subject string `json:"subject"` 

}



func (item EmailNotifications) GiveID() uint {
	return item.Model.ID
}

type Contact struct {
    Model
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



func (item Contact) GiveID() uint {
	return item.Model.ID
}

type Address struct {
    
    CountryCode string `json:"country_code"` 
PostCode string `json:"post_code"` 
City string `json:"city"` 
Address string `json:"address"` 

}


type Project struct {
    Model
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
ProjectState IProjectState `gorm:"-:all"` 
ProjectStateID uint `json:"project_state_id"` 
ProjectStateName string `json:"project_state_name"` 
ProjectTimeState IProjectTimeState `gorm:"-:all"` 
ProjectTimeStateID uint `json:"project_time_state_id"` 
ProjectTimeStateName string `json:"project_time_state_name"` 
FulfillmentDate time.Time `json:"fulfillment_date"` 
FirmID uint `json:"firm_id" gorm:"foreignKey:FirmID"` 
ProjectManager uint `json:"project_manager_id" gorm:"foreignKey:UserID"` 
PartiallyBillable bool `json:"partially_billable"` 
Tasks []Task `json:"tasks"` 

}



func (item Project) GiveID() uint {
	return item.Model.ID
}

type Task struct {
    Model
    ProjectID uint `json:"project_id"` 
BatchID uint `json:"batch_id"` 
OrderWithinBatch uint `json:"order_within_batch"` 
TaskType string `json:"task_type"` 
ProjectManager uint `json:"project_manager_id" gorm:"foreignKey:UserID"` 
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
PreparerID uint `json:"preparer_id" gorm:"foreignKey:UserID"` 
ReviewerID uint `json:"reviewer_id" gorm:"foreignKey:UserID"` 
TaskCustomerProps TaskCustomerProps `json:"task_cust_props" gorm:"embedded;embeddedPrefix:tcp_"` 
TaskSupplierProps TaskSupplierProps `json:"task_suppl_props" gorm:"embedded;embeddedPrefix:tsp_"` 
Questions []Question `json:"questions"` 
TaskStateChangeComments []TaskStateChangeComment `json:"task_state_change_comments"` 

}



func (item Task) GiveID() uint {
	return item.Model.ID
}

type TaskConfig struct {
    Model
    AddToDefault bool `json:"add_to_default"` 
TaskType string `json:"task_type" gorm:"unique;not null"` 
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
PreparerID uint `json:"preparer_id" gorm:"foreignKey:UserID"` 
ReviewerID uint `json:"reviewer_id" gorm:"foreignKey:UserID"` 
TranslBilling bool `json:"transl_billing"` 

}



func (item TaskConfig) GiveID() uint {
	return item.Model.ID
}

type ActualStateChanges struct {
    Model
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



func (item ActualStateChanges) GiveID() uint {
	return item.Model.ID
}

type Permission struct {
    Model
    PermissionName string `json:"permission_name"` 

}



func (item Permission) GiveID() uint {
	return item.Model.ID
}

type Batch struct {
    Model
    BatchState IBatchState `gorm:"-:all"` 
BatchStateID uint `json:"batch_state_id"` 
BatchTimeState IBatchTimeState `gorm:"-:all"` 
BatchTimeStateID uint `json:"batch_time_state_id"` 

}



func (item Batch) GiveID() uint {
	return item.Model.ID
}

type SupplierPrice struct {
    Model
    SupplierID uint `json:"supplier" gorm:"foreignKey:UserID"` 
TaskTypeID uint `json:"task_type"` 
SourceLang string `json:"source_lang"` 
TargetLang string `json:"target_lang"` 
Price float64 `json:"price"` 

}



func (item SupplierPrice) GiveID() uint {
	return item.Model.ID
}

type UPMLogger struct {
    Model
    UserID uint `json:"user_id"` 
Info string `json:"info"` 
FuncName string `json:"func_name"` 
EventStartTime time.Time `json:"log_begin"` 
EventEndTime time.Time `json:"log_end"` 
CriticalError bool `json:"critical_error"` 
ErrorMessage string `json:"error_message"` 
dbHandler dbc.DbHandler `gorm:"-:all"` 

}



func (item UPMLogger) GiveID() uint {
	return item.Model.ID
}

type ClientOfferTask struct {
    Model
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



func (item ClientOfferTask) GiveID() uint {
	return item.Model.ID
}

type Role struct {
    Model
    RoleName string `json:"role_name"` 
Users []User `json:"users" gorm:"many2many:user_role;"` 
Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;"` 

}



func (item Role) GiveID() uint {
	return item.Model.ID
}

type TaskCustomerProps struct {
    
    CustomerDueDate time.Time `json:"customer_due_date"` 
CustQuantity float64 `json:"customer_quantity"` 
CustomerPrice float64 `json:"customer_price"` 
CustomerPriceType globalconstants.PriceType `gorm:"-:all"` 
CustomerPriceTypeID uint `json:"customer_price_type_id"` 
CustomerUnitType globalconstants.UnitType `gorm:"-:all"` 
CustomerUnitTypeID uint `json:"customer_unit_type_id"` 
CustomerUniquePrice bool `json:"customer_unique_price"` 

}


type TaskSupplierProps struct {
    
    SupplierID uint `json:"supplier_id" gorm:"foreignKey:UserID"` 
SupplierDueDate time.Time `json:"supplier_due_date"` 
SupplierPriceType globalconstants.PriceType `gorm:"-:all"` 
SupplierPriceTypeID uint `json:"supplier_price_type_id"` 
SupplierUnitType globalconstants.UnitType `gorm:"-:all"` 
SupplierUnitTypeID uint `json:"supplier_unit_type_id"` 
SupplierQuantity float64 `json:"supplier_quantity"` 
SupplierToBill bool `json:"supplier_to_bill"` 
SupplierTimeState ISupplierTimeState `gorm:"-:all"` 
SupplierTimeStateID uint `json:"supplier_time_state_id"` 
BilledBySupplier bool `json:"billed_by_supplier"` 
BillingFulfillmentDate string `json:"billing_fulfillment_date"` 
BillingInvoiceNumber string `json:"billing_invoice_number"` 
SupplierUniquePrice bool `json:"supplier_unique_price"` 

}


type TaskTimeStateChanges struct {
    Model
    FromTaskTimeStateCode string `json:"from_task_time_state_code"` 
ToTaskTimeStateCode string `json:"to_task_time_state_code"` 

}



func (item TaskTimeStateChanges) GiveID() uint {
	return item.Model.ID
}
