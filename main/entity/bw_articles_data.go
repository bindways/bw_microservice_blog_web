package entity

import (
	"github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_blog/entity"
)

type BwArticleData struct {
	Articles    []entity.BwArticle `bson:"article" json:"article"`
	ProjectName string             `bson:"projectName" json:"projectName"`
}
