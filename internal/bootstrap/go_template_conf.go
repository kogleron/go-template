package bootstrap

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/kogleron/go-template/internal/apps"
)

func NewGoTemplateConf() apps.GoTemplateConf {
	conf := apps.GoTemplateConf{}

	err := envconfig.Process("GOTEMPLATE", &conf)
	if err != nil {
		panic(err)
	}

	return conf
}
