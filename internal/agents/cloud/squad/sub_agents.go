package cloudsquad

import "context"
// SubAgentCloudListInstances - Listar instâncias
type SubAgentCloudListInstances struct{}

// Execute runs instance listing
func (s *SubAgentCloudListInstances) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"instances": []string{}}, nil
}

// SubAgentCloudListVolumes - Listar volumes
type SubAgentCloudListVolumes struct{}

// Execute runs volume listing
func (s *SubAgentCloudListVolumes) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"volumes": []string{}}, nil
}

// SubAgentCloudListNetworks - Listar redes
type SubAgentCloudListNetworks struct{}

// Execute runs network listing
func (s *SubAgentCloudListNetworks) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"networks": []string{}}, nil
}

// SubAgentCloudListTemplates - Listar templates
type SubAgentCloudListTemplates struct{}

// Execute runs template listing
func (s *SubAgentCloudListTemplates) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"templates": []string{}}, nil
}

// SubAgentCloudListResources - Listar recursos (domínios, contas)
type SubAgentCloudListResources struct{}

// Execute runs resource listing
func (s *SubAgentCloudListResources) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"resources": []string{}}, nil
}