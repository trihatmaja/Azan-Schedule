package cache_test

import (
	//"bytes"
	//"compress/zlib"
	//"io"
	//"os"

	"testing"

	//"time"

	gomemcache "github.com/bradfitz/gomemcache/memcache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/trihatmaja/Azan-Schedule/cache"
)

type MemcachedSuite struct {
	suite.Suite
	c      *cache.Memcached
	opt    cache.OptionsMemcached
	client *gomemcache.Client
}

func (suite *MemcachedSuite) TearDownSuite() {
	suite.client.Delete("azan-v1:test-set")
	suite.client.Delete("azan-v1:test-get")
}

func (suite *MemcachedSuite) SetupSuite() {
	suite.opt = cache.OptionsMemcached{
		Server:    []string{"127.0.0.1:11211"},
		PrefixKey: "azan-v1",
	}
	suite.c = cache.NewMemcached(suite.opt)
	suite.client = gomemcache.New("127.0.0.1:11211")
}

func (suite *MemcachedSuite) TestNewMemcached() {
	k := cache.NewMemcached(suite.opt)

	assert.NotNil(suite.T(), k)
}

/*
	func (suite *MemcachedSuite) TestSet() {
		cc := suite.c

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

		dt, _ := json.Marshal(data)
		err := cc.Set("test-set", dt)

		assert.Nil(suite.T(), err)
	}

	func (suite *MemcachedSuite) TestGet() {
		cc := suite.c

		var tmpdata azan.CalcResult

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

		dt, _ := json.Marshal(data)
		cc.Set("test-get", dt)

		_, err := cc.Get("test-get")

		assert.Nil(suite.T(), err)

		json.Unmarshal(cl, &tmpdata)

		assert.Equal(suite.T(), "Jakarta", tmpdata.City)
	}
*/
func TestMemcachedSuite(t *testing.T) {
	suite.Run(t, new(MemcachedSuite))
}
