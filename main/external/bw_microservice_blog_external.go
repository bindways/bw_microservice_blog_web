package external

import (
	"encoding/json"
	"fmt"
	"github.com/bindways/bw_microservice_share/bw_helper/bw_feign_client_helper"
	"github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_blog/dto"
	"github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_blog/entity"
	bw_router2 "github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_gateway_microservice/bw_router"
	"github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_oauth2/bw_entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BwMicroserviceBlogExternal struct {
}

/**
 * Get article by id and project
 */
func (t *BwMicroserviceBlogExternal) GetArticleByIdAndProject(
	idArticle primitive.ObjectID, project string) (article entity.BwArticle, err error) {

	urlFill := fmt.Sprintf("%s/project/%s/article/%s", bw_router2.BwMicroserviceBlog.GetLocalFullRouterHttp(), project, idArticle.Hex())
	responseBytes, err := bw_feign_client_helper.BwGetWithParams(urlFill, nil, bw_entity.BwMasterToken)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseBytes, &article)
	return
}

/**
 * Get article by id and project
 */
func (t *BwMicroserviceBlogExternal) GetArticleByName(
	urlName string, project string) (article entity.BwArticle, err error) {

	urlFill := fmt.Sprintf("%s/project/%s/article/name/%s", bw_router2.BwMicroserviceBlog.GetLocalFullRouterHttp(), project, urlName)
	responseBytes, err := bw_feign_client_helper.BwGetWithParams(urlFill, nil, bw_entity.BwMasterToken)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseBytes, &article)
	return
}

/**
 * Get all articles from project and limited size
 */
func (t *BwMicroserviceBlogExternal) GetArticlesByProjectLimitedSize(project string) (articlesD1 dto.BwArticleD1List, err error) {
	urlFill := fmt.Sprintf("%s/%d/articles/d1/project/%s", bw_router2.BwMicroserviceBlog.GetLocalFullRouterHttp(), 5, project)
	responseBytes, err := bw_feign_client_helper.BwGetWithParams(urlFill, nil, bw_entity.BwMasterToken)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseBytes, &articlesD1)
	return
}

// get articles by project name
func (t *BwMicroserviceBlogExternal) GetArticlesByProject(project string) (articles []entity.BwArticle, err error) {
	urlFill := fmt.Sprintf("%s/articles/project/%s", bw_router2.BwMicroserviceBlog.GetLocalFullRouterHttp(), project)
	responseBytes, err := bw_feign_client_helper.BwGetWithParams(urlFill, nil, bw_entity.BwMasterToken)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseBytes, &articles)
	return
}
