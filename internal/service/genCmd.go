package service

import (
	"fmt"
)

func (s *service) generateCmd() error {
	if err := s.mkDir(fmt.Sprintf("%s/cmd/root", s.genCfg.Root)); err != nil {
		return err
	}

	if err := s.mkDir(fmt.Sprintf("%s/cmd/serve", s.genCfg.Root)); err != nil {
		return err
	}

	if err := s.genTmpl(fmt.Sprintf("%s/cmd/root/root.go", s.genCfg.Root), "root", s.genCfg); err != nil {
		return err
	}

	if err := s.genTmpl(fmt.Sprintf("%s/cmd/serve/serve.go", s.genCfg.Root), "serve", s.genCfg); err != nil {
		return err
	}

	if err := s.genTmpl(fmt.Sprintf("%s/cmd/%s.go", s.genCfg.Root, s.genCfg.GoPackage), "main", s.genCfg); err != nil {
		return err
	}

	return nil
}
