package server

import (
	"bw_microservice_blog_web/main/core/api"
	"github.com/bindways/bw_microservice_share/bw_helper/bw_error_helper"
	"github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_gateway_frontend/bw_router"
	"github.com/bindways/bw_microservice_share/bw_server"
	"github.com/bindways/bw_microservice_share2/bw_fiber"
	"github.com/gofiber/fiber/v2"
)

/**
 * SesServiceCheckoutServer config_files a server gin
 * This is used when there aren't the hashtag "#" in angular url, and we can receive
 * any kind of angular router url specific to load app, so if there aren't a router
 * we pass to index angular file index.html.
 */
func BwWebBlogServer() {
	app := fiber.New()
	new(bw_fiber.BwMetricConfig).
		Constructor1(app).
		ConfigMiddleware().
		ConfigCustomHandler("health")
	new(api.BwArticleWebController).Constructor1().Controller(app)
	bw_server.BwMicroservicePrintServer(&bw_router.BwMicroserviceBlogWeb)
	err := app.Listen(bw_router.BwMicroserviceBlogWeb.GetPortHttp())
	bw_error_helper.CheckPanic(err)
}
