package db

import (
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"webapp/db/models"
)

type Item struct {
	URL   string  `json:"url"`
	PRICE float64 `json:"price"`
}

var (
	db    *gorm.DB
	err   error
	mutex sync.Mutex
)

func Init(dsn string) error {

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.Product{})

	return err
}

func CreateItem(url string, price float64) error {
	mutex.Lock()
	newProduct := models.Product{Url: url, Price: price, AddedAt: time.Now()}
	res := db.Create(&newProduct)

	if res.Error != nil {
		return res.Error
	}
	mutex.Unlock()

	return nil
}

func GetItems() ([]Item, error) {
	mutex.Lock()
	var products []models.Product

	var productsList []Item

	res := db.Find(&products)

	if res.Error != nil {
		return nil, res.Error
	}

	for _, product := range products {
		productsList = append(productsList, Item{URL: product.Url, PRICE: product.Price})
	}
	defer mutex.Unlock()

	return productsList, nil
}

func GetItem(url string) (Item, error) {
	mutex.Lock()
	var product models.Product

	res := db.First(&product, "url = ?", url)

	if res.Error != nil {
		return Item{}, res.Error
	}
	defer mutex.Unlock()

	return Item{URL: product.Url, PRICE: product.Price}, nil
}

func EditItem(url string, price float64) error {
	mutex.Lock()
	res := db.Model(&models.Product{}).Where("Url", url).Update("Price", price)
	defer mutex.Unlock()

	return res.Error
}

func DeleteItem(url string) error {
	mutex.Lock()
	res := db.Delete(&models.Product{}, "Url", url)
	defer mutex.Unlock()

	return res.Error
}
