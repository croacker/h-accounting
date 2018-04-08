package persistsql

import (
	"fmt"

	"../ofd"
	"github.com/jinzhu/gorm"
)

//CheckHeader Заголовок чека
type CheckHeader struct {
	gorm.Model
	CashTotalSum         int
	DateTime             int
	Discount             int
	DiscountSum          int
	EcashTotalSum        int
	FiscalDocumentNumber int
	FiscalDriveNumber    string
	FiscalSign           int64
	KktNumber            string
	KktRegID             string
	Markup               string
	MarkupSum            int
	Modifiers            int
	Nds0                 int
	Nds10                int
	Nds18                int
	NdsCalculated10      int
	NdsCalculated18      int
	NdsNo                int
	OperationType        int
	Operator             string
	RequestNumber        int
	RetailPlaceAddress   string
	ShiftNumber          int
	StornoItems          string
	TaxationType         int
	TotalSum             int
	User                 string
	UserInn              string
	Items                []CheckItem
}

type CheckItem struct {
	gorm.Model
	Modifiers       string
	Name            string
	Nds0            int
	Nds10           int
	Nds18           int
	NdsCalculated10 int
	NdsCalculated18 int
	NdsNo           int
	Price           int
	Quantity        float64
	Sum             int
	Storno          bool
	CheckHeaderID   uint
}

func NewCheckHeader(ofdCheck *ofd.OfdCheck) *CheckHeader {
	return &CheckHeader{
		CashTotalSum:         ofdCheck.CashTotalSum,
		DateTime:             ofdCheck.DateTime,
		Discount:             ofd.ToInt(ofdCheck.Discount),
		DiscountSum:          ofd.ToInt(ofdCheck.DiscountSum),
		EcashTotalSum:        ofdCheck.EcashTotalSum,
		FiscalDocumentNumber: ofdCheck.FiscalDocumentNumber,
		FiscalDriveNumber:    ofdCheck.FiscalDriveNumber,
		FiscalSign:           ofdCheck.FiscalSign,
		KktNumber:            ofd.ToString(ofdCheck.KktNumber),
		KktRegID:             ofdCheck.KktRegID,
		Markup:               ofd.ToString(ofdCheck.Markup),
		MarkupSum:            ofd.ToInt(ofdCheck.MarkupSum),
		Modifiers:            ofd.ToInt(ofdCheck.Modifiers),
		Nds0:                 ofd.ToInt(ofdCheck.Nds0),
		Nds10:                ofdCheck.Nds10,
		Nds18:                ofdCheck.Nds18,
		NdsCalculated10:      ofd.ToInt(ofdCheck.NdsCalculated10),
		NdsCalculated18:      ofd.ToInt(ofdCheck.NdsCalculated18),
		NdsNo:                ofd.ToInt(ofdCheck.NdsNo),
		OperationType:        ofdCheck.OperationType,
		Operator:             ofdCheck.Operator,
		RequestNumber:        ofdCheck.RequestNumber,
		RetailPlaceAddress:   ofd.ToString(ofdCheck.RetailPlaceAddress),
		ShiftNumber:          ofdCheck.ShiftNumber,
		StornoItems:          ofd.ToString(ofdCheck.StornoItems),
		TaxationType:         ofdCheck.TaxationType,
		TotalSum:             ofdCheck.TotalSum,
		User:                 ofdCheck.User,
		UserInn:              ofdCheck.UserInn,
		Items:                NewCheckItems(ofdCheck),
	}
}

func NewCheckItems(ofdCheck *ofd.OfdCheck) []CheckItem {
	items := make([]CheckItem, 0)
	for _, ofdItem := range ofdCheck.Items {
		item := CheckItem{
			Modifiers:       ofd.ToString(ofdItem.Modifiers),
			Name:            ofdItem.Name,
			Nds0:            ofd.ToInt(ofdItem.Nds0),
			Nds10:           ofd.ToInt(ofdItem.Nds10),
			Nds18:           ofdItem.Nds18,
			NdsCalculated10: ofd.ToInt(ofdItem.NdsCalculated10),
			NdsCalculated18: ofd.ToInt(ofdItem.NdsCalculated18),
			NdsNo:           ofd.ToInt(ofdItem.NdsNo),
			Price:           ofdItem.Price,
			Quantity:        ofdItem.Quantity,
			Sum:             ofdItem.Sum,
			Storno:          ofdItem.Storno,
		}
		items = append(items, item)
	}
	return items
}

type CheckHeaderDao struct {
	db *gorm.DB
}

func (dao CheckHeaderDao) Create(checkHeader *CheckHeader) {
	dao.db.Create(checkHeader)
}

func (dao CheckHeaderDao) FirstOrCreate(сheckHeader *CheckHeader) *CheckHeader {
	dao.db.FirstOrCreate(&сheckHeader, сheckHeader)
	fmt.Println("Price Id:", сheckHeader.ID)
	return сheckHeader
}

func (dao CheckHeaderDao) Save(сheckHeader *CheckHeader) {
	dao.db.Save(сheckHeader)
}

func (dao CheckHeaderDao) FindById(id uint) *CheckHeader {
	var сheckHeader CheckHeader
	dao.db.First(&сheckHeader, id)
	return &сheckHeader
}
