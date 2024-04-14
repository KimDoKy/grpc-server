package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct{}

// 환경변수 파일을 불러올 함수
func NewConfig(path string) *Config {
	c := new(Config)
	if file, err := os.Open(path); err != nil {
		panic(err)
	} else {
		defer file.Close()
		if err = toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			return c
		}
	}
}