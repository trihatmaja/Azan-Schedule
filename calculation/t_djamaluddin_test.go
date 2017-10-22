package calculation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	calculation "github.com/trihatmaja/Azan-Schedule/calculation"
)

type TDjamaluddinSuite struct {
	suite.Suite
	Latitude     float64
	Longitude    float64
	Timezone     float64
	City         string
	TDjamaluddin *calculation.TDjamaluddin
}

func (suite *TDjamaluddinSuite) SetupSuite() {
	suite.Latitude = -6.18
	suite.Longitude = 106.83
	suite.Timezone = 7
	suite.City = "Jakarta"
	suite.TDjamaluddin = calculation.NewTDjamaluddin()
}

func (suite *TDjamaluddinSuite) TestNewTDjamaluddin() {
	c := calculation.NewTDjamaluddin()

	assert.NotNil(suite.T(), c)
}

func (suite *TDjamaluddinSuite) TestAzanJakarta() {
	schedule := suite.TDjamaluddin
	azan := schedule.Calculate(suite.Latitude, suite.Longitude, suite.Timezone, suite.City)

	assert.NotNil(suite.T(), azan)

	assert.Equal(suite.T(), suite.City, azan.City)
	assert.Equal(suite.T(), 366, len(azan.Schedule))
}

func TestTDjamaluddinSuite(t *testing.T) {
	suite.Run(t, new(TDjamaluddinSuite))
}
