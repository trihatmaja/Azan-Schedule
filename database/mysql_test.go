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

func (suite *MySQLSuite) TearDownSuite() {
	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", suite.User, suite.Password, suite.Host, suite.Port, suite.Database, suite.Charset))
	db.Exec("truncate table azan.azan")
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
	data := []azan.CalcResult{
		{
			City:  "jakarta",
			Month: "January",
			Year:  2017,
			Schedule: []azan.AzanSchedule{
				{
					Date:    1,
					Fajr:    "04:00",
					Sunrise: "05:00",
					Zuhr:    "12:00",
					Asr:     "15:00",
					Maghrib: "18:00",
					Isya:    "19:00",
				},
				{
					Date:    2,
					Fajr:    "04:00",
					Sunrise: "05:00",
					Zuhr:    "12:00",
					Asr:     "15:00",
					Maghrib: "18:00",
					Isya:    "19:00",
				},
			},
		},
		{
			City:  "jakarta",
			Month: "February",
			Year:  2017,
			Schedule: []azan.AzanSchedule{
				{
					Date:    1,
					Fajr:    "04:00",
					Sunrise: "05:00",
					Zuhr:    "12:00",
					Asr:     "15:00",
					Maghrib: "18:00",
					Isya:    "19:00",
				},
				{
					Date:    2,
					Fajr:    "04:00",
					Sunrise: "05:00",
					Zuhr:    "12:00",
					Asr:     "15:00",
					Maghrib: "18:00",
					Isya:    "19:00",
				},
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
	data := []azan.CalcResult{
		{
			City:  "jakarta",
			Month: "January",
			Year:  2017,
			Schedule: []azan.AzanSchedule{
				{
					Date:    1,
					Fajr:    "04:00",
					Sunrise: "05:00",
					Zuhr:    "12:00",
					Asr:     "15:00",
					Maghrib: "18:00",
					Isya:    "19:00",
				},
				{
					Date:    2,
					Fajr:    "04:00",
					Sunrise: "05:00",
					Zuhr:    "12:00",
					Asr:     "15:00",
					Maghrib: "18:00",
					Isya:    "19:00",
				},
			},
		},
		{
			City:  "jakarta",
			Month: "February",
			Year:  2017,
			Schedule: []azan.AzanSchedule{
				{
					Date:    1,
					Fajr:    "04:00",
					Sunrise: "05:00",
					Zuhr:    "12:00",
					Asr:     "15:00",
					Maghrib: "18:00",
					Isya:    "19:00",
				},
				{
					Date:    2,
					Fajr:    "04:00",
					Sunrise: "05:00",
					Zuhr:    "12:00",
					Asr:     "15:00",
					Maghrib: "18:00",
					Isya:    "19:00",
				},
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
	assert.Equal(suite.T(), 4, len(k))
	assert.Equal(suite.T(), "jakarta", k[0].City)
	assert.Equal(suite.T(), "2017-02-02", k[3].Date)
}

func TestMySQLSuite(t *testing.T) {
	suite.Run(t, new(MySQLSuite))
}
