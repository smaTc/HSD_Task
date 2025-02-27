package hsdtask

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB = nil

func InitDB(dbPath string) {
	if database != nil {
		return
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database = db
	database.AutoMigrate(Task{})
}

func RemoveDb(path string) {
	database = nil
	os.Remove(path)
}

func CreateData[T any](data T) error {
	tx := database.Create(data)
	err := tx.Error

	if err != nil {
		Logger.Println("db create error: ", err)
		return err
	}

	return nil
}

func ReadAllData[T any](data T) error {
	tx := database.Find(data)
	err := tx.Error

	if err != nil {
		Logger.Println("db read all error: ", err)
		return err
	}

	return nil
}

func ReadByIdData[T any](data T, ids []int) error {
	tx := database.Find(data, ids)
	err := tx.Error

	if err != nil {
		Logger.Println("db id read error: ", err)
		return err
	}

	return nil
}

func UpdateData[T any](data T) error {
	tx := database.Updates(data)
	err := tx.Error

	if err != nil {
		Logger.Println("db update error: ", err)
		return err
	}

	return nil
}

func DeleteData[T any](data T) error {
	tx := database.Delete(data)
	err := tx.Error

	if err != nil {
		Logger.Println("db delete error: ", err)
		return err
	}

	return nil
}
