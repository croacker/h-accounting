package persist

import (
	"../ofd"
	"gopkg.in/mgo.v2/bson"
)

//CheckTotal Итоги чека
type CheckTotal struct {
	Id                   bson.ObjectId `bson:"_id"`
	ShopId               bson.ObjectId
	DateTime             int
	CashTotalSum         int
	Discount             interface{}
	DiscountSum          interface{}
	EcashTotalSum        int
	FiscalDocumentNumber int
	FiscalDriveNumber    string
	FiscalSign           int64
	KktNumber            interface{}
	KktRegID             string
	Markup               interface{}
	MarkupSum            interface{}
	Modifiers            interface{}
	Nds0                 interface{}
	Nds10                int
	Nds18                int
	NdsCalculated10      interface{}
	NdsCalculated18      interface{}
	NdsNo                interface{}
	OperationType        int
	Operator             string
	RequestNumber        int
	RetailPlaceAddress   interface{}
	ShiftNumber          int
	StornoItems          interface{}
	TaxationType         int
	TotalSum             int
}

//Фабричный метод
func NewCheckTotal(shop *Shop, check *ofd.OfdCheck) *CheckTotal {
	return &CheckTotal{bson.NewObjectId(),
		shop.Id,
		check.DateTime,
		check.CashTotalSum,
		check.Discount,
		check.DiscountSum,
		check.EcashTotalSum,
		check.FiscalDocumentNumber,
		check.FiscalDriveNumber,
		check.FiscalSign,
		check.KktNumber,
		check.KktRegID,
		check.Markup,
		check.MarkupSum,
		check.Modifiers,
		check.Nds0,
		check.Nds10,
		check.Nds18,
		check.NdsCalculated10,
		check.NdsCalculated18,
		check.NdsNo,
		check.OperationType,
		check.Operator,
		check.RequestNumber,
		check.RetailPlaceAddress,
		check.ShiftNumber,
		check.StornoItems,
		check.TaxationType,
		check.TotalSum,
	}
}
