package quickbooks

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SalesItemLine struct {
	Id                  string `json:",omitempty"`
	DetailType          string
	SalesItemLineDetail *TSalesItemLineDetail
	Amount              float64
	Description         string `json:",omitempty"`
	LineNum             uint64 `json:",omitempty"`
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
			Description string `json:",omitempty"`
			LineNum     uint64 `json:",omitempty"`
		}
	}
	DetailType  string
	LineNum     uint64 `json:",omitempty"`
	Description string `json:",omitempty"`
}

type DescriptionOnlyLine struct {
	Id                    string `json:",omitempty"`
	DetailType            string
	DescriptionLineDetail struct {
		TaxCodeRef  *TNameValue `json:",omitempty"`
		ServiceDate TDate       `json:",omitempty"`
	}
	Description string `json:",omitempty"`
	LineNum     uint64 `json:",omitempty"`
	Amount      float64
}

type SubTotalLine struct {
	Id                 string `json:",omitempty"`
	SubTotalLineDetail *struct {
		ItemRef *TNameValue
	} `json:",omitempty"`
	DetailType  string
	Amount      float64
	Description string `json:",omitempty"`
	LineNum     uint64 `json:",omitempty"`
}

type DiscountLine struct {
	Id                 string `json:",omitempty"`
	DiscountLineDetail *struct {
		ClassRef           *TNameValue `json:",omitempty"`
		TaxCodeRef         *TNameValue `json:",omitempty"`
		DiscountAccountRef *TNameValue `json:",omitempty"`
		PercentBased       bool
		DiscountPercent    float64
	}
	DetailType  string
	Amount      float64
	Description string
	LineNum     uint64
}

type InvoiceLineData interface{}

type InvoiceLine struct {
	*DescriptionOnlyLine
	*GroupLine
	*SalesItemLine
	*SubTotalLine
	*DiscountLine
}

func (i *InvoiceLine) UnmarshalJSON(d []byte) error {
	if bytes.Compare(d, []byte(`null`)) == 0 {
		return nil
	}
	m := map[string]interface{}{}
	if err := json.Unmarshal(d, &m); err == nil {
		if detailType, ok := m["DetailType"]; ok {
			switch detailType {
			case "DescriptionOnlyLine":
				i.DescriptionOnlyLine = &DescriptionOnlyLine{}
				return json.Unmarshal(d, &i.DescriptionOnlyLine)
			case "GroupLine":
				i.GroupLine = &GroupLine{}
				return json.Unmarshal(d, &i.GroupLine)
			case "SalesItemLineDetail":
				i.SalesItemLine = &SalesItemLine{}
				return json.Unmarshal(d, &i.SalesItemLine)
			case "SubTotalLineDetail":
				i.SubTotalLine = &SubTotalLine{}
				return json.Unmarshal(d, &i.SubTotalLine)
			default:
				return errors.New(fmt.Sprintf("unknown invoice line type from data: %s", detailType))
			}
		}
	}

	return errors.New(fmt.Sprintf("unknown invoice line type from data: %s", d))
}

func (i *InvoiceLine) MarshalJSON() ([]byte, error) {
	if i.DescriptionOnlyLine != nil {
		i.SalesItemLine.DetailType = "DescriptionOnly"
		return json.Marshal(i.DescriptionOnlyLine)
	} else if i.GroupLine != nil {
		i.SalesItemLine.DetailType = "GroupLineDetail"
		return json.Marshal(i.GroupLine)
	} else if i.SalesItemLine != nil {
		i.SalesItemLine.DetailType = "SalesItemLineDetail"
		return json.Marshal(i.SalesItemLine)
	} else if i.SubTotalLine != nil {
		return json.Marshal(i.SubTotalLine)
	} else if i.DiscountLine != nil {
		return json.Marshal(i.DiscountLine)
	}
	return []byte("null"), nil
}

var _ = InvoiceLineData(SalesItemLine{})
var _ = InvoiceLineData(SalesItemLine{})
var _ = InvoiceLineData(DescriptionOnlyLine{})

type GlobalTaxCalculationId string

const (
	TAX_CALCULATION_TAX_EXCLUDED   GlobalTaxCalculationId = "TaxExcluded"
	TAX_CALCULATION_TAX_INCLUSIVE                         = "TaxInclusive"
	TAX_CALCULATION_NOT_APPLICABLE                        = "NotApplicable"
)

type Invoice struct {
	Id                           string `json:",omitempty"`
	Line                         *[]InvoiceLine
	CustomerRef                  TNameValue
	SyncToken                    string
	CurrencyRef                  *TNameValue
	DocNumber                    string                  `json:",omitempty"`
	BillEmail                    *TEmail                 `json:",omitempty"`
	TxnDate                      string                  `json:",omitempty"`
	ShipFromAddress              *TAddress               `json:",omitempty"`
	ShipDate                     *TDate                  `json:",omitempty"`
	TrackingNum                  string                  `json:",omitempty"`
	ClassRef                     *TNameValue             `json:",omitempty"`
	PrintStatus                  string                  `json:",omitempty"`
	SalesTermRef                 *TNameValue             `json:",omitempty"`
	TxnSource                    string                  `json:",omitempty"`
	LinkedTxn                    []*TLinkedTxn           `json:",omitempty"`
	DepositToAccountRef          *TNameValue             `json:",omitempty"`
	GlobalTaxCalculationEnum     *GlobalTaxCalculationId `json:",omitempty"`
	AllowOnlineACHPayment        bool                    `json:",omitempty"`
	TransactionLocationType      string                  `json:",omitempty"`
	DueDate                      string                  `json:",omitempty"`
	MetaData                     *TMetaData              `json:",omitempty"`
	PrivateNote                  string                  `json:",omitempty"`
	BillEmailCc                  *TEmail                 `json:",omitempty"`
	CustomerMemo                 *TValue                 `json:",omitempty"`
	EmailStatus                  string                  `json:",omitempty"`
	ExchangeRate                 float64                 `json:",omitempty"`
	Deposit                      float64                 `json:",omitempty"`
	TxnTaxDetail                 *TTxnTaxDetail          `json:",omitempty"`
	AllowOnlineCreditCardPayment bool                    `json:",omitempty"`
	CustomField                  []*TCustomField         `json:",omitempty"`
	ShipAddr                     *TAddress               `json:",omitempty"`
	DepartmentRef                *TNameValue             `json:",omitempty"`
	BillEmailBcc                 *TEmail                 `json:",omitempty"`
	ShipMethodRef                *TNameValue             `json:",omitempty"`
	BillAddr                     *TAddress               `json:",omitempty"`
	ApplyTaxAfterDiscount        bool                    `json:",omitempty"`
	HomeBalance                  float64                 `json:",omitempty"`
	DeliveryInfo                 *struct {
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
