package modelsAgent

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/misnaged/annales/logger"
	"github.com/pontus-devoteam/agent-sdk-go/pkg/agent"
	"github.com/pontus-devoteam/agent-sdk-go/pkg/model/providers/openai"
	"github.com/pontus-devoteam/agent-sdk-go/pkg/runner"

	"gmg/models"
)

type IModelsAgent interface {
	GenerateModels() (*models.AgentResponse, error)
	SetTools(root, goPackage string, cfg models.ModelsCfg) error
	Close()
}

type modelsAgent struct {
	cfg *Config

	genCfg    models.ModelsCfg
	root      string
	goPackage string

	provider *openai.Provider
	agent    *agent.Agent
	runner   *runner.Runner

	mu sync.Mutex
}

func NewModelsAgent(cfg *Config) (IModelsAgent, error) {
	a := &modelsAgent{
		cfg: cfg,
		mu:  sync.Mutex{},
	}

	if err := a.init(); err != nil {
		return nil, fmt.Errorf("models agent init: %w", err)
	}

	return a, nil
}

func (a *modelsAgent) init() (err error) {
	logger.Log().Infof("models-agent: init")

	a.provider = openai.NewProvider(a.cfg.OpenAIKey)
	a.provider.SetDefaultModel(a.cfg.DefaultModel)

	a.agent = agent.NewAgent("GMG-ModelsAgent")
	a.agent.WithModel(a.cfg.DefaultModel)

	a.runner = runner.NewRunner()
	a.runner.WithDefaultProvider(a.provider)

	return nil
}

func (a *modelsAgent) GenerateModels() (*models.AgentResponse, error) {
	fmt.Println("~~ Agent generating models")

	prompt := a.getRunFlowInput()

	logger.Log().Debugf("agent input prompt: %s", prompt)

	result, err := a.runner.RunSync(a.agent, &runner.RunOptions{
		Input: prompt,
	})
	if err != nil {
		return nil, fmt.Errorf("models agent: run sync %w", err)
	}

	fmt.Println("~~ Agent received result")

	logger.Log().Debugf("agent result:\n%s", result.FinalOutput)

	r := &models.AgentResponse{
		Items: make([]*models.AgentFile, 0),
	}
	if err = json.Unmarshal([]byte(result.FinalOutput.(string)), &r.Items); err != nil {
		return nil, fmt.Errorf("models agent: unmarshal response %w", err)
	}

	return r, nil
}

func (a *modelsAgent) Close() {
	logger.Log().Infof("models-agent: close")
}
