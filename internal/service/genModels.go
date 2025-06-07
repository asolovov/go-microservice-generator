package service

import "fmt"

func (s *service) genModels() error {
	if err := s.mkDir(fmt.Sprintf("%s/models", s.genCfg.Root)); err != nil {
		return err
	}

	if err := s.genTmpl(fmt.Sprintf("%s/models/models.go", s.genCfg.Root), "models", s.genCfg); err != nil {
		return err
	}

	return nil
}
