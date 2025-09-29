package dto

import (
	"github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_blog/dto"
	"github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_blog/entity"
)

type BwArticleD2 struct {

	/**
	 * Article data
	 */
	Article entity.BwArticle `bson:"article" json:"article"`

	/**
	 * Preview about article
	 */
	ArticleD1List dto.BwArticleD1List `bson:"articleD1List" json:"articleD1List"`

	/**
	 * url of project
	 */
	ProjectName string `bson:"projectName" json:"projectName"`
}

func NewBwArticleD2(article entity.BwArticle, articlesD1 dto.BwArticleD1List) BwArticleD2 {
	return BwArticleD2{
		Article:       article,
		ArticleD1List: articlesD1,
		ProjectName:   article.Project,
	}
}
