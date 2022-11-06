package configs

import (
	"github.com/seiyamiyaoka/crm_backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NOTE: https://stackoverflow.com/questions/57928208/how-to-make-multiple-models-auto-migrate-in-gorm
var AllModels = []interface{}{
	&models.Customer{},
	&models.Contact{},
}

func ConnectDB() *gorm.DB {
	dsn := "crm_gackend_user:password@tcp(db:3306)/crm_gackend_database?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(AllModels...)
}

var DB *gorm.DB = ConnectDB()

func GetCollection(client *gorm.DB, collectionName string) *gorm.DB {
	return client.Table(collectionName)
}

func SetUpSeedData() {
	customers := []models.Customer{
		{
			Name:      "John1",
			Role:      "CEO",
			Email:     "fasdfa@examle.com",
			Phone:     "1234567890",
			Contacted: false,
		},
		{
			Name:      "Jane",
			Role:      "CTO",
			Email:     "afsdfa@example.com",
			Phone:     "1234567890",
			Contacted: false,
		},
		{
			Name:      "Jane",
			Role:      "CTO",
			Email:     "afsdfa@example.com",
			Phone:     "1234567890",
			Contacted: false,
		},
	}

	DB.Create(&customers)
}
