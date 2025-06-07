package service

import (
	"fmt"
)

func (s *service) generateServer() error {
	if err := s.mkDir(fmt.Sprintf("%s/internal/server", s.genCfg.Root)); err != nil {
		return err
	}

	if err := s.genTmpl(fmt.Sprintf("%s/internal/server/server.go", s.genCfg.Root), "server", s.genCfg); err != nil {
		return err
	}

	return nil
}

func (s *service) generateGrpcServer() error {
	if err := s.mkDir(fmt.Sprintf("%s/internal/server/grpcServer", s.genCfg.Root)); err != nil {
		return err
	}

	if err := s.genTmpl(fmt.Sprintf("%s/internal/server/grpcServer/cfg.go", s.genCfg.Root), "grpcServerCfg", s.genCfg); err != nil {
		return err
	}

	if err := s.genTmpl(fmt.Sprintf("%s/internal/server/grpcServer/handlers.go", s.genCfg.Root), "grpcHandlers", s.genCfg); err != nil {
		return err
	}

	if err := s.genTmpl(fmt.Sprintf("%s/internal/server/grpcServer/server.go", s.genCfg.Root), "grpc", s.genCfg); err != nil {
		return err
	}

	return nil
}
