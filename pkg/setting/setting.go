package setting

import (
	"fiber-boilerplate/pkg/global"
	"strings"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	reader := strings.NewReader(global.DevConfigYAMLContent)

	vp := viper.New()
	vp.SetConfigType("yaml")

	err := vp.ReadConfig(reader)

	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)

	if err != nil {
		return err
	}

	return nil
}
