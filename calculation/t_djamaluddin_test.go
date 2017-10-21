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
	suite.TDjamaluddin = calculation.NewTDjamaluddin(suite.Latitude, suite.Longitude, suite.Timezone, suite.City)
}

func (suite *TDjamaluddinSuite) TestNewTDjamaluddin() {
	c := calculation.NewTDjamaluddin(suite.Latitude, suite.Longitude, suite.Timezone, suite.City)

	assert.NotNil(suite.T(), c)
}

func (suite *TDjamaluddinSuite) TestAzanJakarta() {
	schedule := suite.TDjamaluddin
	azan := schedule.Calculate()

	assert.NotNil(suite.T(), azan)

	assert.Equal(suite.T(), len(azan), 12)
	assert.Equal(suite.T(), azan[11].Month, "December")
}

func TestTDjamaluddinSuite(t *testing.T) {
	suite.Run(t, new(TDjamaluddinSuite))
}
