package controller

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_frontend_sales/config"
)

type home struct {
	Page string `json:"page" validate:"required"`
}
type response_login struct {
	Token string `json:"token"`
	Key   string `json:"key"`
}
type response_error_login struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
}
type response_loginhome struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Login(c *fiber.Ctx) error {
	type payload_login struct {
		Username  string `json:"username" `
		Password  string `json:"password" `
		Ipaddress string `json:"ipaddress" `
		Timezone  string `json:"timezone" `
	}

	client := new(payload_login)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(response_login{}).
		SetError(response_error_login{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"username":  client.Username,
			"password":  client.Password,
			"ipaddress": client.Ipaddress,
			"timezone":  client.Timezone,
		}).
		Post(config.GetPathAPI() + "api/login")
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
	if resp.StatusCode() == 200 {
		result := resp.Result().(*response_login)
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status": resp.StatusCode(),
			"token":  result.Token,
			"key":    result.Key,
			"time":   time.Since(render_page).String(),
		})
	} else {
		result_err := resp.Error().(*response_error_login)
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  resp.StatusCode(),
			"token":   "",
			"key":     "",
			"message": result_err.Message,
			"time":    time.Since(render_page).String(),
		})
	}

}
func Home(c *fiber.Ctx) error {
	client := new(home)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_loginhome{}).
		SetError(response_loginhome{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"page": client.Page,
		}).
		Post(config.GetPathAPI() + "api/home")
	if err != nil {
		log.Println(err.Error())
	}

	result := resp.Result().(*response_loginhome)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_loginhome)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
