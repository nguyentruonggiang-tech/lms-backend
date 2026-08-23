package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"lms-worker/internal/common/env"
)

type ElasticClient struct {
	client *elasticsearch.Client
	index  string
}

type CourseDoc struct {
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	CategoryID   int     `json:"categoryId"`
	CategoryName string  `json:"categoryName"`
	Level        string  `json:"level"`
	Price        float64 `json:"price"`
	Status       string  `json:"status"`
	CreatedAt    string  `json:"createdAt"`
}

func NewElasticClient(e *env.Env) *ElasticClient {
	cfg := elasticsearch.Config{
		Addresses:              []string{e.ElasticAddrs},
		Username:               e.ElasticUser,
		Password:               e.ElasticPassword,
		CertificateFingerprint: e.ElasticCertFingerprint,
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(fmt.Sprintf("❌ [ELASTICSEARCH] failed to create client: %v", err))
	}

	res, err := client.Ping()
	if err != nil || res.IsError() {
		panic(fmt.Sprintf("❌ [ELASTICSEARCH] failed to connect to %s", e.ElasticAddrs))
	}
	defer res.Body.Close()

	fmt.Println("✅ [ELASTICSEARCH] Connected successfully")
	return &ElasticClient{client: client, index: "courses"}
}

func (c *ElasticClient) IndexCourse(ctx context.Context, doc CourseDoc) error {
	body, _ := json.Marshal(doc)
	req := esapi.IndexRequest{
		Index:      c.index,
		DocumentID: strconv.Itoa(doc.ID),
		Body:       bytes.NewReader(body),
		Refresh:    "true",
	}
	res, err := req.Do(ctx, c.client)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("ES index error: %s", res.String())
	}
	return nil
}
