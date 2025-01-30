package web

import (
	"bw_microservice_blog_web/main/entity"
	"bw_microservice_blog_web/main/external"
	"github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_blog/dto"
	entity2 "github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_blog/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"html/template"
	"net/http"
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
func (t *BwArticleWebService) GetArticles(w http.ResponseWriter, project string) (err error) {
	articles, err := t.microserviceBlogExternalDep.GetArticlesByProject(project)
	if err != nil {
		return
	}
	tmpl, err := template.New("articles.html").
		Funcs(t.pipeServiceDep.PipeDateLong()).
		ParseFiles("static/template/articles.html")
	if err != nil {
		return
	}
	articleData := entity.NewBwArticleData(project, dto.NewBwArticlesD1(articles))
	return tmpl.Execute(w, articleData)
}

/**
 * Get article by name
 */
func (t *BwArticleWebService) GetArticleByName(w http.ResponseWriter, urlName string, project string) (err error) {
	article, err := t.microserviceBlogExternalDep.GetArticleByName(urlName, project)
	if err != nil {
		return
	}
	return t.ProcessArticle(w, project, article)
}

/**
 * Get article by id
 */
func (t *BwArticleWebService) GetArticleById(w http.ResponseWriter, idArticle primitive.ObjectID, project string) (err error) {
	article, err := t.microserviceBlogExternalDep.GetArticleByIdAndProject(idArticle, project)
	if err != nil {
		return
	}
	return t.ProcessArticle(w, project, article)
}

/**
 * Process article building as template html and return to frontend
 */
func (t *BwArticleWebService) ProcessArticle(w http.ResponseWriter, project string, article entity2.BwArticle) (err error) {
	articlesD1, err := t.microserviceBlogExternalDep.GetArticlesByProjectLimitedSize(project)
	if err != nil {
		return
	}
	articleD2 := dto.NewBwArticleD2(article, articlesD1)
	tmpl, err := template.
		New("article.html").
		Funcs(t.pipeServiceDep.PipeDateLong()).
		ParseFiles("static/template/article.html")
	if err != nil {
		return
	}
	return tmpl.Execute(w, articleD2)
}
