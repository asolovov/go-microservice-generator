package models

type GenerateCfg struct {
	Root          string        `yaml:"root" json:"Root"`
	GoPackage     string        `yaml:"go-package" json:"GoPackage"`
	GoModGenerate bool          `yaml:"go-mod-generate" json:"GoModGenerate"`
	Cfg           CfgCfg        `yaml:"cfg" json:"Cfg"`
	App           AppCfg        `yaml:"app" json:"App"`
	Cmd           CmdCfg        `yaml:"cmd" json:"Cmd"`
	Service       ServiceCfg    `yaml:"service" json:"Service"`
	IsServer      bool          `yaml:"-" json:"IsServer"`
	GrpcServer    GrpcServerCfg `yaml:"grpc-server" json:"GrpcServer"`
	Repository    RepositoryCfg `yaml:"repository" json:"Repository"`
	Models        ModelsCfg     `yaml:"models" json:"Models"`
}

type GrpcServerCfg struct {
	Generate     bool         `yaml:"generate" json:"-"`
	DefaultAddr  string       `yaml:"default-addr" json:"DefaultAddr"`
	ProtoService ProtoService `yaml:"proto-service" json:"ProtoService"`
	ProtoModels  ProtoModels  `yaml:"proto-models" json:"ProtoModels"`
}

type ServiceCfg struct {
	Generate bool `yaml:"generate" json:"-"`
}

type CfgCfg struct {
	Generate bool `yaml:"generate" json:"-"`
}

type AppCfg struct {
	Generate bool `yaml:"generate" json:"-"`
}

type CmdCfg struct {
	Generate bool `yaml:"generate" json:"-"`
}

type ProtoService struct {
	Path   string `yaml:"path" json:"Path"`
	PbName string `yaml:"pb-name" json:"PbName"`
	GoName string `yaml:"go-name" json:"GoName"`
}

type ProtoModels struct {
	Path          string `yaml:"path" json:"Path"`
	SameAsService bool   `yaml:"-" json:"SameAsService"`
	GoName        string `yaml:"go-name" json:"GoName"`
}

type RepositoryCfg struct {
	Generate    bool   `yaml:"generate" json:"-"`
	DefaultAddr string `yaml:"default-addr" json:"DefaultAddr"`
}

type ModelsCfg struct {
	AI ModelsAICfg `yaml:"ai" json:"-"`
}

type ModelsAICfg struct {
	Generate         bool               `yaml:"generate" json:"-"`
	DomainModelPaths []string           `yaml:"domain-model-paths" json:"-"`
	PbModelPaths     []string           `yaml:"pb-model-paths" json:"-"`
	MigrationPaths   []string           `yaml:"migration-paths" json:"-"`
	ParsingRules     ModelsParsingRules `yaml:"parsing-rules" json:"-"`
	AdditionalPrompt string             `yaml:"additional-prompt" json:"-"`
}

type ModelsParsingRules struct {
	ProtoDomain map[string]string `yaml:"proto-domain" json:"-"`
	DomainDB    map[string]string `yaml:"domain-db" json:"-"`
}
