package database

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path"
	"time"

	azan "github.com/trihatmaja/Azan-Schedule"
)

type Files struct {
	FileName  string
	OutputDir string
}

type OptionFiles struct {
	OutputDir string
	FileName  string
}

func NewFiles(opt OptionFiles) *Files {
	return &Files{
		FileName:  opt.FileName,
		OutputDir: opt.OutputDir,
	}
}

func (f *Files) Set(data []azan.CalcResult) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(f.OutputDir, f.FileName), js, 0644)
}

func (f *Files) GetAll() ([]azan.DbData, error) {
	return []azan.DbData{}, errors.New("Not Implemented Yet")
}

func (f *Files) GetByCity(city string) (azan.DbData, error) {
	return azan.DbData{}, errors.New("Not Implemented Yet")
}

func (f *Files) GetByDate(date time.Time) (azan.DbData, error) {
	return azan.DbData{}, errors.New("Not Implemented Yet")
}
