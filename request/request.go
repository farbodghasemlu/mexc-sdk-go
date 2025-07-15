package request

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	// "net/url"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type RequestService struct {
	apiKey    string
	apiSecret string
	baseURL   string
	client    *resty.Client
}

func NewRequestService(apiKey, apiSecret, baseURL string) *RequestService {
	return &RequestService{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		baseURL:   baseURL,
		client:    resty.New(),
	}
}

// PublicGet makes a public GET request
func (s *RequestService) PublicGet(endpoint string, params string) (interface{}, error) {
	urlStr := s.baseURL + endpoint
	
	if params != "" {
		strParams := s.jsonToParamStr(params)
		urlStr += "?" + strParams
	}
	
	resp, err := s.client.R().Get(urlStr)
	if err != nil {
		return nil, fmt.Errorf("public GET request failed: %w", err)
	}
	
	return resp, nil
}

// PrivateGet makes an authenticated GET request
func (s *RequestService) PrivateGet(endpoint string, params string) (interface{}, error) {
	urlStr := s.baseURL + endpoint
	timestamp := time.Now().UnixNano() / 1e6
	
	var message string
	if params == "" {
		message = fmt.Sprintf("timestamp=%d", timestamp)
	} else {
		strParams := s.jsonToParamStr(params)
		message = fmt.Sprintf("%s&timestamp=%d", strParams, timestamp)
	}
	
	sign := s.computeHmac256(message)
	fullURL := fmt.Sprintf("%s?%s&timestamp=%d&signature=%s", urlStr, message, timestamp, sign)
	
	resp, err := s.client.R().
		SetHeaders(map[string]string{
			"X-MEXC-APIKEY": s.apiKey,
			"Content-Type":  "application/json",
		}).
		Get(fullURL)
	
	if err != nil {
		return nil, fmt.Errorf("private GET request failed: %w", err)
	}
	
	return resp, nil
}

// PrivatePost makes an authenticated POST request
func (s *RequestService) PrivatePost(endpoint string, params string) (interface{}, error) {
	urlStr := s.baseURL + endpoint
	timestamp := time.Now().UnixNano() / 1e6
	
	var message string
	if params == "" {
		message = fmt.Sprintf("timestamp=%d", timestamp)
	} else {
		strParams := s.jsonToParamStr(params)
		message = fmt.Sprintf("%s&timestamp=%d", strParams, timestamp)
	}
	
	sign := s.computeHmac256(message)
	fullURL := fmt.Sprintf("%s?timestamp=%d&signature=%s", urlStr, timestamp, sign)
	
	var body interface{}
	if params != "" {
		if err := json.Unmarshal([]byte(params), &body); err != nil {
			return nil, fmt.Errorf("failed to unmarshal params: %w", err)
		}
	}
	
	resp, err := s.client.R().
		SetHeaders(map[string]string{
			"X-MEXC-APIKEY": s.apiKey,
			"Content-Type":  "application/json",
		}).
		SetBody(body).
		Post(fullURL)
	
	if err != nil {
		return nil, fmt.Errorf("private POST request failed: %w", err)
	}
	
	return resp, nil
}

// PrivateDelete makes an authenticated DELETE request
func (s *RequestService) PrivateDelete(endpoint string, params string) (interface{}, error) {
	urlStr := s.baseURL + endpoint
	timestamp := time.Now().UnixNano() / 1e6
	
	var message string
	if params == "" {
		message = fmt.Sprintf("timestamp=%d", timestamp)
	} else {
		strParams := s.jsonToParamStr(params)
		message = fmt.Sprintf("%s&timestamp=%d", strParams, timestamp)
	}
	
	sign := s.computeHmac256(message)
	fullURL := fmt.Sprintf("%s?%s&timestamp=%d&signature=%s", urlStr, message, timestamp, sign)
	
	resp, err := s.client.R().
		SetHeaders(map[string]string{
			"X-MEXC-APIKEY": s.apiKey,
			"Content-Type":  "application/json",
		}).
		Delete(fullURL)
	
	if err != nil {
		return nil, fmt.Errorf("private DELETE request failed: %w", err)
	}
	
	return resp, nil
}

// PrivatePut makes an authenticated PUT request
func (s *RequestService) PrivatePut(endpoint string, params string) (interface{}, error) {
	urlStr := s.baseURL + endpoint
	timestamp := time.Now().UnixNano() / 1e6
	
	var message string
	if params == "" {
		message = fmt.Sprintf("timestamp=%d", timestamp)
	} else {
		strParams := s.jsonToParamStr(params)
		message = fmt.Sprintf("%s&timestamp=%d", strParams, timestamp)
	}
	
	sign := s.computeHmac256(message)
	fullURL := fmt.Sprintf("%s?timestamp=%d&signature=%s", urlStr, timestamp, sign)
	
	var body interface{}
	if params != "" {
		if err := json.Unmarshal([]byte(params), &body); err != nil {
			return nil, fmt.Errorf("failed to unmarshal params: %w", err)
		}
	}
	
	resp, err := s.client.R().
		SetHeaders(map[string]string{
			"X-MEXC-APIKEY": s.apiKey,
			"Content-Type":  "application/json",
		}).
		SetBody(body).
		Put(fullURL)
	
	if err != nil {
		return nil, fmt.Errorf("private PUT request failed: %w", err)
	}
	
	return resp, nil
}

// Helper function to convert JSON to query parameters
func (s *RequestService) jsonToParamStr(jsonParams string) string {
	var paramsArr []string
	m := make(map[string]string)
	
	if err := json.Unmarshal([]byte(jsonParams), &m); err != nil {
		return ""
	}
	
	for key, value := range m {
		paramsArr = append(paramsArr, fmt.Sprintf("%s=%s", key, value))
	}
	
	return strings.Join(paramsArr, "&")
}

// Helper function to compute HMAC signature
func (s *RequestService) computeHmac256(message string) string {
	key := []byte(s.apiSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}