package web

import (
	"github.com/bindways/bw_microservice_share/bw_helper/bw_date_helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"html/template"
)

type BwPipeService struct {
}

/**
 * Pipe to show date
 */
func (t *BwPipeService) PipeDate() template.FuncMap {
	return template.FuncMap{
		"pipeDate": func(id primitive.ObjectID) string {
			return id.Hex()
		},
	}
}

/**
 * Pile to long date
 */
func (t *BwPipeService) PipeDateLong() template.FuncMap {
	return template.FuncMap{
		"pipeDateLong": bw_date_helper.ConvertLongDate,
	}
}

/**
 * Convert id to string
 */
func (t *BwPipeService) PipeObjectIdToHex() template.FuncMap {
	return template.FuncMap{
		"pipeObjectIdToHex": t.convertObjectIdToHex,
	}
}

/**
 * Convert to hex string from id
 */
func (t *BwPipeService) convertObjectIdToHex(id primitive.ObjectID) string {
	return id.Hex()
}
