package main

import (
	_ "embed"

	"fiber-boilerplate/pkg/app"
	"fiber-boilerplate/pkg/global"
)

//go:embed config/devconfig.yaml
var devConfigYAMLContent string

func init() {
	global.DevConfigYAMLContent = devConfigYAMLContent
}

func main() {
	app.Bootstrap()
}
