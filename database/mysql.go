package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	azan "github.com/trihatmaja/Azan-Schedule"
)

type MySQL struct {
	db *sql.DB
}

type OptionMySQL struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
	Charset  string
}

func NewMySQL(opt OptionMySQL) (*MySQL, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", opt.User, opt.Password, opt.Host, opt.Port, opt.Database, opt.Charset))
	if err != nil {
		return &MySQL{}, err
	}

	return &MySQL{db: db}, nil
}

func (m *MySQL) Set(data []azan.CalcResult) error {
	for _, v1 := range data {
		for _, v2 := range v1.Schedule {
			tm, err := time.Parse("2006-January-2", fmt.Sprintf("%d-%s-%d", v1.Year, v1.Month, v2.Date))
			if err != nil {
				return err
			}
			_, err = m.db.Exec("INSERT INTO azan(city, dt, fajr, sunrise, zuhr, asr, maghrib, isya) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", v1.City, tm, v2.Fajr, v2.Sunrise, v2.Zuhr, v2.Asr, v2.Maghrib, v2.Isya)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *MySQL) GetAll() ([]azan.CalcResult, error) {
	return []azan.CalcResult{}, errors.New("Not Implemented Yet")
}

func (m *MySQL) GetByCity(city string) (azan.CalcResult, error) {
	return azan.CalcResult{}, errors.New("Not Implemented Yet")
}

func (m *MySQL) GetByDate(date time.Time) (azan.CalcResult, error) {
	return azan.CalcResult{}, errors.New("Not Implemented Yet")
}
