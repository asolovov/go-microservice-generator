package service

import (
	"fmt"
)

func (s *service) generateApplication() error {
	if err := s.mkDir(fmt.Sprintf("%s/internal", s.genCfg.Root)); err != nil {
		return err
	}

	if err := s.genTmpl(fmt.Sprintf("%s/internal/applicaion.go", s.genCfg.Root), "application", s.genCfg); err != nil {
		return err
	}

	return nil
}
