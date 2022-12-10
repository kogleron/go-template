//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/kogleron/go-template/internal/apps"
	"github.com/kogleron/go-template/internal/bootstrap"
)

func initGoTemplate() (*apps.GoTemplate, error) { //nolint
	wire.Build(
		bootstrap.NewGoTemplateConf,
		apps.NewGoTemplate,
	)

	return &apps.GoTemplate{}, nil
}
