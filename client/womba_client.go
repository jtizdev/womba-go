package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// WombaClient is the HTTP client for Womba API
type WombaClient struct {
	BaseURL string
	APIKey  string
	Client  *http.Client
}

// GenerateRequest represents the request to generate tests
type GenerateRequest struct {
	StoryKey       string `json:"issue_key"`
	UploadToZephyr bool   `json:"upload_to_zephyr"`
}

// TestCase represents a single test case
type TestCase struct {
	Title          string                   `json:"title"`
	Description    string                   `json:"description"`
	Steps          []map[string]interface{} `json:"steps"`
	Preconditions  string                   `json:"preconditions,omitempty"`
	ExpectedResult string                   `json:"expected_result"`
	Priority       string                   `json:"priority"`
	TestType       string                   `json:"test_type"`
}

// TestPlan represents the test plan structure
type TestPlan struct {
	Story            map[string]interface{}   `json:"story"`
	TestCases        []TestCase               `json:"test_cases"`
	Metadata         map[string]interface{}   `json:"metadata"`
	Summary          string                   `json:"summary"`
	CoverageAnalysis string                   `json:"coverage_analysis"`
}

// GenerateResponse represents the response from test generation
type GenerateResponse struct {
	TestPlan      TestPlan               `json:"test_plan"`
	ZephyrResults map[string]interface{} `json:"zephyr_results,omitempty"`
}

// ErrorResponse represents an error from the API
type ErrorResponse struct {
	Error      string `json:"error"`
	Detail     string `json:"detail"`
	StatusCode int    `json:"status_code"`
}

// NewWombaClient creates a new Womba API client
func NewWombaClient(baseURL, apiKey string) *WombaClient {
	return &WombaClient{
		BaseURL: baseURL,
		APIKey:  apiKey,
		Client: &http.Client{
			Timeout: 120 * time.Second, // AI generation can take time
		},
	}
}

// GenerateTests generates test cases for a Jira story
func (c *WombaClient) GenerateTests(storyKey string, upload bool) (*GenerateResponse, error) {
	reqBody := GenerateRequest{
		StoryKey:       storyKey,
		UploadToZephyr: upload,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/api/v1/test-plans/generate", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != 200 {
		var errResp ErrorResponse
		if err := json.Unmarshal(body, &errResp); err != nil {
			return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
		}
		return nil, fmt.Errorf("API error %d: %s - %s", resp.StatusCode, errResp.Error, errResp.Detail)
	}

	var result GenerateResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// HealthCheck checks if the API is healthy
func (c *WombaClient) HealthCheck() (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/health", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return result, nil
}
