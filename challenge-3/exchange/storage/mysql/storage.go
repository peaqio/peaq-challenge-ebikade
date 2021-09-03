package storage

import (
	"fmt"
	"os"

	"github.com/ebikode/peaq-challenge/challenge-3/exchange/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// EDatabase for database connection and operations
type EDatabase struct {
	db *gorm.DB
}

// Config for application configs
type Config struct {
}

// New ...
func New() *Config {
	return &Config{}
}

// InitDB ..
func (config *Config) InitDB() (*EDatabase, error) {

	var err error
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	charset := "utf8mb4"

	// dbHost = "127.0.0.1"
	// dbPort = "3306"
	// dbName = "peaq_analytics_goDB"
	// dbUser = "root"
	// dbPass = ""

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True", dbUser, dbPass, dbHost, dbPort, dbName, charset) //Build connection string
	fmt.Println(dbURI)

	edb := new(EDatabase)

	edb.db, err = gorm.Open(mysql.New(
		mysql.Config{
			DSN: dbURI,
		}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		fmt.Print(err)
		return edb, err
	}

	// edb.db.Migrator().DropTable(
	// 	&models.Rate{},
	// 	&models.GrowthRecord{},
	// )

	// Migrating tables to database
	edb.db.Migrator().AutoMigrate(
		&models.Rate{},
		&models.GrowthRecord{},
	) //Database migration

	// defer mdb.db.Close()

	return edb, nil

}
