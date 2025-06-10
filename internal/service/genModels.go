package service

import (
	"fmt"
	"os"
)

func (s *service) genModels() error {
	if err := s.agent.SetTools(s.genCfg.Root, s.genCfg.GoPackage, s.genCfg.Models); err != nil {
		return fmt.Errorf("set agent tools: %w", err)
	}

	resp, err := s.agent.GenerateModels()
	if err != nil {
		return fmt.Errorf("agent generate models: %w", err)
	}

	for _, f := range resp.Items {
		if err = os.WriteFile(fmt.Sprintf("%s%s", s.genCfg.Root, f.FileName), []byte(f.Code), 0777); err != nil {
			return fmt.Errorf("write file %s: %w", f.FileName, err)
		}
	}

	return nil
}
