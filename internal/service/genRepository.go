package service

import (
	"fmt"
)

func (s *service) generateRepository() error {
	if err := s.mkDir(fmt.Sprintf("%s/internal/repository", s.genCfg.Root)); err != nil {
		return err
	}

	if err := s.genTmpl(fmt.Sprintf("%s/internal/repository/cfg.go", s.genCfg.Root), "repositoryCfg", s.genCfg); err != nil {
		return err
	}

	if err := s.genTmpl(fmt.Sprintf("%s/internal/repository/repository.go", s.genCfg.Root), "repository", s.genCfg); err != nil {
		return err
	}

	return nil
}
