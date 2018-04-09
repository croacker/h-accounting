package persistsql

import (
	"strings"

	"../ofd"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

// Save - Сохранить чеки в хранилище
func Save(checks *ofd.OfdChecks) error {
	var err error
	db := GetDb()
	defer db.Close()

	sailerDao := SailerDao{db}
	shopDao := ShopDao{db}
	productDao := ProductDao{db}
	productCathegoryDao := ProductCathegoryDao{db}
	priceDao := PriceDao{db}
	checkHeaderDao := CheckHeaderDao{db}

	commonCathegory := productCathegoryDao.FirstOrCreate("Общие")

	for _, check := range *checks {
		sailerName := strings.Trim(check.User, " ")
		inn := strings.Trim(check.UserInn, " ")
		sailer := sailerDao.FirstOrCreate(sailerName, inn)

		shop := shopDao.FirstOrCreate(sailer, check.RetailPlaceAddress)

		for _, item := range check.Items {
			productName := strings.Trim(item.Name, " ")
			product := productDao.FirstOrCreate(productName, commonCathegory)

			price := &Price{
				Sailer:   *sailer,
				Product:  *product,
				Price:    item.Price,
				DateTime: check.DateTime,
			}
			priceDao.FirstOrCreate(price)
		}

		checkHeader := checkHeaderDao.NewCheckHeader(*shop, &check)
		checkHeaderDao.FirstOrCreate(checkHeader)
	}

	return err
}

func Init() {
	db := GetDb()
	migrate(db)
	defer db.Close()
}

func ProductsList() []Product {
	db := GetDb()
	dao := ProductDao{db}
	return dao.GetAll()
}

func PricesList() []Price {
	db := GetDb()
	dao := PriceDao{db}
	return dao.GetAll()
}

func SailersList() []Sailer {
	db := GetDb()
	dao := SailerDao{db}
	return dao.GetAll()
}

func ShopsList() []Shop {
	db := GetDb()
	dao := ShopDao{db}
	return dao.GetAll()
}

func ChecksList() []CheckHeader {
	db := GetDb()
	dao := CheckHeaderDao{db}
	return dao.GetAll()
}

func GetDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./db/h-accounting.db")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&ProductCathegory{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Sailer{})
	db.AutoMigrate(&Shop{})
	db.AutoMigrate(&Price{})
	db.AutoMigrate(&CheckHeader{})
	db.AutoMigrate(&CheckItem{})
}
