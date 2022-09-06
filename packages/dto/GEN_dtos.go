package dto

import (
	"globalconstants"
	"time"
)

// ---------- GENERATED ---------- \\

type ProjectResponseDTO struct {
	ID                  uint
	CustomerID          uint              `json:"customer_id"`
	ClientOfferID       uint              `json:"client_offer_id"`
	ExternalProjectCode string            `json:"external_project_code"`
	PurchaseOrder       string            `json:"purchase_order"`
	ProjectArrivalDate  time.Time         `json:"project_arrival_date"`
	ProjectDeadlineDate time.Time         `json:"project_deadline_date"`
	FixedPrice          float64           `json:"fixed_price"`
	SumToBill           float64           `json:"sum_to_bill"`
	Currency            string            `json:"currency_name"`
	ProjectState        string            `json:"project_state"`
	ProjectStateID      uint              `json:"project_state_id"`
	ProjectTimeState    string            `json:"project_time_state"`
	ProjectTimeStateID  uint              `json:"project_time_state_id"`
	FulfillmentDate     time.Time         `json:"fulfillment_date"`
	Firm                uint              `json:"firm_id"`
	ProjectManager      uint              `json:"project_manager_id"`
	PartiallyBillable   bool              `json:"partially_billable"`
	Tasks               []TaskResponseDTO `json:"tasks"`
}

type ClientOfferResponseDTO struct {
	ID                      uint
	CustomerID              uint    `json:"customer_id"`
	OfferRequestID          string  `json:"offerrequest_id"`
	FixDeadline             bool    `json:"fixdeadline_deadline"`
	DeadlineIntervalDays    int     `json:"deadlineintervaldays_days"`
	DeadlineIntervalHours   int     `json:"deadlineintervalhours_hours"`
	FixedPrice              float64 `json:"fixedprice_price"`
	Currency                string  `json:"currency"`
	CurrencyID              uint    `json:"currency_id"`
	PrSpecInfo              string  `json:"prspecinfo_info"`
	FirmID                  uint    `json:"firm_id"`
	PartiallyBillable       bool    `json:"partiallybillable_billable"`
	ClientOfferTimeState    string  `json:"client_offer_time_state"`
	ClientOfferTimeStateID  uint    `json:"clientoffertimestate_state_id"`
	ClientOfferState        string  `json:"client_offer_state"`
	ClientOfferStateID      uint    `json:"clientofferstate_state_id"`
	CopyOfferID             uint    `json:"copyoffer_id"`
	ClientOfferContactMail  string  `json:"clientoffercontactmail_mail"`
	ClientOfferContactName  string  `json:"clientoffercontactname_name"`
	ClientOfferInnerDetails string  `json:"clientofferinnerdetails_details"`
	DiscountPercentage      float64 `json:"discountpercentage_percentage"`
	DiscountSum             float64 `json:"discountsum_sum"`
	VATType                 string  `json:"vat_type"`
	VATPercentage           int     `json:"vat_percentage"`
	VATAmount               float64 `json:"vat_amount"`
	GrossAmount             float64 `json:"gross_amount"`
}

type CustomerRequestDTO struct {
	ID                  uint
	Name                string                      `json:"name"`
	Address             AddressRequestDTO           `json:"address"`
	Prices              []CustomerPriceRequestDTO   `json:"prices"`
	BillPID             int64                       `json:"id"`
	Email               string                      `json:"email"`
	TaxCode             string                      `json:"taxcode"`
	Iban                string                      `json:"iban"`
	Swift               string                      `json:"swift"`
	AccountNumber       string                      `json:"account_number"`
	Phone               string                      `json:"phone"`
	GeneralLedgerNumber string                      `json:"general_ledger_number"`
	TaxType             string                      `json:"tax_type"`
	CustGeneralInfo     string                      `json:"cust_general_info"`
	CurrencyName        string                      `json:"currency_name"`
	Currency            globalconstants.Currency    `json:"currency"`
	DefaultFirmName     string                      `json:"default_firm_name"`
	DefaultFirm         uint                        `json:"default_firm"`
	InvoiceLangName     string                      `json:"invoice_language_name"`
	InvoiceLang         globalconstants.InvoiceLang `json:"invoice_language"`
	PaymentDueDays      int                         `json:"payment_due_days"`
}

type CustomerPriceRequestDTO struct {
	ID         uint
	CustomerID uint    `json:"customer_id"`
	TaskTypeID uint    `json:"task_type"`
	SourceLang string  `json:"source_lang"`
	TargetLang string  `json:"target_lang"`
	CurrencyID uint    `json:"currency"`
	Price      float64 `json:"price"`
}

type ClientOfferTaskRequestDTO struct {
	ID                uint
	ProjectID         uint                        `json:"project_id"`
	BatchID           uint                        `json:"batch_id"`
	OrderWithinBatch  uint                        `json:"order_within_batch"`
	TaskTypeID        uint                        `json:"task_type"`
	ProjectManager    uint                        `json:"project_manager_id"`
	TaskState         string                      `json:"task_state"`
	TaskStateID       uint                        `json:"task_state_id"`
	TaskTimeState     string                      `json:"task_time_state"`
	TaskTimeStateID   uint                        `json:"task_time_state_id"`
	SourceLang        string                      `json:"source_lang"`
	TargetLang        string                      `json:"target_lang"`
	PrepDisabled      bool                        `json:"prepare_disabled"`
	ReviewDisabled    bool                        `json:"review_disabled"`
	TaskCustomerProps TaskCustomerPropsRequestDTO `json:"task_cust_props"`
	TaskSupplierProps TaskSupplierPropsRequestDTO `json:"task_suppl_props"`
	TaskSpeciality    string                      `json:"task_speciality"`
}

type UPMLoggerResponseDTO struct {
	ID             uint
	UserID         uint      `json:"user_id"`
	Info           string    `json:"info"`
	FuncName       string    `json:"func_name"`
	EventStartTime time.Time `json:"log_begin"`
	EventEndTime   time.Time `json:"log_end"`
	CriticalError  bool      `json:"critical_error"`
	ErrorMessage   string    `json:"error_message"`
}

type TaskResponseDTO struct {
	ID                uint
	ProjectID         uint                         `json:"project_id"`
	BatchID           uint                         `json:"batch_id"`
	OrderWithinBatch  uint                         `json:"order_within_batch"`
	TaskTypeID        uint                         `json:"task_type"`
	ProjectManager    uint                         `json:"project_manager_id"`
	TaskStateName     string                       `json:"task_state_name"`
	TaskStateID       uint                         `json:"task_state_id"`
	TaskTimeStateName string                       `json:"task_time_state_name"`
	TaskTimeStateID   uint                         `json:"task_time_state_id"`
	SourceLang        string                       `json:"source_lang"`
	TargetLang        string                       `json:"target_lang"`
	PrepDisabled      bool                         `json:"prepare_disabled"`
	PrepareStateName  string                       `json:"prepare_state_name"`
	PrepareStateID    uint                         `json:"prepare_state_id"`
	ReviewDisabled    bool                         `json:"review_disabled"`
	ReviewStateName   string                       `json:"review_state_name"`
	ReviewStateID     uint                         `json:"review_state_id"`
	TaskCustomerProps TaskCustomerPropsResponseDTO `json:"task_cust_props"`
	TaskSupplierProps TaskSupplierPropsResponseDTO `json:"task_suppl_props"`
	TaskSpeciality    string                       `json:"task_speciality"`
}

type TaskOfferedResponseDTO struct {
	ID                 uint
	TaskID             uint   `json:"task_id"`
	SupplierID         uint   `json:"supplier_id"`
	TaskOfferedState   string `json:"task_offered_state"`
	TaskOfferedStateID uint   `json:"task_offered_state_id"`
}

type TaskSupplierPropsRequestDTO struct {
	ID                     uint
	SupplierID             uint      `json:"supplier_id"`
	SupplierDueDateDate    string    `json:"supplier_due_date_date"`
	SupplierDueDateHour    int       `json:"supplier_due_date_hour"`
	SupplierDueDate        time.Time `json:"supplier_due_date"`
	SupplierPriceType      string    `json:"supplier_price_type"`
	SupplierPriceTypeID    uint      `json:"supplier_price_type_id"`
	SupplierUnitType       string    `json:"supplier_unit_type"`
	SupplierUnitTypeID     uint      `json:"supplier_unit_type_id"`
	SupplierQuantity       float64   `json:"supplier_quantity"`
	SupplierToBill         bool      `json:"supplier_to_bill"`
	SupplierTimeState      string    `json:"supplier_time_state"`
	SupplierTimeStateID    uint      `json:"supplier_time_state_id"`
	BilledBySupplier       bool      `json:"billed_by_supplier"`
	BillingFulfillmentDate string    `json:"billing_fulfillment_date"`
	BillingInvoiceNumber   string    `json:"billing_invoice_number"`
	SupplierUniquePrice    bool      `json:"supplier_unique_price"`
}

type TaskOfferedRequestDTO struct {
	ID                 uint
	TaskID             uint   `json:"task_id"`
	SupplierID         uint   `json:"supplier_id"`
	TaskOfferedState   string `json:"task_offered_state"`
	TaskOfferedStateID uint   `json:"task_offered_state_id"`
}

type DefaultSupplierPriceRequestDTO struct {
	ID         uint
	TaskTypeID uint                     `json:"task_type"`
	SourceLang string                   `json:"source_lang"`
	TargetLang string                   `json:"target_lang"`
	Currency   globalconstants.Currency `json:"currency"`
	Price      float64                  `json:"price"`
}

type ClientOfferTaskCustomerPropsRequestDTO struct {
	ID                  uint
	CustomerID          uint      `json:"customer_id"`
	CustomerDueDate     time.Time `json:"customer_due_date"`
	CustQuantity        float64   `json:"customer_quantity"`
	CustomerPrice       float64   `json:"customer_price"`
	CustomerPriceType   string    `json:"customer_price_type"`
	CustomerPriceTypeID uint      `json:"customer_price_type_id"`
	CustomerUnitType    string    `json:"customer_unit_type"`
	CustomerUnitTypeID  uint      `json:"customer_unit_type_id"`
	CustomerUniquePrice bool      `json:"customer_unique_price"`
}

type BillingLogResponseDTO struct {
	ID              uint
	BillTStamp      string  `json:"billing_stamp"`
	RespBody        string  `json:"response_body"`
	ItemName        string  `json:"item_name"`
	UnitPrice       float64 `json:"unit_price"`
	Currency        string  `json:"currency"`
	Quantity        float64 `json:"item_quantity"`
	Unit            string  `json:"item_unit"`
	VAT             string  `json:"vat"`
	Entitlement     string  `json:"entitlement"`
	Comment         string  `json:"comment"`
	Language        string  `json:"billing_language"`
	FulfillmentDate string  `json:"fulfillment_date"`
	DueDate         string  `json:"due_date"`
}

type BatchResponseDTO struct {
	ID               uint
	Tasks            []TaskResponseDTO `json:"tasks"`
	BatchState       string            `json:"batch_state"`
	BatchStateID     uint              `json:"batch_state_id"`
	BatchTimeState   string            `json:"batch_time_state"`
	BatchTimeStateID uint              `json:"batch_time_state_id"`
}

type SupplierPriceResponseDTO struct {
	ID         uint
	SupplierID uint    `json:"supplier"`
	TaskTypeID uint    `json:"task_type"`
	SourceLang string  `json:"source_lang"`
	TargetLang string  `json:"target_lang"`
	CurrencyID uint    `json:"currency"`
	Default    bool    `json:"default"`
	Price      float64 `json:"price"`
}

type TaskRequestDTO struct {
	ID                uint
	ProjectID         uint                        `json:"project_id"`
	BatchID           uint                        `json:"batch_id"`
	OrderWithinBatch  uint                        `json:"order_within_batch"`
	TaskType          string                      `json:"task_type"`
	ProjectManager    uint                        `json:"project_manager_id"`
	TaskStateName     string                      `json:"task_state_name"`
	TaskStateID       uint                        `json:"task_state_id"`
	TaskTimeStateName string                      `json:"task_time_state_name"`
	TaskTimeStateID   uint                        `json:"task_time_state_id"`
	SourceLang        string                      `json:"source_lang"`
	TargetLang        string                      `json:"target_lang"`
	PrepDisabled      bool                        `json:"prepare_disabled"`
	PrepareStateName  string                      `json:"prepare_state_name"`
	PrepareStateID    uint                        `json:"prepare_state_id"`
	ReviewDisabled    bool                        `json:"review_disabled"`
	ReviewStateName   string                      `json:"review_state_name"`
	ReviewStateID     uint                        `json:"review_state_id"`
	PrepBillable      bool                        `json:"prep_billable"`
	ReviewBillable    bool                        `json:"review_billable"`
	PreparerID        uint                        `json:"preparer_id"`
	ReviewerID        uint                        `json:"reviewer_id"`
	TaskSpeciality    string                      `json:"task_speciality"`
	TaskCustomerProps TaskCustomerPropsRequestDTO `json:"task_cust_props"`
	TaskSupplierProps TaskSupplierPropsRequestDTO `json:"task_suppl_props"`
}

type TaskCustomerPropsRequestDTO struct {
	ID                  uint
	CustomerDueDateDate string    `json:"customer_due_date_date"`
	CustomerDueDateHour int       `json:"customer_due_date_hour"`
	CustomerDueDate     time.Time `json:"customer_due_date"`
	CustQuantity        float64   `json:"customer_quantity"`
	CustomerPrice       float64   `json:"customer_price"`
	CustomerPriceType   string    `json:"customer_price_type"`
	CustomerPriceTypeID uint      `json:"customer_price_type_id"`
	CustomerUnitType    string    `json:"customer_unit_type"`
	CustomerUnitTypeID  uint      `json:"customer_unit_type_id"`
	CustomerUniquePrice bool      `json:"customer_unique_price"`
}

type SupplierPriceRequestDTO struct {
	ID         uint
	SupplierID uint    `json:"supplier"`
	TaskTypeID uint    `json:"task_type"`
	SourceLang string  `json:"source_lang"`
	TargetLang string  `json:"target_lang"`
	CurrencyID uint    `json:"currency"`
	Price      float64 `json:"price"`
}

type ClientOfferTaskCustomerPropsResponseDTO struct {
	ID                  uint
	CustomerID          uint      `json:"customer_id"`
	CustomerDueDate     time.Time `json:"customer_due_date"`
	CustQuantity        float64   `json:"customer_quantity"`
	CustomerPrice       float64   `json:"customer_price"`
	CustomerPriceType   string    `json:"customer_price_type"`
	CustomerPriceTypeID uint      `json:"customer_price_type_id"`
	CustomerUnitType    string    `json:"customer_unit_type"`
	CustomerUnitTypeID  uint      `json:"customer_unit_type_id"`
	CustomerUniquePrice bool      `json:"customer_unique_price"`
}

type ClientOfferRequestDTO struct {
	ID                      uint
	CustomerID              uint    `json:"customer_id"`
	OfferRequestID          string  `json:"offerrequest_id"`
	FixDeadline             bool    `json:"fixdeadline_deadline"`
	DeadlineIntervalDays    int     `json:"deadlineintervaldays_days"`
	DeadlineIntervalHours   int     `json:"deadlineintervalhours_hours"`
	FixedPrice              float64 `json:"fixedprice_price"`
	Currency                string  `json:"currency"`
	CurrencyID              uint    `json:"currency_id"`
	PrSpecInfo              string  `json:"prspecinfo_info"`
	FirmID                  uint    `json:"firm_id"`
	PartiallyBillable       bool    `json:"partiallybillable_billable"`
	ClientOfferTimeState    string  `json:"client_offer_time_state"`
	ClientOfferTimeStateID  uint    `json:"clientoffertimestate_state_id"`
	ClientOfferState        string  `json:"client_offer_state"`
	ClientOfferStateID      uint    `json:"clientofferstate_state_id"`
	CopyOfferID             uint    `json:"copyoffer_id"`
	ClientOfferContactMail  string  `json:"clientoffercontactmail_mail"`
	ClientOfferContactName  string  `json:"clientoffercontactname_name"`
	ClientOfferInnerDetails string  `json:"clientofferinnerdetails_details"`
	DiscountPercentage      float64 `json:"discountpercentage_percentage"`
	DiscountSum             float64 `json:"discountsum_sum"`
}

type PermissionResponseDTO struct {
	ID             uint
	PermissionName string `json:"permission_name"`
}

type DefaultSupplierPriceResponseDTO struct {
	ID         uint
	TaskTypeID uint                     `json:"task_type"`
	SourceLang string                   `json:"source_lang"`
	TargetLang string                   `json:"target_lang"`
	Currency   globalconstants.Currency `json:"currency"`
	Price      float64                  `json:"price"`
}

type ClientOfferTaskResponseDTO struct {
	ID                uint
	ProjectID         uint                         `json:"project_id"`
	BatchID           uint                         `json:"batch_id"`
	OrderWithinBatch  uint                         `json:"order_within_batch"`
	TaskTypeID        uint                         `json:"task_type"`
	ProjectManager    uint                         `json:"project_manager_id"`
	TaskState         string                       `json:"task_state"`
	TaskStateID       uint                         `json:"task_state_id"`
	TaskTimeState     string                       `json:"task_time_state"`
	TaskTimeStateID   uint                         `json:"task_time_state_id"`
	SourceLang        string                       `json:"source_lang"`
	TargetLang        string                       `json:"target_lang"`
	PrepDisabled      bool                         `json:"prepare_disabled"`
	ReviewDisabled    bool                         `json:"review_disabled"`
	TaskCustomerProps TaskCustomerPropsResponseDTO `json:"task_cust_props"`
	TaskSupplierProps TaskSupplierPropsResponseDTO `json:"task_suppl_props"`
	TaskSpeciality    string                       `json:"task_speciality"`
}

type AddressRequestDTO struct {
	ID          uint
	CountryCode string `json:"country_code"`
	PostCode    string `json:"post_code"`
	City        string `json:"city"`
	Address     string `json:"address"`
}

type BatchRequestDTO struct {
	ID               uint
	Tasks            []TaskRequestDTO `json:"tasks"`
	BatchState       string           `json:"batch_state"`
	BatchStateID     uint             `json:"batch_state_id"`
	BatchTimeState   string           `json:"batch_time_state"`
	BatchTimeStateID uint             `json:"batch_time_state_id"`
}

type TaskConfigRequestDTO struct {
	ID                 uint
	AddToDefault       bool   `json:"add_to_default"`
	TaskType           string `json:"task_type"`
	SourceLang         string `json:"source_lang"`
	TargetLang         string `json:"target_lang"`
	DefTDDInterval     int    `json:"def_tdd_interval"`
	DefTDDIntervalType string `json:"def_tdd_interval_type"`
	TDDTime            string `json:"tdd_time"`
	TaskSpeciality     string `json:"task_speciality"`
	Billable           bool   `json:"billable"`
	CustomerUnitType   string `json:"customer_unit_type"`
	SupplierUnitType   string `json:"supplier_unit_type"`
	PrepDisabled       bool   `json:"prep_disabled"`
	PrepareState       string `json:"prepare_state"`
	PrepareStateID     uint   `json:"prepare_state_id"`
	PrepBillable       bool   `json:"prep_billable"`
	ReviewDisabled     bool   `json:"review_disabled"`
	ReviewState        string `json:"review_state"`
	ReviewStateID      uint   `json:"review_state_id"`
	ReviewBillable     bool   `json:"review_billable"`
	PreparerUserName   string `json:"preparer_user_name"`
	PreparerID         uint   `json:"preparer_id"`
	ReviewerUserName   string `json:"reviewer_user_name"`
	ReviewerID         uint   `json:"reviewer_id"`
	TranslBilling      bool   `json:"transl_billing"`
}

type FirmResponseDTO struct {
	ID          uint
	Name        string             `json:"name"`
	FirmAddress AddressResponseDTO `json:"address"`
	MainEmail   string             `json:"main_email"`
}

type TaskSupplierPropsResponseDTO struct {
	ID                     uint
	SupplierID             uint      `json:"supplier_id"`
	SupplierDueDate        time.Time `json:"supplier_due_date"`
	SupplierPriceType      string    `json:"supplier_price_type"`
	SupplierPriceTypeID    uint      `json:"supplier_price_type_id"`
	SupplierUnitType       string    `json:"supplier_unit_type"`
	SupplierUnitTypeID     uint      `json:"supplier_unit_type_id"`
	SupplierQuantity       float64   `json:"supplier_quantity"`
	SupplierToBill         bool      `json:"supplier_to_bill"`
	SupplierTimeState      string    `json:"supplier_time_state"`
	SupplierTimeStateID    uint      `json:"supplier_time_state_id"`
	BilledBySupplier       bool      `json:"billed_by_supplier"`
	BillingFulfillmentDate string    `json:"billing_fulfillment_date"`
	BillingInvoiceNumber   string    `json:"billing_invoice_number"`
	SupplierUniquePrice    bool      `json:"supplier_unique_price"`
}

type TaskConfigResponseDTO struct {
	ID                 uint
	AddToDefault       bool   `json:"add_to_default"`
	TaskType           string `json:"task_type"`
	SourceLang         string `json:"source_lang"`
	TargetLang         string `json:"target_lang"`
	DefTDDInterval     int    `json:"def_tdd_interval"`
	DefTDDIntervalType string `json:"def_tdd_interval_type"`
	TDDTime            string `json:"tdd_time"`
	TaskSpeciality     string `json:"task_speciality"`
	Billable           bool   `json:"billable"`
	CustomerUnitType   string `json:"customer_unit_type"`
	SupplierUnitType   string `json:"supplier_unit_type"`
	PrepDisabled       bool   `json:"prep_disabled"`
	PrepareState       string `json:"prepare_state"`
	PrepareStateID     uint   `json:"prepare_state_id"`
	PrepBillable       bool   `json:"prep_billable"`
	ReviewDisabled     bool   `json:"review_disabled"`
	ReviewState        string `json:"review_state"`
	ReviewStateID      uint   `json:"review_state_id"`
	ReviewBillable     bool   `json:"review_billable"`
	PreparerID         uint   `json:"preparer_id"`
	ReviewerID         uint   `json:"reviewer_id"`
	TranslBilling      bool   `json:"transl_billing"`
}

type CustomerPriceResponseDTO struct {
	ID         uint
	CustomerID uint    `json:"customer_id"`
	TaskTypeID uint    `json:"task_type"`
	SourceLang string  `json:"source_lang"`
	TargetLang string  `json:"target_lang"`
	CurrencyID uint    `json:"currency"`
	Default    bool    `json:"default"`
	Price      float64 `json:"price"`
}

type AddressResponseDTO struct {
	ID          uint
	CountryCode string `json:"country_code"`
	PostCode    string `json:"post_code"`
	City        string `json:"city"`
	Address     string `json:"address"`
}

type ContactRequestDTO struct {
	ID         uint
	CustomerID uint   `json:"customer_id"`
	Default    bool   `json:"default"`
	Salutation string `json:"salutation"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Position   string `json:"position"`
	Email      string `json:"email"`
	Email2     string `json:"email2"`
	Phone      string `json:"phone"`
	Phone2     string `json:"phone2"`
	Phone3     string `json:"phone3"`
	Phone4     string `json:"phone4"`
	Fax        string `json:"fax"`
	Comment    string `json:"comment"`
}

type ProjectRequestDTO struct {
	ID                      uint
	CustomerID              uint                     `json:"customer_id"`
	ClientOfferID           uint                     `json:"client_offer_id"`
	ExternalProjectCode     string                   `json:"external_project_code"`
	PurchaseOrder           string                   `json:"purchase_order"`
	ProjectArrivalDate      time.Time                `json:"project_arrival_date"`
	ProjectDeadlineDateDate string                   `json:"project_deadline_date_date"`
	ProjectDeadlineDateHour int                      `json:"project_deadline_date_hour"`
	ProjectDeadlineDate     time.Time                `json:"project_deadline_date"`
	FixedPrice              float64                  `json:"fixed_price"`
	SumToBill               float64                  `json:"sum_to_bill"`
	CurrencyName            string                   `json:"currency_name"`
	Currency                globalconstants.Currency `json:"currency"`
	ProjectState            string                   `json:"project_state"`
	ProjectStateID          uint                     `json:"project_state_id"`
	ProjectTimeState        string                   `json:"project_time_state"`
	ProjectTimeStateID      uint                     `json:"project_time_state_id"`
	FulfillmentDateDate     string                   `json:"fulfillment_date_date"`
	FulfillmentDate         time.Time                `json:"fulfillment_date"`
	FirmName                string                   `json:"firm_name"`
	Firm                    uint                     `json:"firm_id"`
	ProjectManagerUserName  string                   `json:"project_manager_user_name"`
	ProjectManager          uint                     `json:"project_manager_id"`
	PartiallyBillable       bool                     `json:"partially_billable"`
	Tasks                   []TaskRequestDTO         `json:"tasks"`
}

type DefaultCustomerPriceRequestDTO struct {
	ID         uint
	TaskTypeID uint                     `json:"task_type"`
	SourceLang string                   `json:"source_lang"`
	TargetLang string                   `json:"target_lang"`
	Currency   globalconstants.Currency `json:"currency"`
	Price      float64                  `json:"price"`
}

type RoleResponseDTO struct {
	ID          uint
	RoleName    string                  `json:"role_name"`
	Permissions []PermissionResponseDTO `json:"permissions"`
	UserID      uint                    `json:"user_id"`
}

type ContactResponseDTO struct {
	ID         uint
	CustomerID uint   `json:"customer_id"`
	Default    bool   `json:"default"`
	Salutation string `json:"salutation"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Position   string `json:"position"`
	Email      string `json:"email"`
	Email2     string `json:"email2"`
	Phone      string `json:"phone"`
	Phone2     string `json:"phone2"`
	Phone3     string `json:"phone3"`
	Phone4     string `json:"phone4"`
	Fax        string `json:"fax"`
	Comment    string `json:"comment"`
}

type CustomerResponseDTO struct {
	ID                  uint
	Name                string                      `json:"name"`
	Address             AddressResponseDTO          `json:"address"`
	Prices              []CustomerPriceResponseDTO  `json:"prices"`
	BillPID             int64                       `json:"id"`
	Email               string                      `json:"email"`
	TaxCode             string                      `json:"taxcode"`
	Iban                string                      `json:"iban"`
	Swift               string                      `json:"swift"`
	AccountNumber       string                      `json:"account_number"`
	Phone               string                      `json:"phone"`
	GeneralLedgerNumber string                      `json:"general_ledger_number"`
	TaxType             string                      `json:"tax_type"`
	CustGeneralInfo     string                      `json:"cust_general_info"`
	CurrencyName        string                      `json:"currency_name"`
	Currency            globalconstants.Currency    `json:"currency"`
	DefaultFirm         uint                        `json:"default_firm"`
	InvoiceLangName     string                      `json:"invoice_language_name"`
	InvoiceLang         globalconstants.InvoiceLang `json:"invoice_language"`
	PaymentDueDays      int                         `json:"payment_due_days"`
}

type TaskCustomerPropsResponseDTO struct {
	ID                  uint
	CustomerID          uint      `json:"customer_id"`
	CustomerDueDate     time.Time `json:"customer_due_date"`
	CustQuantity        float64   `json:"customer_quantity"`
	CustomerPrice       float64   `json:"customer_price"`
	CustomerPriceType   string    `json:"customer_price_type"`
	CustomerPriceTypeID uint      `json:"customer_price_type_id"`
	CustomerUnitType    string    `json:"customer_unit_type"`
	CustomerUnitTypeID  uint      `json:"customer_unit_type_id"`
	CustomerUniquePrice bool      `json:"customer_unique_price"`
}

type DefaultCustomerPriceResponseDTO struct {
	ID         uint
	TaskTypeID uint                     `json:"task_type"`
	SourceLang string                   `json:"source_lang"`
	TargetLang string                   `json:"target_lang"`
	Currency   globalconstants.Currency `json:"currency"`
	Price      float64                  `json:"price"`
}

type PermissionRequestDTO struct {
	ID             uint
	PermissionName string `json:"permission_name"`
}

type UserResponseDTO struct {
	ID                      uint
	Name                    string                     `json:"name"`
	UserName                string                     `json:"user_name"`
	Email                   string                     `json:"email"`
	Active                  bool                       `json:"active"`
	Roles                   []RoleResponseDTO          `json:"roles"`
	Car                     bool                       `json:"car"`
	PrimaryPhoneNumber      string                     `json:"primary_phone_number"`
	SecondaryPhoneNumber    string                     `json:"secondary_phone_number"`
	VendorInvoiceFolderLink string                     `json:"vendor_invoice_folder_link"`
	BillingCurrencyName     string                     `json:"billing_currency_name"`
	BillingCurrency         globalconstants.Currency   `json:"billing_currency"`
	Prices                  []SupplierPriceResponseDTO `json:"prices"`
	Specialities            string                     `json:"specialities"`
	Languages               string                     `json:"languages"`
}

type FirmRequestDTO struct {
	ID          uint
	Name        string            `json:"name"`
	FirmAddress AddressRequestDTO `json:"address"`
	MainEmail   string            `json:"main_email"`
}

type UserRequestDTO struct {
	ID                      uint
	Name                    string                    `json:"name"`
	UserName                string                    `json:"user_name"`
	Email                   string                    `json:"email"`
	Active                  bool                      `json:"active"`
	Roles                   []RoleRequestDTO          `json:"roles"`
	Car                     bool                      `json:"car"`
	PrimaryPhoneNumber      string                    `json:"primary_phone_number"`
	SecondaryPhoneNumber    string                    `json:"secondary_phone_number"`
	VendorInvoiceFolderLink string                    `json:"vendor_invoice_folder_link"`
	BillingCurrencyName     string                    `json:"billing_currency_name"`
	BillingCurrency         globalconstants.Currency  `json:"billing_currency"`
	Prices                  []SupplierPriceRequestDTO `json:"prices"`
	Specialities            string                    `json:"specialities"`
	Languages               string                    `json:"languages"`
}

type RoleRequestDTO struct {
	ID          uint
	RoleName    string                 `json:"role_name"`
	Permissions []PermissionRequestDTO `json:"permissions"`
	UserID      uint                   `json:"user_id"`
}
