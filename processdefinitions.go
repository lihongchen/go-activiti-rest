package activiti

import (
	"fmt"
)

// GetProcessDefinition retrieves process definition by ID
// Endpoint: GET repository/process-definitions/{processDefinitionId}
func (c *ActClient) GetProcessDefinition(pid string) (*ActProcessDefinition, error) {
	pd := &ActProcessDefinition{}
	url := fmt.Sprintf("%s%s%s", c.BaseURL, "/process-definitions/", pid)

	fmt.Println(url)
	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s", c.BaseURL, "/process-definitions/", pid), nil)
	if err != nil {
		return pd, err
	}
	if err = c.SendWithBasicAuth(req, &pd); err != nil {
		return pd, err
	}

	return pd, nil
}

// GetProcessDefinitions retrieves all process definitions
// Endpoint: GET repository/process-definitions
func (c *ActClient) GetProcessDefinitions() (ActListProcessDefinitions, error) {
	pds := ActListProcessDefinitions{}
	url := fmt.Sprintf("%s%s", c.BaseURL, "/process-definitions")
	fmt.Println(url)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return pds, err
	}

	if err = c.SendWithBasicAuth(req, &pds); err != nil {
		return pds, err
	}
	return pds, nil
}

//GetProcessDefinitionMeta 获取process definition 元数据
func (c *ActClient) GetProcessDefinitionMeta(pid string) (*ActProcessDefinitionMeta, error) {
	pd := &ActProcessDefinitionMeta{}
	url := fmt.Sprintf("%s%s%s%s", c.BaseURL, "/process-definitions/", pid, "/meta")

	fmt.Println(url)
	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s%s", c.BaseURL, "/process-definitions/", pid, "/meta"), nil)
	if err != nil {
		return pd, err
	}
	if err = c.SendWithBasicAuth(req, &pd); err != nil {
		return pd, err
	}
	return pd, nil
}
