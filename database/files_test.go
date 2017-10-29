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
	err := db.Set(data)

	assert.Nil(suite.T(), err)
}

func (suite *FilesSuite) TestValidate() {
	opt := database.OptionFiles{
		OutputDir: suite.OutputDir,
		FileName:  suite.FileName,
	}

	db := database.NewFiles(opt)

	var lat float64 = -6.82
	var long float64 = 106.83

	k, e := db.Validate(lat, long, "jakarta")

	assert.Nil(suite.T(), e)
	assert.Equal(suite.T(), true, k)
}

func TestFilesSuite(t *testing.T) {
	suite.Run(t, new(FilesSuite))
}
