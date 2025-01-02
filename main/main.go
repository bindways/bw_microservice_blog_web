package main

import (
	"bw_microservice_blog_web/main/config/server"
	"github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_gateway_frontend/bw_router"
	"github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_logger/bw_logger"
)

func main() {
	new(bw_logger.LogServerClient).Constructor(bw_router.BwMicroserviceBlogWeb.NameMicroservice).Connect()
	server.BwWebBlogServer()
}
