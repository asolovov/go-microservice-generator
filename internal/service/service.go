package service

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/misnaged/annales/logger"
	"gopkg.in/yaml.v3"

	"gmg/external/modelsAgent"
	"gmg/models"
	"gmg/templates"
)

type IService interface {
	GenerateApplication(cfgPath string) error
	GenerateLayout() error
	Close()
}

type service struct {
	genCfg *models.GenerateCfg

	agent modelsAgent.IModelsAgent
}

func New(agent modelsAgent.IModelsAgent) (IService, error) {
	s := &service{
		agent: agent,
	}

	if err := s.init(); err != nil {
		return nil, fmt.Errorf("init service: %w", err)
	}

	return s, nil
}

func (s *service) init() (err error) {
	logger.Log().Debug("service: init")
	return nil
}

func (s *service) GenerateApplication(cfgPath string) error {
	if err := s.parseCfg(cfgPath); err != nil {
		return fmt.Errorf("parse cfg: %w", err)
	}

	fmt.Printf("=== Generating application %s in %s ===\n", s.genCfg.GoPackage, s.genCfg.Root)

	if err := s.GenerateLayout(); err != nil {
		return fmt.Errorf("generate layout: %w", err)
	}

	fmt.Printf("=== APPLICAION GENERATED ===\n")
	fmt.Printf("*** Run 'go mod tidy' ***\n")

	return nil
}

func (s *service) GenerateLayout() error {
	fmt.Println("== Generating layout")

	if s.genCfg.GoModGenerate {
		fmt.Println("= Mod init")
		if err := s.modInit(); err != nil {
			return fmt.Errorf("module init: %w", err)
		}
	}

	if s.genCfg.Models.AI.Generate {
		if err := s.genModels(); err != nil {
			return fmt.Errorf("generate models: %w", err)
		}
	}

	if s.genCfg.Repository.Generate {
		fmt.Println("= Generating Repository")
		if err := s.generateRepository(); err != nil {
			return fmt.Errorf("generate repository: %w", err)
		}
	}

	if s.genCfg.Service.Generate {
		fmt.Println("= Generating Service")
		if err := s.generateService(); err != nil {
			return fmt.Errorf("generate service: %w", err)
		}
	}

	if s.genCfg.IsServer {
		fmt.Println("= Generating Server")
		if err := s.generateServer(); err != nil {
			return fmt.Errorf("generate server: %w", err)
		}

		if s.genCfg.GrpcServer.Generate {
			fmt.Println("= Generating GRPC Server")
			if err := s.generateGrpcServer(); err != nil {
				return fmt.Errorf("generate grpc server: %w", err)
			}
		}
	}

	if s.genCfg.Cfg.Generate {
		fmt.Println("= Generating Config")
		if err := s.generateConfig(); err != nil {
			return fmt.Errorf("generate config: %w", err)
		}
	}

	if s.genCfg.App.Generate {
		fmt.Println("= Generating Application")
		if err := s.generateApplication(); err != nil {
			return fmt.Errorf("generate application: %w", err)
		}
	}

	if s.genCfg.Cmd.Generate {
		fmt.Println("= Generating Cmd")
		if err := s.generateCmd(); err != nil {
			return fmt.Errorf("generate cmd: %w", err)
		}
	}

	fmt.Println("== Layout Done ==")

	return nil
}

func (s *service) Close() {
	logger.Log().Debug("service: close")
}

func (s *service) parseCfg(cfgPath string) error {
	s.genCfg = &models.GenerateCfg{}

	f, err := os.Open(cfgPath)
	if err != nil {
		return fmt.Errorf("open cfg: %w", err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("read cfg: %w", err)
	}

	if err = yaml.Unmarshal(data, s.genCfg); err != nil {
		return fmt.Errorf("parse cfg: %w", err)
	}

	s.genCfg.IsServer = s.genCfg.GrpcServer.Generate

	if s.genCfg.GrpcServer.Generate {
		s.genCfg.GrpcServer.ProtoModels.SameAsService = s.genCfg.GrpcServer.ProtoModels.Path == s.genCfg.GrpcServer.ProtoService.Path
	}

	return nil
}

func (s *service) modInit() error {
	fmt.Println("  * initializing go module")

	cmd := exec.Command("go", "mod", "init", s.genCfg.GoPackage)
	cmd.Dir = s.genCfg.Root
	return cmd.Run()
}

func (s *service) mkDir(path string) error {
	fmt.Printf("  * creating %s\n", path)

	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("err create %s dir: %s", path, err)
	}

	return nil
}

func (s *service) genTmpl(path string, name string, data interface{}) error {
	fmt.Printf("  * generating %s\n", path)

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create %s file: %w", name, err)
	}
	defer f.Close()

	if err = templates.Tmpls.ExecuteTemplate(f, name, data); err != nil {
		return fmt.Errorf("template execute %s: %w", name, err)
	}

	return nil
}
