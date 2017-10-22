package azan

import (
	"errors"
)

// azan
type Azan struct {
	db    DbProvider
	cache CacheProvider
	calc  CalcProvider
}

func New(database DbProvider, cache CacheProvider, calculation CalcProvider) *Azan {
	return &Azan{
		db:    database,
		cache: cache,
		calc:  calculation,
	}
}

func (a *Azan) Generate(latitude, longitude, timezone float64, city string) error {
	test, _ := a.db.GetByCity(city)
	if test.City != "" {
		return errors.New("city exist")
	}
	res := a.calc.Calculate(latitude, longitude, timezone, city)

	err := a.db.Set(res)
	if err != nil {
		return err
	}
	return nil
}

func (a *Azan) GetAll() ([]CalcResult, error) {
	return a.db.GetAll()
}

func (a *Azan) GetByCity(city string) (CalcResult, error) {
	return a.db.GetByCity(city)
}

func (a *Azan) GetCache(key string) ([]byte, error) {
	return a.cache.Get(key)
}

func (a *Azan) SetCache(key string, data []byte) error {
	return a.cache.Set(key, data)
}
