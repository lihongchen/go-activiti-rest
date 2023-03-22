package activiti

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type TaskAction string

const (
	TASK_ACTION_COMPLETE TaskAction = "complete"
	TASK_ACTION_CLAIM    TaskAction = "claim"
	TASK_ACTION_DELEGATE TaskAction = "delegate"
	TASK_ACTION_RESOLVE  TaskAction = "resolve"
)

type (
	// JSONTime overrides MarshalJson method to format in ISO8601
	JSONTime time.Time

	// Client represents a Activiti 6.x REST API Client
	ActClient struct {
		Client  *http.Client
		Token   string
		BaseURL string
		Log     io.Writer // If user set log file name all requests will be logged there
	}

	expirationTime int64

	// ErrorResponse https://www.activiti.org/userguide/#_error_response_body
	ActErrorResponse struct {
		Response     *http.Response `json:"-"`
		statusCode   string         `json:"statusCode"`
		errorMessage string         `json:"errorMessage"`
	}

	ActProcessDefinition struct {
		ProcessDefinition ProcessDefinition `json:"entry,omitempty"`
	}

	ProcessDefinition struct {
		ID              string `json:"id,omitempty"`
		Key             string `json:"key,omitempty"`
		Name            string `json:"name,omitempty"`
		Category        string `json:"category,omitempty"`
		Version         int    `json:"version,omitempty"`
		AppName         string `json:"appName,omitempty"`
		AppVersion      string `json:"appVersion,omitempty"`
		ServiceFullName string `json:"serviceFullName,omitempty"`
		ServiceName     string `json:"serviceName,omitempty"`
		ServiceType     string `json:"serviceType,omitempty"`
		ServiceVersion  string `json:"serviceVersion,omitempty"`
	}
	ActProcessDefinitions struct {
		ProcessDefinitions []ActProcessDefinition `json:"entries,omitempty"`
		Pagination         Pagination             `json:"pagination,omitempty"`
	}
	ActListProcessDefinitions struct {
		List ActProcessDefinitions
	}
	Pagination struct {
		Count        int  `json:"count,omitempty"`
		HasMoreItems bool `json:"hasMoreItems,omitempty"`
		MaxItems     int  `json:"maxItems,omitempty"`
		SkipCount    int  `json:"skipCount,omitempty"`
		TotalItems   int  `json:"totalItems,omitempty"`
	}

	ActProcessInstance struct {
		ProcessInstance ProcessInstance `json:"entry,omitempty"`
	}

	ProcessInstance struct {
		ID                       string `json:"id,omitempty"`
		AppName                  string `json:"appName,omitempty"`
		Initiator                string `json:"initiator,omitempty"`
		ProcessDefinitionId      string `json:"processDefinitionId,omitempty"`
		ProcessDefinitionKey     string `json:"processDefinitionKey,omitempty"`
		ProcessDefinitionName    string `json:"processDefinitionName,omitempty"`
		ProcessDefinitionVersion int    `json:"processDefinitionVersion,omitempty"`
		ServiceFullName          string `json:"serviceFullName,omitempty"`
		ServiceName              string `json:"serviceName,omitempty"`
		ServiceType              string `json:"serviceType,omitempty"`
		ServiceVersion           string `json:"serviceVersion,omitempty"`
		StartDate                string `json:"startDate,omitempty"`
		Status                   string `json:"status,omitempty"`
	}
	ActProcessInstances struct {
		ProcessInstances []ActProcessInstance `json:"entries,omitempty"`
		Pagination       Pagination           `json:"pagination,omitempty"`
	}

	ActListProcessInstances struct {
		List ActProcessInstances
	}

	ActStartProcessInstance struct {
		ProcessDefinitionId  string                 `json:"processDefinitionId,omitempty"`
		ProcessDefinitionKey string                 `json:"processDefinitionKey,omitempty"`
		PayloadType          string                 `json:"payloadType,omitempty"`
		BusinessKey          string                 `json:"businessKey,omitempty"`
		Name                 string                 `json:"name,omitempty"`
		Variables            map[string]interface{} `json:"variables,omitempty"`
	}
	Task struct {
		Name                string `json:"name,omitempty"`
		ID                  string `json:"id,omitempty"`
		AppName             string `json:"appName,omitempty"`
		Assignee            string `json:"assignee,omitempty"`
		CreatedDate         string `json:"createdDate,omitempty"`
		FormKey             string `json:"formKey,omitempty"`
		ProcessDefinitionId string `json:"processDefinitionId,omitempty"`
		ProcessInstanceId   string `json:"processInstanceId,omitempty"`
		ServiceFullName     string `json:"serviceFullName,omitempty"`
		ServiceName         string `json:"serviceName,omitempty"`
		ServiceType         string `json:"serviceType,omitempty"`
		ServiceVersion      string `json:"serviceVersion,omitempty"`
		Status              string `json:"status,omitempty"`
		TaskDefinitionKey   string `json:"taskDefinitionKey,omitempty"`
		Priority            int    `json:"priority,omitempty"`
		Standalone          bool   `json:"standalone,omitempty"`
		BusinessKey         string `json:"businessKey,omitempty"`
		CompletedBy         string `json:"completedBy,omitempty"`
		CompletedDate       string `json:"completedDate,omitempty"`
	}
	ActTask struct {
		Task Task `json:"entry,omitempty"`
	}
	ActTasks struct {
		Tasks      []ActTask  `json:"entries,omitempty"`
		Pagination Pagination `json:"pagination,omitempty"`
	}

	ActListTasks struct {
		List ActTasks
	}

	ActUser struct {
		ID         string `json:"id,omitempty"`
		FirstName  string `json:"firstName,omitempty"`
		LastName   string `json:"lastName,omitempty"`
		Email      string `json:"email,omitempty"`
		URL        string `json:"url,omitempty"`
		PictureURL string `json:"pictureUrl,omitempty"`
		Password   string `json:"password,omitempty"`
	}

	ActUsers struct {
		Users []ActUser `json:"data,omitempty"`
		Total int       `json:"total,omitempty"`
		Start int       `json:"start,omitempty"`
		Sort  string    `json:"sort,omitempty"`
		Order string    `json:"order,omitempty"`
		Size  int       `json:"size,omitempty"`
	}
	ProcessDefinitionMeta struct {
		ID          string   `json:"id,omitempty"`
		Name        string   `json:"name,omitempty"`
		Description string   `json:"description,omitempty"`
		Version     int      `json:"version,omitempty"`
		Groups      []string `json:"groups,omitempty"`
	}
	ActProcessDefinitionMeta struct {
		Entry ProcessDefinitionMeta `json:"entry,omitempty"`
	}
)

// Error method implementation for ErrorResponse struct
func (r *ActErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %s", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.errorMessage)
}

// MarshalJSON for JSONTime
func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf(`"%s"`, time.Time(t).UTC().Format(time.RFC3339))
	return []byte(stamp), nil
}

func (e *expirationTime) UnmarshalJSON(b []byte) error {
	var n json.Number
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}
	i, err := n.Int64()
	if err != nil {
		return err
	}
	*e = expirationTime(i)
	return nil
}
