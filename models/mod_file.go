package models

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"golang.org/x/mod/modfile"
)

func GetModFile() (*modfile.File, error) {
	data, err := ioutil.ReadFile("./go.mod")
	if err != nil {
		return nil, errors.Wrap(err, "Fail to load ./go.mod")
	}
	file, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Fail to parse ./go.mod")
	}
	return file, nil
}
