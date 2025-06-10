package modelsAgent

import (
	"fmt"

	"gmg/models"
)

func (a *modelsAgent) SetTools(root, goPackage string, cfg models.ModelsCfg) error {
	a.root = root
	a.genCfg = cfg
	a.goPackage = goPackage

	a.agent.SetSystemInstructions(system)

	fmt.Println("~~ Agent tools Set")

	return nil
}
