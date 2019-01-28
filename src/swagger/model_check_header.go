/*
 * OFD Checks API
 *
 * This is a OFD Checks API
 *
 * API version: 1.0.0
 * Contact: croacker@homeworld.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type CheckHeader struct {

	Shop *Shop `json:"shop"`

	OperationType string `json:"operationType,omitempty"`

	Operator string `json:"operator,omitempty"`

	RequestNumber string `json:"requestNumber,omitempty"`

	DateTime time.Time `json:"dateTime,omitempty"`

	ShiftNumber string `json:"shiftNumber,omitempty"`

	FiscalDocumentNumber string `json:"fiscalDocumentNumber,omitempty"`

	FiscalDriveNumber string `json:"fiscalDriveNumber,omitempty"`

	FiscalSign string `json:"fiscalSign,omitempty"`

	KktNumber string `json:"kktNumber,omitempty"`

	KktRegID string `json:"kktRegID,omitempty"`

	CashTotalSum float32 `json:"cashTotalSum,omitempty"`

	EcashTotalSum float32 `json:"ecashTotalSum,omitempty"`

	TotalSum float32 `json:"totalSum,omitempty"`

	Discount float32 `json:"discount,omitempty"`

	DiscountSum float32 `json:"discountSum,omitempty"`

	Markup string `json:"markup,omitempty"`

	MarkupSum float32 `json:"markupSum,omitempty"`

	Modifiers string `json:"modifiers,omitempty"`

	Nds0 float32 `json:"nds0,omitempty"`

	Nds10 float32 `json:"nds10,omitempty"`

	Nds18 float32 `json:"nds18,omitempty"`

	NdsCalculated10 float32 `json:"ndsCalculated10,omitempty"`

	NdsCalculated18 float32 `json:"ndsCalculated18,omitempty"`

	NdsNo float32 `json:"ndsNo,omitempty"`

	TaxationType string `json:"taxationType,omitempty"`

	StornoItems string `json:"stornoItems,omitempty"`
}