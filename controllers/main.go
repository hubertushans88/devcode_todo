package controllers

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(ctx *fiber.Ctx, code int, status string, message string, data interface{}) error{
	rsp := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return ctx.Status(code).JSON(rsp)
}

func SendInternalError(ctx *fiber.Ctx)error{
	return ctx.SendStatus(500)
}