package database

import (
	"database/sql"
	"fmt"
	"strings"
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

func (m *MySQL) Set(data azan.CalcResult) error {

	tx, err := m.db.Begin()

	res, err := tx.Exec("INSERT INTO city(city, latitude, longitude, timezone) VALUES(?, ?, ?, ?)", strings.ToLower(data.City), data.Latitude, data.Longitude, data.Timezone)
	if err != nil {
		tx.Rollback()
		return err
	}

	city_id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, v2 := range data.Schedule {
		tm, err := time.Parse("2006-January-2", v2.Date)
		if err != nil {
			continue
		}

		_, err = tx.Exec("INSERT INTO schedule(city_id, dt, fajr, sunrise, zuhr, asr, maghrib, isya) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", city_id, tm, v2.Fajr, v2.Sunrise, v2.Zuhr, v2.Asr, v2.Maghrib, v2.Isya)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func (m *MySQL) Validate(lat, long float64, city string) (bool, error) {
	var val bool

	err := m.db.QueryRow("SELECT EXISTS (SELECT id FROM city where city = ? and cast(latitude as decimal) = cast(? as decimal) and cast(longitude as decimal) = cast(? as decimal))", city, lat, long).Scan(&val)
	if err != nil {
		return false, err
	}

	if val {
		return true, nil
	}

	return false, nil
}

func (m *MySQL) GetAll() ([]azan.CalcResult, error) {
	var retval []azan.CalcResult

	rows1, err := m.db.Query(`select * from city`)
	if err != nil {
		return []azan.CalcResult{}, err
	}

	defer rows1.Close()

	for rows1.Next() {
		var cityId int
		var dt azan.CalcResult

		err = rows1.Scan(&cityId, &dt.City, &dt.Latitude, &dt.Longitude, &dt.Timezone)
		if err != nil {
			return []azan.CalcResult{}, err
		}

		rows2, err := m.db.Query(`select * from schedule where city_id = ?`, cityId)
		if err != nil {
			return []azan.CalcResult{}, err
		}

		defer rows2.Close()

		for rows2.Next() {
			var as azan.AzanSchedule
			var scheduleId int
			var ccity_id int
			err = rows2.Scan(&scheduleId, &ccity_id, &as.Date, &as.Fajr, &as.Sunrise, &as.Zuhr, &as.Asr, &as.Maghrib, &as.Isya)
			if err != nil {
				return []azan.CalcResult{}, err
			}

			dt.Schedule = append(dt.Schedule, as)
		}

		retval = append(retval, dt)
	}

	return retval, nil

}

func (m *MySQL) GetByCity(city string) (azan.CalcResult, error) {
	var retval azan.CalcResult

	rows1 := m.db.QueryRow(`select * from city where city = ?`, strings.ToLower(city))

	var cityId int

	err := rows1.Scan(&cityId, &retval.City, &retval.Latitude, &retval.Longitude, &retval.Timezone)
	if err != nil {
		return azan.CalcResult{}, err
	}

	rows2, err := m.db.Query(`select * from schedule where city_id = ?`, cityId)
	if err != nil {
		return azan.CalcResult{}, err
	}

	defer rows2.Close()

	for rows2.Next() {
		var as azan.AzanSchedule
		var scheduleId int
		var ccity_id int
		err = rows2.Scan(&scheduleId, &ccity_id, &as.Date, &as.Fajr, &as.Sunrise, &as.Zuhr, &as.Asr, &as.Maghrib, &as.Isya)
		if err != nil {
			return azan.CalcResult{}, err
		}

		retval.Schedule = append(retval.Schedule, as)
	}

	return retval, nil
}

func (m *MySQL) GetByCityDate(city string, date time.Time) (azan.CalcResult, error) {
	var retval azan.CalcResult

	rows1 := m.db.QueryRow(`select * from city where city = ?`, strings.ToLower(city))

	var cityId int

	err := rows1.Scan(&cityId, &retval.City, &retval.Latitude, &retval.Longitude, &retval.Timezone)
	if err != nil {
		return azan.CalcResult{}, err
	}

	rows2, err := m.db.Query(`select * from schedule where city_id = ? and dt = ?`, cityId, date)
	if err != nil {
		return azan.CalcResult{}, err
	}

	defer rows2.Close()

	for rows2.Next() {
		var as azan.AzanSchedule
		var scheduleId int
		var ccity_id int
		err = rows2.Scan(&scheduleId, &ccity_id, &as.Date, &as.Fajr, &as.Sunrise, &as.Zuhr, &as.Asr, &as.Maghrib, &as.Isya)
		if err != nil {
			return azan.CalcResult{}, err
		}

		retval.Schedule = append(retval.Schedule, as)
	}

	return retval, nil
}

func (m *MySQL) GetByCityMonth(city string, month int) (azan.CalcResult, error) {
	var retval azan.CalcResult

	rows1 := m.db.QueryRow(`select * from city where city = ?`, strings.ToLower(city))

	var cityId int

	err := rows1.Scan(&cityId, &retval.City, &retval.Latitude, &retval.Longitude, &retval.Timezone)
	if err != nil {
		return azan.CalcResult{}, err
	}

	rows2, err := m.db.Query(`select * from schedule where city_id = ? and month(dt) = ?`, cityId, month)
	if err != nil {
		return azan.CalcResult{}, err
	}

	defer rows2.Close()

	for rows2.Next() {
		var as azan.AzanSchedule
		var scheduleId int
		var ccity_id int
		err = rows2.Scan(&scheduleId, &ccity_id, &as.Date, &as.Fajr, &as.Sunrise, &as.Zuhr, &as.Asr, &as.Maghrib, &as.Isya)
		if err != nil {
			return azan.CalcResult{}, err
		}

		retval.Schedule = append(retval.Schedule, as)
	}

	return retval, nil
}

func (m *MySQL) GetCities() ([]azan.CalcResult, error) {

	var retval []azan.CalcResult

	rows1, err := m.db.Query(`select city, latitude, longitude, timezone from city`)
	if err != nil {
		return []azan.CalcResult{}, err
	}

	for rows1.Next() {

		var a azan.CalcResult

		err := rows1.Scan(&a.City, &a.Latitude, &a.Longitude, &a.Timezone)
		if err != nil {
			return []azan.CalcResult{}, err
		}

		retval = append(retval, a)
	}

	return retval, nil
}
