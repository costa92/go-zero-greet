package config

import (
	"github.com/tal-tech/go-zero/rest"
	"greet/libs"
)

type Config struct {
	rest.RestConf
	Sms libs.Sms
}
