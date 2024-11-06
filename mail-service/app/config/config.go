package config

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/golobby/dotenv"
)

func Init() (Conf, error) {
	conf := Conf{}
	//Load Environment file
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b), "../")
	os.Chdir(filepath.Dir(d))
	file, err := os.Open(".env")

	if err != nil {
		return conf, errors.New("error loading .env file")
	}
	err = dotenv.NewDecoder(file).Decode(&conf)
	if err != nil {
		return conf, errors.New("cannot decode .env file")
	}
	return conf, err
}
