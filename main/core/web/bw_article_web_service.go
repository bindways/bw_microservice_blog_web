package web

import (
	dto2 "bw_microservice_blog_web/main/dto"
	"bw_microservice_blog_web/main/entity"
	"bw_microservice_blog_web/main/external"
	"github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_blog/dto"
	entity2 "github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_blog/entity"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"html/template"
)

type BwArticleWebService struct {
	microserviceBlogExternalDep *external.BwMicroserviceBlogExternal
	pipeServiceDep              *BwPipeService
}

func (t *BwArticleWebService) Constructor1() *BwArticleWebService {
	return t.Constructor2(
		new(external.BwMicroserviceBlogExternal),
		new(BwPipeService),
	)
}

func (t *BwArticleWebService) Constructor2(
	articleService *external.BwMicroserviceBlogExternal,
	pipeService *BwPipeService) *BwArticleWebService {
	t.microserviceBlogExternalDep = articleService
	t.pipeServiceDep = pipeService
	return t
}

/**
 * Get all articles to customer and SEO
 */
func (t *BwArticleWebService) GetArticles(c *fiber.Ctx, projectName string) (err error) {
	articles, err := t.microserviceBlogExternalDep.GetArticlesByProject(projectName)
	if err != nil {
		return
	}
	tmpl, err := template.New("articles.html").
		Funcs(t.pipeServiceDep.PipeDateLong()).
		ParseFiles("static/template/articles.html")
	if err != nil {
		return
	}
	articlesD1 := dto.NewBwArticlesD1(articles)
	articleData := entity.NewBwArticleData(projectName, articlesD1)
	return tmpl.Execute(c.Type("html"), articleData)
}

/**
 * Get article by name
 */
func (t *BwArticleWebService) GetArticleByName(c *fiber.Ctx, urlName string, projectName string) (err error) {
	article, err := t.microserviceBlogExternalDep.GetArticleByName(urlName, projectName)
	if err != nil {
		return
	}
	return t.ProcessArticle(c, projectName, article)
}

/**
 * Get article by id
 */
func (t *BwArticleWebService) GetArticleById(c *fiber.Ctx, idArticle primitive.ObjectID, project string) (err error) {
	article, err := t.microserviceBlogExternalDep.GetArticleByIdAndProject(idArticle, project)
	if err != nil {
		return
	}
	return t.ProcessArticle(c, project, article)
}

/**
 * Process article building as template html and return to frontend
 */
func (t *BwArticleWebService) ProcessArticle(c *fiber.Ctx, project string, article entity2.BwArticle) (err error) {
	articlesD1, err := t.microserviceBlogExternalDep.GetArticlesByProjectLimitedSize(project)
	if err != nil {
		return
	}
	articleD2 := dto2.NewBwArticleD2(article, articlesD1)
	tmpl, err := template.
		New("article.html").
		Funcs(t.pipeServiceDep.PipeDateLong()).
		ParseFiles("static/template/article.html")
	if err != nil {
		return
	}
	return tmpl.Execute(c.Type("html"), articleD2)
}
