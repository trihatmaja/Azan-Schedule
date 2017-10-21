package database_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	azan "github.com/trihatmaja/Azan-Schedule"
	database "github.com/trihatmaja/Azan-Schedule/database"
)

type FilesSuite struct {
	suite.Suite
	FileName  string
	OutputDir string
}

func (suite *FilesSuite) TearDownSuite() {
	os.RemoveAll("./schedule.json")
}

func (suite *FilesSuite) SetupSuite() {
	suite.FileName = "schedule.json"
	suite.OutputDir = "."
}

func (suite *FilesSuite) TestFiles() {
	opt := database.OptionFiles{
		OutputDir: suite.OutputDir,
		FileName:  suite.FileName,
	}

	db := database.NewFiles(opt)

	assert.NotNil(suite.T(), db)
}

func (suite *FilesSuite) TestSet() {
	opt := database.OptionFiles{
		OutputDir: suite.OutputDir,
		FileName:  suite.FileName,
	}

	db := database.NewFiles(opt)

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

	err := db.Set(data)

	assert.Nil(suite.T(), err)
}

func TestFilesSuite(t *testing.T) {
	suite.Run(t, new(FilesSuite))
}
