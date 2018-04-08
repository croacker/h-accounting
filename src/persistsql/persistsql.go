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

	shopDao := ShopDao{db}
	productDao := ProductDao{db}
	productCathegoryDao := ProductCathegoryDao{db}
	priceDao := PriceDao{db}

	commonCathegory := productCathegoryDao.FirstOrCreate("Общие")

	for _, check := range *checks {

		shopName := strings.Trim(check.User, " ")
		inn := strings.Trim(check.UserInn, " ")
		shop := shopDao.FirstOrCreate(shopName, inn)

		for _, item := range check.Items {
			productName := strings.Trim(item.Name, " ")
			product := productDao.FirstOrCreate(productName, commonCathegory)

			price := &Price{
				Shop:     *shop,
				Product:  *product,
				Price:    item.Price,
				DateTime: check.DateTime,
			}
			priceDao.FirstOrCreate(price)
		}
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

func ShopsList() []Shop {
	db := GetDb()
	dao := ShopDao{db}
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
	db.AutoMigrate(&Shop{})
	db.AutoMigrate(&Price{})
	db.AutoMigrate(&CheckHeader{})
	db.AutoMigrate(&CheckItem{})
}
