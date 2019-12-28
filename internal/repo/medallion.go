package repo

import (
	"fmt"
	"log"
	"time"

	"github.com/SmitaJShetty/rebu/internal/model"
	"github.com/jinzhu/gorm"
)

//DataRetriever interface
type DataRetriever interface {
	GetTripCount(medallionNumber string, pickupDate *time.Time) (int, error)
}

//MedallionRepo repo for medallion
type MedallionRepo struct {
	db *gorm.DB
}

//NewMedallionRepo returns medallion repo
func NewMedallionRepo() *MedallionRepo {
	return &MedallionRepo{
		db: getDB(),
	}
}

func getDB() *gorm.DB {
	db, err := gorm.Open("mysql", "user:pass@(localhost)/testdb?charset=utf8&parseTime=true&loc=local")
	if err != nil {
		log.Printf("error while opening connection to db, err: %v", err)
		return nil
	}
	return db
}

//GetTripCount returns trip count for a medallion and pickupdate
func (mr *MedallionRepo) GetTripCount(medallionNumber string, pickupDate *time.Time) (int, error) {
	if mr.db == nil {
		return 0, fmt.Errorf("cannot connect as connection to db was not created")
	}

	var trips []*model.Trip
	db := mr.db.Where("medallionNum=? AND pickupDate = ?", medallionNumber, pickupDate).Find(&trips)
	if db.Error != nil {
		return 0, db.Error
	}

	if db.RecordNotFound() {
		return 0, fmt.Errorf("record not found for medallion number %s and pickupdate %s", medallionNumber, pickupDate)
	}

	return len(trips), nil
}
