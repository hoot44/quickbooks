package quickbooks

type InvoiceLine interface{}

type SalesItemLine struct {
	Id                  string `json:",omitempty"`
	DetailType          string
	SalesItemLineDetail *TSalesItemLineDetail
	Amount              float64
	Description         string  `json:",omitempty"`
	LineNum             float64 `json:",omitempty"`
}

type GroupLine struct {
	Id              string
	GroupLineDetail *struct {
		Quantity float64 `json:",omitempty"`
		Line     *[]struct {
			Id                  string `json:",omitempty"`
			DetailType          string
			SalesItemLineDetail *struct {
				TaxInclusiveAmt        float64     `json:",omitempty"`
				DiscountAmt            float64     `json:",omitempty"`
				ItemRef                *TNameValue `json:",omitempty"`
				ClassRef               *TNameValue `json:",omitempty"`
				TaxCodeRef             *TNameValue `json:",omitempty"`
				PriceLevelRef          *TNameValue `json:",omitempty"`
				Percent                float64     `json:",omitempty"`
				MarkUpIncomeAccountRef *TNameValue `json:",omitempty"`
			}
			Amount      float64
			Description string  `json:",omitempty"`
			LineNum     float64 `json:",omitempty"`
		}
	}
	DetailType  string
	LineNum     float64 `json:",omitempty"`
	Description string  `json:",omitempty"`
}

type DescriptionOnlyLine struct {
	Id                    string `json:",omitempty"`
	DetailType            string
	DescriptionLineDetail struct {
		TaxCodeRef  *TNameValue `json:",omitempty"`
		ServiceDate TDate       `json:",omitempty"`
	}
	Description string  `json:",omitempty"`
	LineNum     float64 `json:",omitempty"`
	Amount      float64
}

var _ = InvoiceLine(SalesItemLine{})
var _ = InvoiceLine(SalesItemLine{})
var _ = InvoiceLine(DescriptionOnlyLine{})

type CreateInvoice struct {
	CustomerRef *TNameValue
	Line        *[]InvoiceLine
	CurrencyRef *TNameValue
}

type GlobalTaxCalculationId string

const (
	TAX_CALCULATION_TAX_EXCLUDED   GlobalTaxCalculationId = "TaxExcluded"
	TAX_CALCULATION_TAX_INCLUSIVE                         = "TaxInclusive"
	TAX_CALCULATION_NOT_APPLICABLE                        = "NotApplicable"
)

type Invoice struct {
	Id              string
	Line            *[]InvoiceLine
	CustomerRef     *TNameValue
	SyncToken       string
	CurrencyRef     *TNameValue
	DocNumber       string      `json:",omitempty"`
	BillEmail       *TEmail     `json:",omitempty"`
	TxnDate         string      `json:",omitempty"`
	ShipFromAddress *TAddress   `json:",omitempty"`
	ShipDate        *TDate      `json:",omitempty"`
	TrackingNum     string      `json:",omitempty"`
	ClassRef        *TNameValue `json:",omitempty"`
	PrintStatus     string      `json:",omitempty"`
	SalesTermRef    *TNameValue `json:",omitempty"`
	TxnSource       string      `json:",omitempty"`
	LinkedTxn       []*struct {
		TxnId     string `json:",omitempty"`
		TxnType   string `json:",omitempty"`
		TxnLineId string `json:",omitempty"`
	} `json:",omitempty"`
	DepositToAccountRef      *TNameValue             `json:",omitempty"`
	GlobalTaxCalculationEnum *GlobalTaxCalculationId `json:",omitempty"`
	AllowOnlineACHPayment    bool                    `json:",omitempty"`
	TransactionLocationType  string                  `json:",omitempty"`
	DueDate                  string                  `json:",omitempty"`
	MetaData                 *TMetaData              `json:",omitempty"`
	PrivateNote              string                  `json:",omitempty"`
	BillEmailCc              *TEmail                 `json:",omitempty"`
	CustomerMemo             *TValue                 `json:",omitempty"`
	EmailStatus              string                  `json:",omitempty"`
	ExchangeRate             float64                 `json:",omitempty"`
	Deposit                  float64                 `json:",omitempty"`
	TxnTaxDetail             *struct {
		TxnTaxCodeRef *TNameValue `json:",omitempty"`
		TotalTax      float64     `json:",omitempty"`
		TaxLine       *[]struct {
			DetailType    string
			TaxLineDetail *struct {
				TaxRateRef          *TNameValue
				NetAmountTaxable    float64 `json:",omitempty"`
				PercentBased        bool    `json:",omitempty"`
				TaxInclusiveAmount  float64 `json:",omitempty"`
				OverrideDeltaAmount float64 `json:",omitempty"`
				TaxPercent          float64 `json:",omitempty"`
			}
			Amount float64 `json:",omitempty"`
		}
	} `json:",omitempty"`
	AllowOnlineCreditCardPayment bool `json:",omitempty"`
	CustomField                  []*struct {
		DefinitionId string
		StringValue  string `json:",omitempty"`
		Name         string `json:",omitempty"`
		Type         string `json:",omitempty"`
	}
	ShipAddr              *TAddress   `json:",omitempty"`
	DepartmentRef         *TNameValue `json:",omitempty"`
	BillEmailBcc          *TEmail     `json:",omitempty"`
	ShipMethodRef         *TNameValue `json:",omitempty"`
	BillAddr              *TAddress   `json:",omitempty"`
	ApplyTaxAfterDiscount bool        `json:",omitempty"`
	HomeBalance           float64     `json:",omitempty"`
	DeliveryInfo          *struct {
		DeliveryType string     `json:",omitempty`
		DeliveryTime *TDateTime `json:",omitempty"`
	}
	TotalAmt        float64     `json:",omitempty"`
	InvoiceLink     string      `json:",omitempty"`
	RecurDataRef    *TNameValue `json:",omitempty"`
	TaxExemptionRef *TNameValue `json:",omitempty"`
	Balance         float64     `json:",omitempty"`
	HomeTotalAmt    float64     `json:",omitempty"`
	FreeFormAddress bool        `json:",omitempty"`
}
