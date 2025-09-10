package api

import (
	web2 "bw_microservice_blog_web/main/core/web"
	"fmt"
	"github.com/bindways/bw_microservice_share2/bw_fiber"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type BwArticleWebController struct {
	articleServiceDep *web2.BwArticleWebService
}

func (t *BwArticleWebController) Constructor1() *BwArticleWebController {
	t.articleServiceDep = (&web2.BwArticleWebService{}).Constructor1()
	return t
}

func (t *BwArticleWebController) Controller(app *fiber.App) {

	/**
	 * return static assets
	 */
	app.Get("/:project/blog/web/assets/*", func(c *fiber.Ctx) error {
		//project := c.Params("project")
		remainPath := c.Params("*")
		return c.SendFile(fmt.Sprintf("static/assets/%s", remainPath))
	})

	/**
	 * Get articles
	 */
	app.Get("/:project/blog/web/", func(c *fiber.Ctx) (err error) {
		project := c.Params("project")
		if err = t.articleServiceDep.GetArticles(c, project); err != nil {
			return bw_fiber.NewErrorResponseFiber(c, err, http.StatusBadRequest)
		}
		return
	})

	/**
	 * Get article by id
	 */
	app.Get("/:project/blog/web/article/name/:urlName", func(c *fiber.Ctx) (err error) {
		urlName := c.Params("urlName")
		project := c.Params("project")
		if err := t.articleServiceDep.GetArticleByName(c, urlName, project); err != nil {
			return bw_fiber.NewErrorResponseFiber(c, err, http.StatusBadRequest)
		}
		return c.SendStatus(http.StatusOK)
	})

	/**
	 * Get article by id (old format)
	 */
	app.Get("/:project/blog/web/article/:id", func(c *fiber.Ctx) (err error) {
		idArticle, err := primitive.ObjectIDFromHex(c.Params("id"))
		project := c.Params("project")
		if err = t.articleServiceDep.GetArticleById(c, idArticle, project); err != nil {
			return bw_fiber.NewErrorResponseFiber(c, err, http.StatusBadRequest)
		}
		return c.SendStatus(http.StatusOK)
	})
}
