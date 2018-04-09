package persistsql

import (
	"fmt"
	"strings"

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
	// RetailPlaceAddress   string
	ShiftNumber  int
	StornoItems  string
	TaxationType int
	TotalSum     int
	Shop         Shop `gorm:"foreignkey:ShopId"`
	ShopId       uint
	// User                 string
	// UserInn              string
	Items []CheckItem
}

type CheckItem struct {
	gorm.Model
	Modifiers       string
	Product         Product `gorm:"foreignkey:ProductId"`
	ProductId       uint
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

func (dao CheckHeaderDao) NewCheckHeader(shop Shop, ofdCheck *ofd.OfdCheck) *CheckHeader {
	return &CheckHeader{
		CashTotalSum:         ofdCheck.CashTotalSum,
		DateTime:             ofdCheck.DateTime,
		Discount:             ofd.ToInt(ofdCheck.Discount),
		DiscountSum:          ofd.ToInt(ofdCheck.DiscountSum),
		EcashTotalSum:        ofdCheck.EcashTotalSum,
		FiscalDocumentNumber: ofdCheck.FiscalDocumentNumber,
		FiscalDriveNumber:    strings.Trim(ofdCheck.FiscalDriveNumber, " "),
		FiscalSign:           ofdCheck.FiscalSign,
		KktNumber:            ofd.ToString(ofdCheck.KktNumber),
		KktRegID:             strings.Trim(ofdCheck.KktRegID, " "),
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
		Operator:             strings.Trim(ofdCheck.Operator, " "),
		RequestNumber:        ofdCheck.RequestNumber,
		// RetailPlaceAddress:   ofd.ToString(ofdCheck.RetailPlaceAddress),
		ShiftNumber:  ofdCheck.ShiftNumber,
		StornoItems:  ofd.ToString(ofdCheck.StornoItems),
		TaxationType: ofdCheck.TaxationType,
		TotalSum:     ofdCheck.TotalSum,
		// User:                 strings.Trim(ofdCheck.User, " "),
		// UserInn:              strings.Trim(ofdCheck.UserInn, " "),
		Shop:  shop,
		Items: dao.NewCheckItems(ofdCheck),
	}
}

func (dao CheckHeaderDao) NewCheckItems(ofdCheck *ofd.OfdCheck) []CheckItem {
	productDao := ProductDao{dao.db}
	items := make([]CheckItem, 0)
	for _, ofdItem := range ofdCheck.Items {
		product := productDao.FindByName(ofdItem.Name)
		item := CheckItem{
			Modifiers:       ofd.ToString(ofdItem.Modifiers),
			Product:         *product,
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
	dao.db.FirstOrCreate(сheckHeader,
		CheckHeader{DateTime: сheckHeader.DateTime,
			FiscalDriveNumber: сheckHeader.FiscalDriveNumber,
			FiscalSign:        сheckHeader.FiscalSign,
			KktRegID:          сheckHeader.KktRegID,
		})
	fmt.Println("Check Id:", сheckHeader.ID)
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

func (dao CheckHeaderDao) GetAll() []CheckHeader {
	var checkHeaders []CheckHeader
	dao.db.Preload("Shop").Find(&checkHeaders)
	return checkHeaders
}
