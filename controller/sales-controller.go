package controller

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_frontend_sales/config"
	"github.com/nikitamirzani323/isbpanel_frontend_sales/helpers"
)

type admindetail struct {
	Username string `json:"username" validate:"required"`
}

type adminsaveiplist struct {
	Sdata     string `json:"sdata" validate:"required"`
	Page      string `json:"page"`
	Username  string `json:"username" validate:"required,alphanum,max=20"`
	Ipaddress string `json:"ipaddress" validate:"required,max=20"`
}
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
