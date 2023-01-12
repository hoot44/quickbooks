package quickbooks

type TNameValue struct {
	Value string `json:"value,omitempty"`
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
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
