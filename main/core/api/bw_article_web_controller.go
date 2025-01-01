package api

import (
	web2 "bw_microservice_blog_web/main/core/web"
	"encoding/json"
	bw_helper "github.com/bindways/bw_microservice_share2/bw_gin"
	"github.com/gin-gonic/gin"
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

func (t *BwArticleWebController) Controller(engine *gin.Engine) {

	engine.GET("/:project/blog/web/",
		func(context *gin.Context) {
			project := context.Param("project")
			if err := t.articleServiceDep.GetArticles(context.Writer, project); err != nil {
				_ = json.NewEncoder(context.Writer).Encode(err.Error())
				return
			}
			context.Writer.WriteHeader(http.StatusOK)
		},
	)

	//Get article by id
	engine.GET("/:project/blog/web/article/name/:urlName",
		func(context *gin.Context) {
			urlName := context.Param("urlName")
			project := context.Param("project")
			if err := t.articleServiceDep.GetArticleByName(context.Writer, urlName, project); err != nil {
				bw_helper.NewErrorResponse400(context, err)
				return
			}
			context.Writer.WriteHeader(http.StatusOK)
		},
	)

	//Get article by id (old format)
	engine.GET("/:project/blog/web/article/:id",
		func(context *gin.Context) {
			idArticle, err := primitive.ObjectIDFromHex(context.Param("id"))
			project := context.Param("project")
			if err = t.articleServiceDep.GetArticleById(context.Writer, idArticle, project); err != nil {
				_ = json.NewEncoder(context.Writer).Encode(err.Error())
				return
			}
			context.Writer.WriteHeader(http.StatusOK)
		},
	)

}
