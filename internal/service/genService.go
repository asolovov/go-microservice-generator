package service

import (
	"fmt"
)

func (s *service) generateService() error {
	if err := s.mkDir(fmt.Sprintf("%s/internal/service", s.genCfg.Root)); err != nil {
		return err
	}

	if err := s.genTmpl(fmt.Sprintf("%s/internal/service/cfg.go", s.genCfg.Root), "serviceCfg", s.genCfg); err != nil {
		return err
	}

	if err := s.genTmpl(fmt.Sprintf("%s/internal/service/service.go", s.genCfg.Root), "service", s.genCfg); err != nil {
		return err
	}

	return nil
}
