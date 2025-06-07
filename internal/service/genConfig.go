package service

import (
	"fmt"
)

func (s *service) generateConfig() error {
	if err := s.mkDir(fmt.Sprintf("%s/config", s.genCfg.Root)); err != nil {
		return err
	}

	if err := s.genTmpl(fmt.Sprintf("%s/config/init.go", s.genCfg.Root), "configInit", s.genCfg); err != nil {
		return err
	}

	if err := s.genTmpl(fmt.Sprintf("%s/config/scheme.go", s.genCfg.Root), "configScheme", s.genCfg); err != nil {
		return err
	}

	return nil
}
