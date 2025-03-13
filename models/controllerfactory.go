package models

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"plugin"
)

type ControllerFactory struct {
	Controllers map[string]Controller
}

func NewControllerFactory(filePath string) *ControllerFactory {
	factory := &ControllerFactory{Controllers: make(map[string]Controller)}
	if err := factory.loadConfig(filePath); err != nil {
		log.Fatalf("Failed to load controller factory config: %v - %v", filePath, err)
		return nil
	} else {
		log.Printf("Successfully loaded controller factory config: %v", filePath)
		return factory
	}
}

func (f *ControllerFactory) Register(domain, command string, controller Controller) {
	key := fmt.Sprintf("%s:%s", domain, command)
	f.Controllers[key] = controller
}

func (f *ControllerFactory) GetController(domain, command string) (Controller, bool) {
	key := fmt.Sprintf("%s:%s", domain, command)
	controller, exists := f.Controllers[key]
	return controller, exists
}

func (f *ControllerFactory) loadConfig(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var config ControllerFactoryConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}

	for _, mapping := range config.Mappings {
		controller, err := createController(mapping.AssemblyPath, mapping.Domain, mapping.Command)
		if err != nil {
			return err
		}
		f.Register(mapping.Domain, mapping.Command, controller)
	}

	return nil
}

func createController(assemblyPath, domain, command string) (Controller, error) {
	p, err := plugin.Open(assemblyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open plugin: %v", err)
	}

	symController, err := p.Lookup("NewController")
	if err != nil {
		return nil, fmt.Errorf("failed to lookup NewController: %v", err)
	}

	newController, ok := symController.(func(string, string) Controller)
	if !ok {
		return nil, fmt.Errorf("invalid NewController signature")
	}

	return newController(domain, command), nil
}
