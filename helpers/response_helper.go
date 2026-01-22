package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/fiber-apotek/models"
)

// getStatusText returns "success" or "error" based on the HTTP status code
func getStatusText(status int) string {
	if status >= 200 && status < 300 {
		return "success"
	}
	return "error"
}

// JSONResponse sends a standard JSON response format / structure
func JSONResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	resp := models.Response{
		Status:  getStatusText(status),
		Message: message,
		Data:    data,
	}
	return c.Status(status).JSON(resp)
}

// JSONResponse GetAll sends a standard JSON response format / structure
func JSONResponseGetAll(c *fiber.Ctx, status int, message string, search string, total_items int, current_page int, total_pages int, per_page int, data interface{}) error {
	resp := models.ResponseGetAll{
		Status:      getStatusText(status),
		Message:     message,
		Search:      search,
		TotalItems:  total_items,
		CurrentPage: current_page,
		TotalPages:  total_pages,
		PerPage:     per_page,
		Data:        data,
	}
	return c.Status(status).JSON(resp)
}
