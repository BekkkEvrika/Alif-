package base

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
)

var db *gorm.DB

var MySetting *Setting

var RabbitConnection string

// Connect ...
func Connect() error {
	var err error

	MySetting = new(Setting)
	if MySetting, err = readSettings(); err != nil {
		return err
	}
	dsn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Shanghai",
		MySetting.DBHost, MySetting.DBUser, MySetting.DBPassword, MySetting.DBName, MySetting.DBPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func Migrate(v interface{}) error {
	if !db.Migrator().HasTable(v) {
		err := db.Migrator().CreateTable(v)
		return err
	}
	return db.Error
}

func readSettings() (*Setting, error) {
	content, err := ioutil.ReadFile("setting.json")
	if err != nil {
		return nil, err
	}
	s := Setting{}
	if err := json.Unmarshal([]byte(content), &s); err != nil {
		return nil, err
	}
	return &s, nil
}
