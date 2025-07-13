package utils

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func PostRequest[T any](client *resty.Client, url string, queryParams map[string]string, headers map[string]string, data interface{}) (*resty.Response, T, interface{}, error) {
	var result T
	var errorResult interface{}
	req := client.R().
		SetError(&errorResult).
		SetBody(data).
		SetResult(&result)

	if len(queryParams) > 0 {
		req = req.SetQueryParams(queryParams)
	}

	if len(headers) > 0 {
		req = req.SetHeaders(headers)
	}

	if data != nil {
		if headers["Content-Type"] == "application/x-www-form-urlencoded" {
			if formData, ok := data.(map[string]interface{}); ok {
				formParams := make(map[string]string)
				for k, v := range formData {
					// Convert values to string; adjust based on actual types
					formParams[k] = fmt.Sprintf("%v", v)
				}
				req = req.SetBody(nil).SetFormData(formParams)
			} else {
				return nil, result, errorResult, fmt.Errorf("data must be a map for x-www-form-urlencoded")
			}
		} else {
			req = req.SetBody(data)
		}
	}

	resp, err := req.Post(url)

	if err != nil {
		return nil, result, errorResult, err
	}

	return resp, result, errorResult, nil
}

func GetRequest[T any](client *resty.Client, url string, queryParams map[string]string, headers map[string]string) (*resty.Response, T, interface{}, error) {
	var result T
	var errorResult interface{}
	req := client.R().
		SetResult(&result).
		SetError(&errorResult)

	if len(queryParams) > 0 {
		req = req.SetQueryParams(queryParams)
	}

	if len(headers) > 0 {
		req = req.SetHeaders(headers)
	}

	resp, err := req.Get(url)
	if err != nil {
		return nil, result, errorResult, err
	}
	return resp, result, errorResult, nil
}

func PutRequest[T any](client *resty.Client, url string, queryParams map[string]string, headers map[string]string, data interface{}) (*resty.Response, T, interface{}, error) {
	var result T
	var errorResult interface{}
	req := client.R().
		SetError(&errorResult).
		SetBody(data).
		SetResult(&result)

	if len(queryParams) > 0 {
		req = req.SetQueryParams(queryParams)
	}

	if len(headers) > 0 {
		req = req.SetHeaders(queryParams)
	}

	resp, err := req.Put(url)

	if err != nil {
		return nil, result, errorResult, err
	}
	return resp, result, errorResult, nil
}

func DeleteRequest[T any](client *resty.Client, url string, queryParams map[string]string, headers map[string]string, data interface{}) (*resty.Response, T, interface{}, error) {
	var result T
	var errorResult interface{}
	req := client.R().
		SetError(&errorResult).
		SetBody(data).
		SetResult(&result)

	if len(queryParams) > 0 {
		req = req.SetQueryParams(queryParams)
	}

	if len(headers) > 0 {
		req = req.SetHeaders(headers)
	}

	resp, err := req.Delete(url)
	if err != nil {
		return nil, result, errorResult, err
	}

	return resp, result, errorResult, nil
}

func PatchRequest[T any](client *resty.Client, url string, queryParams map[string]string, headers map[string]string, data interface{}) (*resty.Response, T, interface{}, error) {
	var result T
	var errorResult interface{}
	req := client.R().
		SetError(&errorResult).
		SetBody(data).
		SetResult(&result)

	if len(queryParams) > 0 {
		req = req.SetQueryParams(queryParams)
	}

	if len(headers) > 0 {
		req = req.SetHeaders(headers)
	}

	resp, err := req.Patch(url)
	if err != nil {
		return nil, result, errorResult, err
	}
	return resp, result, errorResult, nil
}

func CheckResponse(resp *resty.Response) error {
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return nil
}
