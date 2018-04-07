package persistsql

import (
	"fmt"
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
	migrate(db)
	defer db.Close()

	productDao := ProductDao{db}
	productCathegoryDao := ProductCathegoryDao{db}
	shopDao := ShopDao{db}

	commonCathegory := productCathegoryDao.CreateIfNotExists("Общие")

	for _, check := range *checks {

		shopName := strings.Trim(check.User, " ")
		inn := strings.Trim(check.UserInn, " ")
		_ = shopDao.CreateIfNotExists(shopName, inn)

		for _, item := range check.Items {
			productName := strings.Trim(item.Name, " ")
			_ = productDao.CreateIfNotExists(productName, commonCathegory)
		}
	}

	return err
}

func Connect() {
	db := GetDb()
	migrate(db)
	defer db.Close()

	// Migrate the schema

	cathegory := ProductCathegory{Name: "Фрукты"}
	db.Create(&cathegory)

	// Create
	db.Create(&Product{Name: "Яблоки", Cathegory: cathegory})
	db.Create(&Product{Name: "Бананы", Cathegory: cathegory})

	// Read
	var product Product
	db.First(&product, 1) // find product with id 1
	fmt.Println(product)
	var product2 Product
	db.First(&product2, "name = ?", "Бананы") // find product with code l1212
	fmt.Println(product2)

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

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
}
