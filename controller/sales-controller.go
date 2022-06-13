package controller

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_frontend_sales/config"
	"github.com/nikitamirzani323/isbpanel_frontend_sales/helpers"
)

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
}
type response_adminsave struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
}

func Sales(c *fiber.Ctx) error {
	render_page := time.Now()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(helpers.Response{}).
		SetError(helpers.Response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"crm_page":   0,
			"crm_status": "PROCESS",
		}).
		Post(config.GetPathAPI() + "api/crm")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*helpers.Response)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*helpers.Response)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Salessve(c *fiber.Ctx) error {
	type payload_crud struct {
		Page            string `json:"page"`
		Sdata           string `json:"sdata" `
		Crm_idusersales int    `json:"crm_idusersales"`
		Crm_idcrmsales  int    `json:"crm_idcrmsales"`
		Crm_idwebagen   int    `json:"crm_idcwebagen"`
		Crm_status      string `json:"crm_status" `
		Crm_status_dua  string `json:"crm_status_dua" `
		Crm_note        string `json:"crm_note" `
		Crm_iduseragen  string `json:"crm_iduseragen" `
		Crm_phone       string `json:"crm_phone" `
		Crm_deposit     int    `json:"crm_deposit" `
	}
	var errors []*helpers.ErrorResponse
	client := new(payload_crud)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	render_page := time.Now()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_adminsave{}).
		SetError(response_adminsave{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"sdata":           client.Sdata,
			"page":            client.Page,
			"crm_idusersales": client.Crm_idusersales,
			"crm_idcrmsales":  client.Crm_idcrmsales,
			"crm_idcwebagen":  client.Crm_idwebagen,
			"crm_status":      client.Crm_status,
			"crm_status_dua":  client.Crm_status_dua,
			"crm_note":        client.Crm_note,
			"crm_iduseragen":  client.Crm_iduseragen,
			"crm_phone":       client.Crm_phone,
			"crm_deposit":     client.Crm_deposit,
		}).
		Post(config.GetPathAPI() + "api/crmsave")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*response_adminsave)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_adminsave)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
