package quickbooks

type ClientPreferredDeliveryMethod string
type ClientTaxExemptionReasonID int

const (
	PRINT ClientPreferredDeliveryMethod = "Print"
	EMAIL                               = "Email"
	NONE                                = "None"
)

const (
	FEDERAL_GOVERNMENT ClientTaxExemptionReasonID = 1
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

type Client struct {
	ID                     string
	SyncToken              string
	DisplayName            string
	Title                  string
	GivenName              string
	MiddleName             string
	Suffix                 string
	FamilyName             string
	PrimaryEmailAddr       struct{ Address string }
	ResaleNum              string
	SecondaryTaxIdentifier string
	ARAccountRef           struct {
		Value string
		Name  string
	}
	DefaultTaxCodeRef struct {
		Value string
		Name  string
	}
	PreferredDeliveryMethod ClientPreferredDeliveryMethod
	GSTIN                   string
	SalesTermRef            struct {
		Value string
		Name  string
	}
	CustomerTypeRef struct{ Value string }
	Fax             struct{ FreeFormNumber string }
	BusinessNumber  string
	BillWithParent  bool
	CurrencyRef     struct {
		Value string
		Name  string
	}
	Mobile          struct{ FreeFormNumber string }
	Job             bool
	BalanceWithJobs float64
	PrimaryPhone    struct{ FreeFormNumber string }
	OpenBalanceDate struct{ Date string }
	Taxable         bool
	AlternatePhone  struct{ FreeFormNumber string }
	MetaData        struct {
		CreateTime      string
		LastUpdatedTime string
	}
	ParentRef struct {
		Value string
		Name  string
	}
	Notes       string
	WebAddr     struct{ URI string }
	Active      bool
	CompanyName string
	Balance     float64
	ShipAddr    struct {
		ID                                string
		PostalCode                        string
		City                              string
		Country                           string
		Line1, Line2, Line3, Line4, Line5 string
		Lat, Long                         string
		CountrySubDivisionCode            string
	}
	PaymentMethodRef struct {
		Value string
		Name  string
	}
	IsProject            bool
	Source               string
	PrimaryTaxIdentifier string
	GSTRegistrationType  string
	PrintOnCheckName     string
	BillAddr             struct {
		ID                                string
		PostalCode                        string
		City                              string
		Country                           string
		Line1, Line2, Line3, Line4, Line5 string
		Lat, Long                         string
		CountrySubDivisionCode            string
	}
	FullyQualifiedName   string
	Level                int64
	TaxExemptionReasonId *ClientTaxExemptionReasonID
	rt                   *RefreshToken
}
