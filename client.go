package activiti

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

// NewClient returns new Client struct
// BaseURL is the activiti rest api url, for example 'http://localhost:8080'
func NewClient(token string, baseURL string) (*ActClient, error) {
	if token == "" || baseURL == "" {
		return nil, errors.New("token and BaseURL are required to create a Client ")
	}

	return &ActClient{
		Client:  &http.Client{},
		Token:   token,
		BaseURL: baseURL,
	}, nil
}

// SetHTTPClient sets *http.Client to current client
func (c *ActClient) SetHTTPClient(client *http.Client) {
	c.Client = client
}

// SetLog will set/change the output destination.
// If log file is set all requests and responses will be logged to this Writer
func (c *ActClient) SetLog(log io.Writer) {
	c.Log = log
}

// Send makes a request to the API, the response body will be
// unmarshaled into v, or if v is an io.Writer, the response will
// be written to it without decoding
func (c *ActClient) Send(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
		data []byte
	)

	// Set default headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "zh-CN,en_US")
	if req.Header.Get("Content-type") == "" {
		req.Header.Set("Content-type", "application/json")
	}

	resp, err = c.Client.Do(req)
	c.log(req, resp)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {

		errResp := &ActErrorResponse{Response: resp}
		data, err = ioutil.ReadAll(resp.Body)
		if err == nil && len(data) > 0 {
			json.Unmarshal(data, errResp)
		}
		fmt.Println(string(data))
		return errResp
	}

	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		io.Copy(w, resp.Body)
		return nil
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	if len(bodyBytes) == 0 {
		return nil
	}
	return json.Unmarshal(bodyBytes, v)
}
func (c *ActClient) GetImg(req *http.Request, v interface{}) ([]byte, error) {
	var (
		err  error
		resp *http.Response
		data []byte
	)

	// Set default headers
	req.Header.Set("Accept-Language", "zh-CN,en_US")

	resp, err = c.Client.Do(req)
	c.log(req, resp)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {

		errResp := &ActErrorResponse{Response: resp}
		data, err = ioutil.ReadAll(resp.Body)
		if err == nil && len(data) > 0 {
			json.Unmarshal(data, errResp)
		}
		fmt.Println(string(data))
		return nil, errResp
	}

	if v == nil {
		return nil, nil
	}

	return ioutil.ReadAll(resp.Body)
}

// SendWithBasicAuth makes a request to the API using username:password basic auth
func (c *ActClient) SendWithBasicAuth(req *http.Request, v interface{}) error {
	req.Header.Add("Authorization", "Bearer "+c.Token)
	return c.Send(req, v)
}

// SendWithBasicAuth makes a request to the API using username:password basic auth
func (c *ActClient) GetImgWithBasicAuth(req *http.Request, v interface{}) ([]byte, error) {
	req.Header.Add("Authorization", "Bearer "+c.Token)
	return c.GetImg(req, v)
}

// NewRequest constructs a request
// Convert payload to a JSON
func (c *ActClient) NewRequest(method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		var b []byte
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}
	return http.NewRequest(method, url, buf)
}

// log will dump request and response to the log file
func (c *ActClient) log(r *http.Request, resp *http.Response) {
	if c.Log != nil {
		var (
			reqDump  string
			respDump []byte
		)

		if r != nil {
			reqDump = fmt.Sprintf("%s %s. Data: %s", r.Method, r.URL.String(), r.Form.Encode())
		}
		if resp != nil {
			respDump, _ = httputil.DumpResponse(resp, true)
		}

		c.Log.Write([]byte(fmt.Sprintf("Request: %s\nResponse: %s\n", reqDump, string(respDump))))
	}
}
