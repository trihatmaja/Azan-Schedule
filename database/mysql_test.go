package database_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	_ "github.com/go-sql-driver/mysql"
	azan "github.com/trihatmaja/Azan-Schedule"
	database "github.com/trihatmaja/Azan-Schedule/database"
)

type MySQLSuite struct {
	suite.Suite
	User     string
	Password string
	Host     string
	Port     string
	Database string
	Charset  string
}

func (suite *MySQLSuite) TearDownTest() {
	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", suite.User, suite.Password, suite.Host, suite.Port, suite.Database, suite.Charset))
	tx, _ := db.Begin()
	_, err := tx.Exec("truncate table city")
	if err != nil {
		tx.Rollback()
	}
	_, err = tx.Exec("truncate table schedule")
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
}

func (suite *MySQLSuite) SetupSuite() {
	suite.User = "trihatmaja"
	suite.Password = "P@ssw0rd.1"
	suite.Host = "10.10.1.13"
	suite.Port = "3306"
	suite.Database = "azan"
	suite.Charset = "utf8"
}

func (suite *MySQLSuite) TestConnections() {
	opt := database.OptionMySQL{
		User:     suite.User,
		Password: suite.Password,
		Host:     suite.Host,
		Port:     suite.Port,
		Database: suite.Database,
		Charset:  suite.Charset,
	}

	db, err := database.NewMySQL(opt)

	assert.NotNil(suite.T(), db)
	assert.Nil(suite.T(), err)
}

func (suite *MySQLSuite) TestSet() {
	data := azan.CalcResult{
		City:      "Jakarta",
		Latitude:  -6.18,
		Longitude: 106.83,
		Timezone:  7,
		Schedule: []azan.AzanSchedule{
			{
				Date:    "2017-January-1",
				Fajr:    "04:00",
				Sunrise: "05:00",
				Zuhr:    "12:00",
				Asr:     "15:00",
				Maghrib: "18:00",
				Isya:    "19:00",
			},
			{
				Date:    "2017-January-2",
				Fajr:    "04:00",
				Sunrise: "05:00",
				Zuhr:    "12:00",
				Asr:     "15:00",
				Maghrib: "18:00",
				Isya:    "19:00",
			},
		},
	}

	opt := database.OptionMySQL{
		User:     suite.User,
		Password: suite.Password,
		Host:     suite.Host,
		Port:     suite.Port,
		Database: suite.Database,
		Charset:  suite.Charset,
	}

	db, err := database.NewMySQL(opt)

	err = db.Set(data)

	assert.Nil(suite.T(), err)
}

func (suite *MySQLSuite) TestGetAll() {
	data := azan.CalcResult{
		City:      "Jakarta",
		Latitude:  -6.18,
		Longitude: 106.83,
		Timezone:  7,
		Schedule: []azan.AzanSchedule{
			{
				Date:    "2017-February-1",
				Fajr:    "04:00",
				Sunrise: "05:00",
				Zuhr:    "12:00",
				Asr:     "15:00",
				Maghrib: "18:00",
				Isya:    "19:00",
			},
			{
				Date:    "2017-February-2",
				Fajr:    "04:00",
				Sunrise: "05:00",
				Zuhr:    "12:00",
				Asr:     "15:00",
				Maghrib: "18:00",
				Isya:    "19:00",
			},
		},
	}

	opt := database.OptionMySQL{
		User:     suite.User,
		Password: suite.Password,
		Host:     suite.Host,
		Port:     suite.Port,
		Database: suite.Database,
		Charset:  suite.Charset,
	}

	db, _ := database.NewMySQL(opt)
	db.Set(data)

	k, err := db.GetAll()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(k))
	assert.Equal(suite.T(), "jakarta", k[0].City)
	assert.Equal(suite.T(), 2, len(k[0].Schedule))
	assert.Equal(suite.T(), "2017-02-02", k[0].Schedule[1].Date)
}

func (suite *MySQLSuite) TestGetByCity() {
	data := azan.CalcResult{
		City:      "Jakarta",
		Latitude:  -6.18,
		Longitude: 106.83,
		Timezone:  7,
		Schedule: []azan.AzanSchedule{
			{
				Date:    "2017-March-1",
				Fajr:    "04:00",
				Sunrise: "05:00",
				Zuhr:    "12:00",
				Asr:     "15:00",
				Maghrib: "18:00",
				Isya:    "19:00",
			},
			{
				Date:    "2017-March-2",
				Fajr:    "04:00",
				Sunrise: "05:00",
				Zuhr:    "12:00",
				Asr:     "15:00",
				Maghrib: "18:00",
				Isya:    "19:00",
			},
		},
	}

	opt := database.OptionMySQL{
		User:     suite.User,
		Password: suite.Password,
		Host:     suite.Host,
		Port:     suite.Port,
		Database: suite.Database,
		Charset:  suite.Charset,
	}

	db, _ := database.NewMySQL(opt)
	db.Set(data)

	k, err := db.GetByCity("Jakarta")

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), k)
	assert.Equal(suite.T(), 2, len(k.Schedule))
	assert.Equal(suite.T(), "2017-03-01", k.Schedule[0].Date)

}

func TestMySQLSuite(t *testing.T) {
	suite.Run(t, new(MySQLSuite))
}
