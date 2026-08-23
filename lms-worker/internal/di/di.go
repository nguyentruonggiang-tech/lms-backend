package di

import (
	"lms-worker/ent"
	"lms-worker/internal/common/elastic"
)

func Injection(entClient *ent.Client, elasticClient *elastic.ElasticClient) {
}
