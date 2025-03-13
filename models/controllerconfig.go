package models

type ControllerConfig struct {
	Domain       string `json:"domain"`
	Command      string `json:"command"`
	AssemblyPath string `json:"assemblyPath"`
}

type ControllerFactoryConfig struct {
	Mappings []ControllerConfig `json:"mappings"`
}
