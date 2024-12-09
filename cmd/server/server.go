package server
import (
	"net/http"
	"github.com/gofiber/fiber/v2"
)
func Run()  {
	app := fiber.New()
	app.Get("/", func (ctx *fiber.ctx) error(
	    return ctx.status(http.statusok).Json(fiber.map{
			"message": "Hello I'm working",
		})
	))
	app.Get("/health",func(ctx *fiber.ctx)error{
		return ctx.status(http.statusok).Json(fiber.map{
			"healthy": true,
	    })
	})
	app.listen(":3000")
}