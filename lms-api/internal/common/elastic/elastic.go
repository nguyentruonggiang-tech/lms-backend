package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"lms-api/internal/common/env"
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

func (c *ElasticClient) DeleteCourse(ctx context.Context, id int) error {
	req := esapi.DeleteRequest{
		Index:      c.index,
		DocumentID: strconv.Itoa(id),
		Refresh:    "true",
	}
	res, err := req.Do(ctx, c.client)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

func (c *ElasticClient) SearchCourses(ctx context.Context, q string, from, size int) ([]CourseDoc, int, error) {
	query := map[string]any{
		"from": from,
		"size": size,
		"query": map[string]any{
			"bool": map[string]any{
				"must": map[string]any{
					"multi_match": map[string]any{
						"query":  q,
						"fields": []string{"title", "description"},
					},
				},
				"filter": map[string]any{
					"term": map[string]any{"status": "published"},
				},
			},
		},
	}

	body, _ := json.Marshal(query)
	req := esapi.SearchRequest{
		Index: []string{c.index},
		Body:  bytes.NewReader(body),
	}
	res, err := req.Do(ctx, c.client)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, 0, fmt.Errorf("ES search error: %s", res.String())
	}

	var result struct {
		Hits struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
			Hits []struct {
				Source CourseDoc `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, 0, err
	}

	docs := make([]CourseDoc, 0, len(result.Hits.Hits))
	for _, h := range result.Hits.Hits {
		docs = append(docs, h.Source)
	}
	return docs, result.Hits.Total.Value, nil
}
