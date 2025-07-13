package gateway

import (
	"cth.release/http-gateway/utils"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) error {
	config := utils.GetConfig()
	var query GeneralRequest

	if err := c.QueryParser(&query); err != nil {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	client := resty.New()
	// 모든 헤더를 map[string][]string으로 가져오기
	rawHeaders := c.GetReqHeaders() // map[string][]string

	// map[string]string으로 변환
	headers := make(map[string]string)
	for key, values := range rawHeaders {
		if len(values) > 0 {
			headers[key] = values[0] // 첫 번째 값 사용
			// 또는 모든 값을 결합하려면: headers[key] = strings.Join(values, ",")
		}
	}

	url, err := utils.XorDecrypt(query.Url, config.Secret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	_, res, errResp, err := utils.GetRequest[interface{}](client, url, nil, headers)

	if err != nil || errResp != nil {
		return c.Status(500).JSON(errResp)
	}

	return c.Status(200).JSON(res)
}

func Post(c *fiber.Ctx) error {
	config := utils.GetConfig()
	var query GeneralRequest

	if err := c.QueryParser(&query); err != nil {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	client := resty.New()
	// 모든 헤더를 map[string][]string으로 가져오기
	rawHeaders := c.GetReqHeaders() // map[string][]string

	// map[string]string으로 변환
	headers := make(map[string]string)
	for key, values := range rawHeaders {
		if len(values) > 0 {
			headers[key] = values[0] // 첫 번째 값 사용
			// 또는 모든 값을 결합하려면: headers[key] = strings.Join(values, ",")
		}
	}

	var data interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	url, err := utils.XorDecrypt(query.Url, config.Secret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	_, res, errResp, err := utils.PostRequest[interface{}](client, url, nil, headers, data)

	if err != nil || errResp != nil {
		return c.Status(500).JSON(errResp)
	}

	return c.Status(200).JSON(res)
}

func Put(c *fiber.Ctx) error {
	config := utils.GetConfig()
	var query GeneralRequest

	if err := c.QueryParser(&query); err != nil {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	client := resty.New()
	// 모든 헤더를 map[string][]string으로 가져오기
	rawHeaders := c.GetReqHeaders() // map[string][]string

	// map[string]string으로 변환
	headers := make(map[string]string)
	for key, values := range rawHeaders {
		if len(values) > 0 {
			headers[key] = values[0] // 첫 번째 값 사용
			// 또는 모든 값을 결합하려면: headers[key] = strings.Join(values, ",")
		}
	}

	var data interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	url, err := utils.XorDecrypt(query.Url, config.Secret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	_, res, errResp, err := utils.PutRequest[interface{}](client, url, nil, headers, data)

	if err != nil || errResp != nil {
		return c.Status(500).JSON(errResp)
	}

	return c.Status(200).JSON(res)
}

func Delete(c *fiber.Ctx) error {
	config := utils.GetConfig()
	var query GeneralRequest

	if err := c.QueryParser(&query); err != nil {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	client := resty.New()
	// 모든 헤더를 map[string][]string으로 가져오기
	rawHeaders := c.GetReqHeaders() // map[string][]string

	// map[string]string으로 변환
	headers := make(map[string]string)
	for key, values := range rawHeaders {
		if len(values) > 0 {
			headers[key] = values[0] // 첫 번째 값 사용
			// 또는 모든 값을 결합하려면: headers[key] = strings.Join(values, ",")
		}
	}

	var data interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	url, err := utils.XorDecrypt(query.Url, config.Secret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	_, res, errResp, err := utils.DeleteRequest[interface{}](client, url, nil, headers, data)

	if err != nil || errResp != nil {
		return c.Status(500).JSON(errResp)
	}

	return c.Status(200).JSON(res)
}

func Patch(c *fiber.Ctx) error {
	config := utils.GetConfig()
	var query GeneralRequest

	if err := c.QueryParser(&query); err != nil {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	client := resty.New()
	// 모든 헤더를 map[string][]string으로 가져오기
	rawHeaders := c.GetReqHeaders() // map[string][]string

	// map[string]string으로 변환
	headers := make(map[string]string)
	for key, values := range rawHeaders {
		if len(values) > 0 {
			headers[key] = values[0] // 첫 번째 값 사용
			// 또는 모든 값을 결합하려면: headers[key] = strings.Join(values, ",")
		}
	}

	var data interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	url, err := utils.XorDecrypt(query.Url, config.Secret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	_, res, errResp, err := utils.PatchRequest[interface{}](client, url, nil, headers, data)

	if err != nil || errResp != nil {
		return c.Status(500).JSON(errResp)
	}

	return c.Status(200).JSON(res)
}
