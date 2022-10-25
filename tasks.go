package activiti

import (
	"errors"
	"fmt"
)

// GetTask retrieves task by ID
// Endpoint: GET runtime/tasks/{taskId}
func (c *ActClient) GetTask(tid string) (*ActTask, error) {
	tk := &ActTask{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s", c.BaseURL, "/tasks/", tid), nil)
	if err != nil {
		return tk, err
	}

	if err = c.SendWithBasicAuth(req, tk); err != nil {
		return tk, err
	}

	return tk, nil
}

// GetTasks retrieves all tasks
// Endpoint: GET runtime/tasks
func (c *ActClient) GetTasks() (*ActListTasks, error) {
	tks := &ActListTasks{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.BaseURL, "/tasks"), nil)
	if err != nil {
		return tks, err
	}
	if err = c.SendWithBasicAuth(req, &tks); err != nil {
		return tks, err
	}

	return tks, nil
}

// TaskAction complete/claim/delegate/resolve a task in activiti
// Endpoint: POST runtime/tasks/{taskId}
func (c *ActClient) TaskActionComplete(tid string) error {
	if tid == "" {
		return errors.New("Task id   are required for task action ")
	}

	params := map[string]string{"payloadType": "CompleteTaskPayload"}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s%s%s", c.BaseURL, "/tasks/", tid, "/complete"), params)
	if err != nil {
		return err
	}

	if err = c.SendWithBasicAuth(req, nil); err != nil {
		return err
	}

	return nil
}

// TaskAction complete/claim/delegate/resolve a task in activiti
// Endpoint: POST runtime/tasks/{taskId}
func (c *ActClient) TaskActionClaim(tid string, assignee string) error {
	if tid == "" {
		return errors.New("Task id   are required for task action ")
	}

	params := map[string]string{"action": string(TASK_ACTION_COMPLETE), "assignee": assignee}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s%s%s", c.BaseURL, "/tasks/", tid, "/claim"), params)
	if err != nil {
		return err
	}

	if err = c.SendWithBasicAuth(req, nil); err != nil {
		return err
	}

	return nil
}

// TaskAction complete/claim/delegate/resolve a task in activiti
// Endpoint: POST runtime/tasks/{taskId}
func (c *ActClient) TaskActionAssign(tid string, assignee string) error {
	if tid == "" {
		return errors.New("Task id   are required for task action ")
	}

	params := map[string]string{"action": string(TASK_ACTION_COMPLETE), "assignee": assignee}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s%s%s", c.BaseURL, "/tasks/", tid, "/assign"), params)
	if err != nil {
		return err
	}

	if err = c.SendWithBasicAuth(req, nil); err != nil {
		return err
	}

	return nil
}
