package entity

import (
	"github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_blog/dto"
)

type BwArticleData struct {
	ArticleD1List []dto.BwArticleD1 `bson:"articleList" json:"articleList"`
	ProjectName   string            `bson:"projectName" json:"projectName"`
}

func NewBwArticleData(projectName string, articles []dto.BwArticleD1) BwArticleData {
	return BwArticleData{
		ArticleD1List: articles,
		ProjectName:   projectName,
	}
}
