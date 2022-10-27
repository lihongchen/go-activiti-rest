package activiti

import (
	"errors"
	"fmt"
)

// GetProcessInstance retrieves process instance by ID
// Endpoint: GET runtime/process-instances/{processInstanceId}
func (c *ActClient) GetProcessInstance(pid string) (*ActProcessInstance, error) {
	pi := &ActProcessInstance{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s", c.BaseURL, "/process-instances/", pid), nil)
	if err != nil {
		return pi, err
	}

	if err = c.SendWithBasicAuth(req, pi); err != nil {
		return pi, err
	}

	return pi, nil
}

// GetProcessInstance retrieves process instance by ID
// Endpoint: GET runtime/process-instances/{processInstanceId}
func (c *ActClient) GetProcessDiagram(pid string) ([]byte, error) {
	var pDiagram []byte

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s%s", c.BaseURL, "/process-instances/", pid, "/model"), nil)
	req.Header.Set("content-type", "image/svg+xml;charset=UTF-8")

	if err != nil {
		return pDiagram, err
	}

	if pDiagram, err = c.GetImgWithBasicAuth(req, pDiagram); err != nil {
		return pDiagram, err
	}
	return pDiagram, nil
}

// GetProcessInstances retrieves all process instances
// Endpoint: GET runtime/process-instances
func (c *ActClient) GetProcessInstances() (*ActListProcessInstances, error) {
	pis := &ActListProcessInstances{}
	url := fmt.Sprintf("%s%s", c.BaseURL, "/process-instances")

	fmt.Println(url)
	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.BaseURL, "/process-instances"), nil)
	if err != nil {
		return pis, err
	}
	if err = c.SendWithBasicAuth(req, pis); err != nil {
		return pis, err
	}

	return pis, nil
}

// startProcessInstance start a process instance in activiti
// Endpoint: POST runtime/process-instances
func (c *ActClient) startProcessInstance(s ActStartProcessInstance) (*ActProcessInstance, error) {
	pi := &ActProcessInstance{}
	s.PayloadType = "StartProcessPayload"
	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.BaseURL, "/process-instances"), s)
	if err != nil {
		return pi, err
	}

	if err = c.SendWithBasicAuth(req, &pi); err != nil {
		return pi, err
	}

	return pi, nil
}

// Start a process instance by process definition id
func (c *ActClient) StartProcessInstanceById(pid string) (*ActProcessInstance, error) {
	if pid == "" {
		return nil, errors.New("Process definition id is required to start a process instance ")
	}

	return c.startProcessInstance(ActStartProcessInstance{ProcessDefinitionId: pid})
}

// Start a process instance by process definition key
func (c *ActClient) StartProcessInstanceByKey(key string) (*ActProcessInstance, error) {
	if key == "" {
		return nil, errors.New("Process definition key is required to start a process instance ")
	}

	return c.startProcessInstance(ActStartProcessInstance{ProcessDefinitionKey: key})
}

// Start a process instance by process definition key and variables
func (c *ActClient) StartProcessInstanceWithVariables(key string, variables map[string]interface{}) (*ActProcessInstance, error) {
	if key == "" {
		return nil, errors.New("key is required to start a process instance ")
	}

	return c.startProcessInstance(ActStartProcessInstance{ProcessDefinitionKey: key, Variables: variables})
}

// Start a process instance by process definition key and variables
func (c *ActClient) StartProcessInstanceWithBusinessKeyAndVariables(key, BusinessKey string, variables map[string]interface{}) (*ActProcessInstance, error) {
	if key == "" {
		return nil, errors.New("key is required to start a process instance ")
	}

	return c.startProcessInstance(ActStartProcessInstance{ProcessDefinitionKey: key, BusinessKey: BusinessKey, Variables: variables})
}
