package utils

import (
	"fmt"
	"time"

	"github.com/magiconair/properties"
)

var PROPERTIES_PATH = "${HOME}/config.properties"
var SECRET_PATH = "${HOME}/secrets.properties"

type ConfigProperties struct {
	Host     string        `properties:"host"`
	Port     string        `properties:"port,default=8080"`
	Password string        `properties:"password"`
	Timeout  time.Duration `properties:"timeout,default=5s"`
}

func NewProperties() *properties.Properties {
	p := properties.MustLoadFiles([]string{
		PROPERTIES_PATH,
		SECRET_PATH,
	}, properties.UTF8, true)
	return p
}

func NewConfigProperies() ConfigProperties {
	p := NewProperties()

	var cfg ConfigProperties
	if err := p.Decode(&cfg); err != nil {
		fmt.Print(err)
	}
	return cfg
}

func ConfigPropertiesToString(cfg ConfigProperties) string {
	return fmt.Sprintf("%+v", cfg)
}
