package repo

import (
	"fmt"
	"log"
	"time"

	"github.com/SmitaJShetty/rebu/internal/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//DataRetriever interface
type DataRetriever interface {
	GetTripCount(medallionNumber []string, pickupDate *time.Time) (int, error)
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
	db, err := gorm.Open("mysql", "testuser:testpass@(127.0.0.1:3306)/carttripdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("error while opening connection to db, err: %v", err)
		return nil
	}
	return db
}

//GetTripCount returns trip count for a medallion and pickupdate
func (mr *MedallionRepo) GetTripCount(medallionList []string, pickupDate *time.Time) ([]*model.TripSummary, error) {
	if mr.db == nil {
		mr.db = getDB()
	}

	var trips []*model.TripSummary
	sql := `select count(medallion) as count, medallion, date(pickup_datetime) as pickup_date from cab_trip_data where date(pickup_datetime)=? and medallion in (?) 
				group by medallion,date(pickup_datetime)`
	db := mr.db.Debug().Raw(sql, pickupDate.Format("2006-01-02"), medallionList).Scan(&trips)
	if db.Error != nil {
		return nil, db.Error
	}

	if db.RecordNotFound() {
		return nil, fmt.Errorf("record not found for medallion number %s and pickupdate %s", medallionList, pickupDate)
	}

	return trips, nil
}
