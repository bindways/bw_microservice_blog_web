package server

import (
	"bw_microservice_blog_web/main/core/api"
	"github.com/bindways/bw_microservice_share/bw_helper/bw_error_helper"
	"github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_gateway_frontend/bw_router"
	"github.com/bindways/bw_microservice_share/bw_server"
	bw_handler_gin "github.com/bindways/bw_microservice_share2/bw_gin"
	"github.com/gin-gonic/gin"
)

/**
 * SesServiceCheckoutServer config_files a server gin
 * This is used when there aren't the hashtag "#" in angular url, and we can receive
 * any kind of angular router url specific to load app, so if there aren't a router
 * we pass to index angular file index.html.
 */
func BwWebBlogServer() {
	engine := gin.Default()
	new(bw_handler_gin.BwMetricConfig).
		Constructor1(engine).
		ConfigMiddleware().
		ConfigCustomHandler("health")
	new(api.BwArticleWebController).Constructor1().Controller(engine)
	new(api.BwAssetsController).ConfigAssetsHandler(engine)
	bw_server.BwMicroservicePrintServer(&bw_router.BwMicroserviceBlogWeb)
	err := engine.Run(bw_router.BwMicroserviceBlogWeb.GetPortHttp())
	bw_error_helper.CheckPanic(err)
}
