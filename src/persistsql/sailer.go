package persistsql

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//Тип продавец(магазин)
type Sailer struct {
	gorm.Model
	//Наименование
	Name string
	//ИНН
	Inn string
}

type SailerDao struct {
	db *gorm.DB
}

func (dao SailerDao) Create(sailer *Sailer) {
	dao.db.Create(sailer)
}

func (dao SailerDao) FirstOrCreate(name string, inn string) *Sailer {
	sailer := &Sailer{Name: name, Inn: inn}
	dao.db.FirstOrCreate(sailer, Sailer{Inn: inn})
	fmt.Println("Sailer Id:", sailer.ID)
	return sailer
}

func (dao SailerDao) Save(sailer *Sailer) {
	dao.db.Save(sailer)
}

func (dao SailerDao) FindById(id uint) *Sailer {
	var sailer Sailer
	dao.db.First(&sailer, id)
	return &sailer
}

func (dao SailerDao) FindByInn(inn string) *Sailer {
	var sailer Sailer
	dao.db.First(&sailer, "inn = ?", inn)
	return &sailer
}

func (dao SailerDao) GetAll() []Sailer {
	var sailers []Sailer
	dao.db.Find(&sailers)
	return sailers
}
