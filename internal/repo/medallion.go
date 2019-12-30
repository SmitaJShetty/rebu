package repo

import (
	"fmt"
	"log"
	"os"
	"time"

	"rebu/internal/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//DataRetriever interface
type DataRetriever interface {
	GetTripCount(medallionList []string, pickupDate *time.Time) ([]*model.TripSummary, error)
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
	dbUser:= os.Getenv("testuser")
	dbPass:= os.Getenv("testpass")
	connStr:= fmt.Sprintf("%s:%s@(127.0.0.1:3306)/carttripdb?charset=utf8&parseTime=True&loc=Local", dbUser,dbPass)
	db, err := gorm.Open("mysql",connStr )
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
	sql := `SELECT COUNT(medallion) AS count, 
				medallion, 
				DATE_FORMAT(pickup_datetime,"%Y-%m-%d") AS pickup_date 
			FROM cab_trip_data WHERE date(pickup_datetime)=? AND medallion IN (?) 
			GROUP BY medallion,
				DATE_FORMAT(pickup_datetime,"%Y-%m-%d") `

	db := mr.db.Debug().Raw(sql, pickupDate.Format("2006-01-02"), medallionList).Scan(&trips)
	if db.Error != nil {
		return nil, db.Error
	}

	if db.RecordNotFound() {
		return nil, fmt.Errorf("record not found for medallion number %s and pickupdate %s", medallionList, pickupDate)
	}

	return trips, nil
}
