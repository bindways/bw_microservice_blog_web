package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BwAssetsController struct {
}

func (t *BwAssetsController) ConfigAssetsHandler(engine *gin.Engine) {
	engine.Static("/bw/blog/web/assets/", "static/assets")

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found router.", "path": c.Request.URL.Path})
	})
}
