package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/easy-cloud-Knet/KWS_Control/request/model"
	"github.com/easy-cloud-Knet/KWS_Control/structure"
)

type CoreClient struct {
	baseURL string
	client  *http.Client
}

func NewCoreClient(core *structure.Core) *CoreClient {
	return &CoreClient{
		baseURL: fmt.Sprintf("http://%s:%d", core.IP, core.Port),
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *CoreClient) doRequest(method, path string, requestBody interface{}, responseBody interface{}) error {
	var reqBodyReader io.Reader
	if requestBody != nil {
		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBodyReader = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, c.baseURL+path, reqBodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	if requestBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	if responseBody != nil {
		if err := json.NewDecoder(resp.Body).Decode(responseBody); err != nil {
			return fmt.Errorf("failed to decode response body: %w", err)
		}
	}

	return nil
}

func (c *CoreClient) CreateVM(req model.CreateVMRequest) (model.CreateVMResponse, error) {
	var response model.CreateVMResponse
	err := c.doRequest(http.MethodPost, "/createVm", req, &response)
	if err != nil {
		return model.CreateVMResponse{}, err
	}
	return response, nil
}
