package quickbooks

type CustomerPreferredDeliveryMethod string
type CustomerTaxExemptionReasonId int

const (
	PRINT CustomerPreferredDeliveryMethod = "Print"
	EMAIL                                 = "Email"
	NONE                                  = "None"
)

const (
	FEDERAL_GOVERNMENT CustomerTaxExemptionReasonId = 1
	STATE_GOVERNMENT
	LOCAL_GOVERNMENT
	TRIBAL_GOVERNMENT
	CHARITABLE_ORGANIZATION
	RELIGIOUS_ORGANIZATION
	EDUCATIONAL_ORGANIZATION
	HOSPITAL
	RESALE
	DIRECT_PAY_PERMIT
	MULTIPLE_POINTS_OF_USE
	DIRECT_MAIL
	AGRICULTURAL_PRODUCTION
	INDUSTRIAL_PRODUCTION_MANUFACTURING
	FOREIGN_DIPLOMAT
)

type CreateCustomer struct {
	DisplayName string
	Suffix      string
	Title       string
	MiddleName  string
	FamilyName  string
	GivenName   string
}

type Customer struct {
	Id                      string
	SyncToken               string
	DisplayName             string                           `json:",omitempty"`
	Title                   string                           `json:",omitempty"`
	GivenName               string                           `json:",omitempty"`
	MiddleName              string                           `json:",omitempty"`
	Suffix                  string                           `json:",omitempty"`
	FamilyName              string                           `json:",omitempty"`
	PrimaryEmailAddr        *TEmail                          `json:",omitempty"`
	ResaleNum               string                           `json:",omitempty"`
	SecondaryTaxIdentifier  string                           `json:",omitempty"`
	ARAccountRef            *TNameValue                      `json:",omitempty"`
	DefaultTaxCodeRef       *TNameValue                      `json:",omitempty"`
	PreferredDeliveryMethod *CustomerPreferredDeliveryMethod `json:",omitempty"`
	GSTIN                   string                           `json:",omitempty"`
	SalesTermRef            *TNameValue                      `json:",omitempty"`
	CustomerTypeRef         *TValue                          `json:",omitempty"`
	Fax                     *TFreeFormNumber                 `json:",omitempty"`
	BusinessNumber          string                           `json:",omitempty"`
	BillWithParent          bool                             `json:",omitempty"`
	CurrencyRef             *TNameValue                      `json:",omitempty"`
	Mobile                  *TFreeFormNumber                 `json:",omitempty"`
	Job                     bool                             `json:",omitempty"`
	BalanceWithJobs         float64                          `json:",omitempty"`
	PrimaryPhone            *TFreeFormNumber                 `json:",omitempty"`
	OpenBalanceDate         *TDate                           `json:",omitempty"`
	Taxable                 bool                             `json:",omitempty"`
	AlternatePhone          *TFreeFormNumber                 `json:",omitempty"`
	MetaData                *TMetaData                       `json:",omitempty"`
	ParentRef               *TNameValue                      `json:",omitempty"`
	Notes                   string                           `json:",omitempty"`
	WebAddr                 *TWebAddress                     `json:",omitempty"`
	Active                  bool                             `json:",omitempty"`
	CompanyName             string                           `json:",omitempty"`
	Balance                 float64                          `json:",omitempty"`
	ShipAddr                *TAddress                        `json:",omitempty"`
	PaymentMethodRef        *TNameValue                      `json:",omitempty"`
	IsProject               bool                             `json:",omitempty"`
	Source                  string                           `json:",omitempty"`
	PrimaryTaxIdentifier    string                           `json:",omitempty"`
	GSTRegistrationType     string                           `json:",omitempty"`
	PrintOnCheckName        string                           `json:",omitempty"`
	BillAddr                *TAddress                        `json:",omitempty"`
	FullyQualifiedName      string                           `json:",omitempty"`
	Level                   int64                            `json:",omitempty"`
	TaxExemptionReasonId    *CustomerTaxExemptionReasonId    `json:",omitempty"`
	rt                      *RefreshToken
}
