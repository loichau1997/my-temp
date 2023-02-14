package model

import "github.com/jinzhu/gorm/dialects/postgres"

type TagsResponse struct {
	Tags []TagDetail `json:"tags"`
}

type TagDetail struct {
	TagID            int            `json:"tagID"`
	ParentTagIDs     []int64        `json:"parentTagIds"`
	AllNamesByLocale postgres.Jsonb `json:"allNamesByLocale"`
}
