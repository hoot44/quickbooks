package quickbooks

type TNameValue struct {
	Value string `json:"value"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}

type TValue struct {
	Value string
}

type TEmail struct {
	Address string
}

type TFreeFormNumber struct {
	FreeFormNumber string
}

type TDate struct {
	Date string
}

type TMetaData struct {
	CreateTime      string
	LastUpdatedTime string
}

type TWebAddress struct {
	URI string
}

type TAddress struct {
	Id                     string
	PostalCode             string
	City                   string
	Country                string
	Line1                  string
	Line2                  string
	Line3                  string
	Line4                  string
	Line5                  string
	Lat                    string
	Long                   string
	CountrySubDivisionCode string
}

type TMarkupRef struct {
	PriceLevelRef          *TNameValue `json:",omitempty"`
	Percent                float64     `json:",omitempty"`
	MarkUpIncomeAccountRef *TNameValue `json:",omitempty"`
}

type TDateTime struct {
	DateTime string `json:",omitempty"`
}
type TSalesItemLineDetail struct {
	TaxInclusiveAmt      float64     `json:",omitempty"`
	DiscountAmt          float64     `json:",omitempty"`
	ItemRef              *TNameValue `json:",omitempty"`
	ClassRef             *TNameValue `json:",omitempty"`
	TaxCodeRef           *TNameValue `json:",omitempty"`
	MarkupInfo           *TMarkupRef `json:",omitempty"`
	ItemAccountRef       *TNameValue `json:",omitempty"`
	ServiceDate          string      `json:",omitempty"`
	DiscountRate         float64     `json:",omitempty"`
	Qty                  float64     `json:",omitempty"`
	UnitPrice            float64     `json:",omitempty"`
	TaxClassificationRef *TNameValue `json:",omitempty"`
}

type TLinkedTxn struct {
	TxnId     string `json:",omitempty"`
	TxnType   string `json:",omitempty"`
	TxnLineId string `json:",omitempty"`
}

type TTxnTaxDetail struct {
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
}

type TCustomField struct {
	DefinitionId string
	StringValue  string `json:",omitempty"`
	Name         string `json:",omitempty"`
	Type         string `json:",omitempty"`
}
